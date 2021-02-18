/**
* Created by cks
* Date: 2020-12-02
* Time: 17:06
*/
package task

import (
	"SLU-System/config"
	"SLU-System/proto"
	"SLU-System/tools"

	"github.com/sirupsen/logrus"
	"github.com/smallnest/rpcx/client"
	"context"
	"encoding/json"
	"strconv"
	"strings"
)

var RpcConnectClientList map[int]client.XClient

func (task *Task) InitConnectRpcClient() (err error) {
	etcdConfig := config.Conf.Common.CommonEtcd
	d := client.NewEtcdV3Discovery(etcdConfig.BasePath, etcdConfig.ServerPathConnect, []string{etcdConfig.Host}, nil)
	if len(d.GetServices()) <= 0 {
		logrus.Panicf("no etcd server find!")
	}
	RpcConnectClientList = make(map[int]client.XClient, len(d.GetServices()))
	for _, connectConf := range d.GetServices() {
		connectConf.Value = strings.Replace(connectConf.Value, "=&tps=0", "", 1)
		serverId, error := strconv.ParseInt(connectConf.Value, 10, 8)
		if error != nil {
			logrus.Panicf("InitComets err，Can't find serverId. error: %s", error)
		}
		d := client.NewPeer2PeerDiscovery(connectConf.Key, "")
		RpcConnectClientList[int(serverId)] = client.NewXClient(etcdConfig.ServerPathConnect, client.Failtry, client.RandomSelect, d, client.DefaultOption)
		logrus.Infof("InitConnectRpcClient addr %s, v %+v", connectConf.Key, RpcConnectClientList[int(serverId)])
	}
	return
}

func (task *Task) pushSingleToConnect(serverId int, userId int, msg []byte) {
	logrus.Infof("pushSingleToConnect Body %s", string(msg))
	pushMsgReq := &proto.PushMsgRequest{
		UserId: userId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpSingleSend,
			SeqId:     tools.GetSnowflakeId(),
			Body:      msg,
		},
	}
	reply := &proto.SuccessReply{}
	//todo lock
	err := RpcConnectClientList[serverId].Call(context.Background(), "PushSingleMsg", pushMsgReq, reply)
	if err != nil {
		logrus.Infof(" pushSingleToConnect Call err %v", err)
	}
	logrus.Infof("reply %s", reply.Msg)
}

func (task *Task) broadcastRoomToConnect(roomId int, msg []byte) {
	pushRoomMsgReq := &proto.PushRoomMsgRequest{
		RoomId: roomId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpRoomSend,
			SeqId:     tools.GetSnowflakeId(),
			Body:      msg,
		},
	}
	reply := &proto.SuccessReply{}
	for _, rpc := range RpcConnectClientList {
		logrus.Infof("broadcastRoomToConnect rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomMsg", pushRoomMsgReq, reply)
		logrus.Infof("reply %s", reply.Msg)
	}
}

func (task *Task) broadcastRoomCountToConnect(roomId, count int) {
	msg := &proto.RedisRoomCountMsg{
		Count: count,
		Op:    config.OpRoomCountSend,
	}
	var body []byte
	var err error
	if body, err = json.Marshal(msg); err != nil {
		logrus.Warnf("broadcastRoomCountToConnect  json.Marshal err :%s", err.Error())
		return
	}
	pushRoomMsgReq := &proto.PushRoomMsgRequest{
		RoomId: roomId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpRoomCountSend,
			SeqId:     tools.GetSnowflakeId(),
			Body:      body,
		},
	}
	reply := &proto.SuccessReply{}
	for _, rpc := range RpcConnectClientList {
		logrus.Infof("broadcastRoomCountToConnect rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomCount", pushRoomMsgReq, reply)
		logrus.Infof("reply %s", reply.Msg)
	}
}

func (task *Task) broadcastRoomInfoToConnect(roomId int, roomUserInfo map[string]string) {
	msg := &proto.RedisRoomInfo{
		Count:        len(roomUserInfo),
		Op:           config.OpRoomInfoSend,
		RoomUserInfo: roomUserInfo,
		RoomId:       roomId,
	}
	var body []byte
	var err error
	if body, err = json.Marshal(msg); err != nil {
		logrus.Warnf("broadcastRoomInfoToConnect  json.Marshal err :%s", err.Error())
		return
	}
	pushRoomMsgReq := &proto.PushRoomMsgRequest{
		RoomId: roomId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpRoomInfoSend,
			SeqId:     tools.GetSnowflakeId(),
			Body:      body,
		},
	}
	reply := &proto.SuccessReply{}
	for _, rpc := range RpcConnectClientList {
		logrus.Infof("broadcastRoomInfoToConnect rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomInfo", pushRoomMsgReq, reply)
		logrus.Infof("broadcastRoomInfoToConnect rpc  reply %v", reply)
	}
}

func (task *Task) broadcastCSluAnswerToConnect(roomId int, msg []byte) {
	m := &proto.Send{}
	if err := json.Unmarshal(msg, m); err != nil {
		logrus.Infof(" json.Unmarshal err:%v ", err)
	}
	answer, source, _ := task.SLUAnswer(string(m.Msg))
	// logrus.Infof(string(msg))
	logrus.Infof("answer %s, source %s", answer, source)
	
	m.Msg = answer
	m.FromUserId = 3
	m.FromUserName = "robot"
	mByte, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf(" json.marshal err:: %s", err.Error())
		return
	}
	pushRoomMsgReq := &proto.PushRoomMsgRequest{
		RoomId: roomId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpSluContent,
			SeqId:     tools.GetSnowflakeId(),
			Body:      mByte,
		},
	}
	reply := &proto.SuccessReply{}
	for _, rpc := range RpcConnectClientList {
		logrus.Infof("broadcastCSluAnswerToConnect rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomMsg", pushRoomMsgReq, reply)
		logrus.Infof("reply %s", reply.Msg)
	}
}

func (task *Task) broadcastAudioSluAnswerToConnect(roomId int, msg []byte) {
	logrus.Infof("audio msg: %s", string(msg))
	m := &proto.Send{}
	if err := json.Unmarshal(msg, m); err != nil {
		logrus.Infof(" json.Unmarshal err:%v ", err)
	}
	text, err := task.Trans(string(m.Msg))
	if err != nil {
		logrus.Errorf(" trans audio to text err :: %s", err.Error())
	}
	
	m.Msg = text
	mByte, err := json.Marshal(m)
	if err != nil {
		logrus.Errorf(" json.marshal err:: %s", err.Error())
		return
	}
	pushRoomMsgReq := &proto.PushRoomMsgRequest{
		RoomId: roomId,
		Msg: proto.Msg{
			Ver:       config.MsgVersion,
			Operation: config.OpSluContent,
			SeqId:     tools.GetSnowflakeId(),
			Body:      mByte,
		},
	}
	reply := &proto.SuccessReply{}
	for _, rpc := range RpcConnectClientList {
		logrus.Infof("broadcastCSluAnswerToConnect rpc  %v", rpc)
		rpc.Call(context.Background(), "PushRoomMsg", pushRoomMsgReq, reply)
		logrus.Infof("reply %s", reply.Msg)
	}
	task.broadcastCSluAnswerToConnect(roomId, mByte)
}