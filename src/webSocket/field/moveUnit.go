package field

import "websocket-master"

func MoveUnit(msg FieldMessage, ws *websocket.Conn)  {
	var resp FieldResponse // TODO реализовать метод хождения хотя бы без поиска пути

	resp = FieldResponse{Event: msg.Event, UserName: usersFieldWs[ws].Login, X: msg.X, Y: msg.Y, ErrorType: "not allow"}
	fieldPipe <- resp
}