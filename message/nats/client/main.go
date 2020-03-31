package client

import (
	"flag"
	"log"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/pborman/uuid"
)

const (
	//url   = "nats://192.168.3.125:4222"
	url = nats.DefaultURL
)

var (
	nc  *nats.Conn
	err error
)

func init() {
	if nc, err = nats.Connect(url); checkErr(err) {
		//
	}
}

func main() {
	var (
		subj = flag.String("subj", "", "subject name")
	)
	flag.Parse()
	log.Println(*subj)
	startClient(*subj)

	time.Sleep(time.Second)
}

//send message to server
func startClient(subj string) {
	for i := 0; i < 1; i++ {
		id := uuid.New()
		log.Println(id)
		nc.Publish(subj, []byte(id+" Sun "+strconv.Itoa(i)))
		nc.Publish(subj, []byte(id+" Rain "+strconv.Itoa(i)))
		nc.Publish(subj, []byte(id+" Fog "+strconv.Itoa(i)))
		nc.Publish(subj, []byte(id+" Cloudy "+strconv.Itoa(i)))
	}
}

func checkErr(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
