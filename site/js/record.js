/**
 * Created by chenhao on 8/25/16.
 */
//兼容
window.URL = window.URL || window.webkitURL;
//获取计算机的设备：摄像头或者录音设备
navigator.getUserMedia = navigator.getUserMedia || navigator.webkitGetUserMedia || navigator.mozGetUserMedia || navigator.msGetUserMedia;

var HZRecorder = function (stream, config) {
    config = config || {};
    config.sampleBits = config.sampleBits || 16;      //采样数位 8, 16
    config.sampleRate = config.sampleRate || 16000;   //采样率(1/6 44100)

    //创建一个音频环境对象
    var audioContext = window.AudioContext || window.webkitAudioContext;
    var context = new audioContext();
    var audioInput = context.createMediaStreamSource(stream);
    // 第二个和第三个参数指的是输入和输出都是单声道,2是双声道。
    var recorder = context.createScriptProcessor(4096, 1, 1);

    var audioData = {
        size: 0          //录音文件长度
        , buffer: []     //录音缓存
        , inputSampleRate: context.sampleRate    //输入采样率
        , inputSampleBits: 16       //输入采样数位 8, 16
        , outputSampleRate: config.sampleRate    //输出采样率
        , outputSampleBits: config.sampleBits       //输出采样数位 8, 16
        , input: function (data) {
            this.buffer.push(new Float32Array(data));
            this.size += data.length;
        }
        , compress: function () { //合并压缩
            //合并
            var data = new Float32Array(this.size);
            var offset = 0;
            for (var i = 0; i < this.buffer.length; i++) {
                data.set(this.buffer[i], offset);
                offset += this.buffer[i].length;
            }
            //压缩
            var compression = parseInt(this.inputSampleRate / this.outputSampleRate);
            var length = data.length / compression;
            var result = new Float32Array(length);
            var index = 0, j = 0;
            while (index < length) {
                result[index] = data[j];
                j += compression;
                index++;
            }
            return result;
        }
        , encodeWAV: function () {
            var sampleRate = Math.min(this.inputSampleRate, this.outputSampleRate);
            var sampleBits = Math.min(this.inputSampleBits, this.outputSampleBits);
            var bytes = this.compress();
            var dataLength = bytes.length * (sampleBits / 8);
            var buffer = new ArrayBuffer(44 + dataLength);
            var data = new DataView(buffer);

            var channelCount = 1;//单声道
            var offset = 0;

            var writeString = function (str) {
                for (var i = 0; i < str.length; i++) {
                    data.setUint8(offset + i, str.charCodeAt(i));
                }
            }

            if (sampleBits === 8) {
                for (var i = 0; i < bytes.length; i++, offset++) {
                    var s = Math.max(-1, Math.min(1, bytes[i]));
                    var val = s < 0 ? s * 0x8000 : s * 0x7FFF;
                    val = parseInt(255 / (65535 / (val + 32768)));
                    data.setInt8(offset, val, true);
                }
            } else {
                for (var i = 0; i < bytes.length; i++, offset += 2) {
                    var s = Math.max(-1, Math.min(1, bytes[i]));
                    data.setInt16(offset, s < 0 ? s * 0x8000 : s * 0x7FFF, true);
                }
            }

            return new Blob([data], { type: 'audio/wav' });
        }
    };

    //开始录音
    this.start = function () {
        audioInput.connect(recorder);
        recorder.connect(context.destination);
    }

    //停止
    this.stop = function () {
        recorder.disconnect();
    }

    //获取音频文件
    this.getBlob = function () {
        this.stop();
        return audioData.encodeWAV();
    }

    //回放
    this.play = function (audio) {
        audio.src = window.URL.createObjectURL(this.getBlob());
    }
    //清除
    this.clear = function(){
        audioData.buffer=[];
        audioData.size=0;
    }

    //上传
    this.upload = function (url, op, roomId, authToken, callback) {
        console.log('this.getBlob()', this.getBlob())
        var reader = new FileReader()
        reader.readAsArrayBuffer(this.getBlob())
        this.clear()
        reader.onload = function() {
            // console.log(reader.result);
            var msg = new Int8Array(reader.result);
            // fd.append("msg", msg);
            var xhr = new XMLHttpRequest();
            // console.log(msg);
            if (callback) {
                xhr.upload.addEventListener("progress", function (e) {
                    callback('uploading', e);
                }, false);
                xhr.addEventListener("load", function (e) {
                    callback('ok', e);
                }, false);
                xhr.addEventListener("error", function (e) {
                    callback('error', e);
                }, false);
                xhr.addEventListener("abort", function (e) {
                    callback('cancel', e);
                }, false);
            }

            let jsonData = {roomId: 1, msg: arrayBufferToBase64(msg), authToken: getLocalStorage("authToken")};
            console.log(reader.result)
            $.ajax({
                type: "POST",
                dataType: "json",
                url: url,
                data: JSON.stringify(jsonData),
                success: function (result) {
                    if (result.code != 0) {
                        swal("request error，please login!");
                    }
                },
                error: function () {
                    swal("sorry, exception!");
                }
            });
            function arrayBufferToBase64( buffer ) {
                var binary = '';
                var bytes = new Uint8Array( buffer );
                var len = bytes.byteLength;
                for (var i = 0; i < len; i++) {
                    binary += String.fromCharCode( bytes[ i ] );
                }
                return window.btoa( binary );
            }
            // xhr.open("POST", url);
            // xhr.send(fd);
            
        }
        

        // alert(authToken);
    }

    //音频采集
    recorder.onaudioprocess = function (e) {
        audioData.input(e.inputBuffer.getChannelData(0));
        //record(e.inputBuffer.getChannelData(0));
    }

};
//抛出异常
HZRecorder.throwError = function (message) {
    alert(message);
    throw new function () { this.toString = function () { return message; } }
}
//是否支持录音
HZRecorder.canRecording = (navigator.getUserMedia != null);
//获取录音机
HZRecorder.get = function (callback, config) {
    if (callback) {
        if (navigator.getUserMedia) {
            navigator.getUserMedia(
                { audio: true } //只启用音频
                , function (stream) {
                    var rec = new HZRecorder(stream, config);
                    callback(rec);
                }
                , function (error) {
                    switch (error.code || error.name) {
                        case 'PERMISSION_DENIED':
                        case 'PermissionDeniedError':
                            HZRecorder.throwError('用户拒绝提供信息。');
                            break;
                        case 'NOT_SUPPORTED_ERROR':
                        case 'NotSupportedError':
                            HZRecorder.throwError('浏览器不支持硬件设备。');
                            break;
                        case 'MANDATORY_UNSATISFIED_ERROR':
                        case 'MandatoryUnsatisfiedError':
                            HZRecorder.throwError('无法发现指定的硬件设备。');
                            break;
                        default:
                            HZRecorder.throwError('无法打开麦克风。异常信息:' + (error.code || error.name));
                            break;
                    }
                });
        } else {
            HZRecorder.throwErr('当前浏览器不支持录音功能。'); return;
        }
    }
};