package grammar

import (
	"fmt"
)

type Connector interface {
	SetName(name string) string
	GetName() string
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) SetName(param string) string {
	pc.name = param
	return pc.name
}

func (pc PhoneConnector) GetName() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("PC Connect:", pc.name)
}

type TVConnector struct {
	name string
}

func (tv TVConnector) SetName(param string) string {
	tv.name = param
	return tv.name
}

func (tv TVConnector) GetName() string {
	return tv.name
}

func (tv TVConnector) Connect() {
	fmt.Println("TV Connected.", tv.name)
}
