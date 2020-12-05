/**
 * Created by cks
 * Date: 2020-11-27
 * Time: 9:37
 */
package proto



/* slu */
func (this *Weather) ToString() string {
	result := "当前温度: " + this.TemperatureNow + 
			", 温度范围: " + this.TemperatureRange + 
			", 当前天气: " + this.WeatherNow + 
			", 预测天气: " + this.WeatherPredict
			// ", 风级: " + this.Wind
			// ", 空气质量: " + this.AirQuality
 	return result
}
type Weather struct {
	Code 			 int 	// 0表示有返回结果，其他为错误码
	TemperatureNow 	 string // 当前温度
	TemperatureRange string // 温度范围
	WeatherNow 		 string // 当前天气
	WeatherPredict 	 string // 预测天气
	Wind 			 string // 风级
	AirQuality 		 string // 空气质量
}

type RedisMsg struct {
	Op           int               `json:"op"`
	ServerId     int               `json:"serverId,omitempty"`
	RoomId       int               `json:"roomId,omitempty"`
	UserId       int               `json:"userId,omitempty"`
	Msg          []byte            `json:"msg"`
	Count        int               `json:"count"`
	RoomUserInfo map[string]string `json:"roomUserInfo"`
}

type RedisRoomInfo struct {
	Op           int               `json:"op"`
	RoomId       int               `json:"roomId,omitempty"`
	Count        int               `json:"count,omitempty"`
	RoomUserInfo map[string]string `json:"roomUserInfo"`
}

type RedisRoomCountMsg struct {
	Count int `json:"count,omitempty"`
	Op    int `json:"op"`
}

type SuccessReply struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
