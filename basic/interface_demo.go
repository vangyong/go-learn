package basic

import (
	"fmt"
)

type Connector interface {
	SetName(name string) string
	GetName() string
	Connect()
}

type PhoneConnector struct {
	Name string
}

func (pc PhoneConnector) SetName(param string) string {
	pc.Name = param
	fmt.Println("SetName:"+pc.Name)
	return pc.Name
}

func (pc PhoneConnector) GetName() string {
	fmt.Println("PhoneConnector GetName:"+pc.Name)
	return pc.Name
}

func (pc PhoneConnector) Connect() {
	fmt.Println(pc.Name)
	fmt.Println("PhoneConnector Connected!")
}

//type TVConnector struct {
//	name string
//}
//
//func (tv TVConnector) SetName(param string) string {
//	tv.name = param
//	return tv.name
//}
//
//func (tv TVConnector) GetName() string {
//	return tv.name
//}
//
//func (tv TVConnector) Connect() {
//	fmt.Println("TV Connected.", tv.name)
//}
