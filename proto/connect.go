/**
 * Created by cks
 * Date: 2020-12-02
 * Time: 15:51
 */
 package proto

 type Msg struct {
	 Ver       int    `json:"ver"`  // protocol version
	 Operation int    `json:"op"`   // operation for request
	 SeqId     string `json:"seq"`  // sequence number chosen by client
	 Body      []byte `json:"body"` // binary body bytes
 }
 
 type PushMsgRequest struct {
	 UserId int
	 Msg    Msg
 }
 
 type PushRoomMsgRequest struct {
	 RoomId int
	 Msg    Msg
 }
 
 type PushRoomCountRequest struct {
	 RoomId int
	 Count  int
 }
 