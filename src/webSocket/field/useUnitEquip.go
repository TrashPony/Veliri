package field

import (
	"github.com/gorilla/websocket"
	"fmt"
)

func UseUnitEquip(msg Message, ws *websocket.Conn)  {
	fmt.Printf("%+v\n", msg)

}
