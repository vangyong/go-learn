package grammar

import (
	"fmt"
)

type empty interface {
}

type USB interface {
	Name() string
	Connect()
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connect:", pc.name)
}

type TVConnector struct {
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("Connected.", tv.name)
}

func Disconnect(usb interface{}) {
	//	if pc, ok := usb.(PhoneConnector); ok {
	//		fmt.Println("Disconnected.", pc.name)
	//		return
	//	}
	//	fmt.Println("Unkmown decive.")

	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected.", v.name)
	default:
		fmt.Println("Unkmown decive.")
	}

}
