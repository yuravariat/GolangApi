package ZMQ

import (
	"flag"
	"fmt"
	"time"

	czmq "github.com/zeromq/goczmq"
)

func Test() {

	flag.Parse()

	pullSock := czmq.NewSock(czmq.Pull)
	defer pullSock.Destroy()

	err := pullSock.Connect("tcp://*:5555")
	if err != nil {
		panic(err)
	}

	go func() {
		pushSock := czmq.NewSock(czmq.Push)
		defer pushSock.Destroy()
		_, err := pushSock.Bind("tcp://*:5555")
		if err != nil {
			panic(err)
		}

		for i := 0; i < 10; i++ {
			str := fmt.Sprint("Message from server ", i)
			err = pushSock.SendFrame([]byte(str), 0)
			if err != nil {
				panic(err)
			}
		}
		time.Sleep(5000 * time.Millisecond)
		for i := 0; i < 10; i++ {
			str := fmt.Sprint("Message from server ", i)
			err = pushSock.SendFrame([]byte(str), 0)
			if err != nil {
				panic(err)
			}
		}
	}()

	for {
		time.Sleep(300 * time.Millisecond)
		msg, _, err := pullSock.RecvFrame()
		if err != nil {
			panic(err)
		}
		println(string(msg[:]))
	}
}
