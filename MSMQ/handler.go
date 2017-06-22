package MSMQ

import (
	"fmt"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// QueueMsg is struct created to store Queue Messages
type QueueMsg struct {
	Label string
	Body  string
}

func (message QueueMsg) String() string {
	return fmt.Sprintf("Label:%s | Body:%s", message.Label, message.Body)
}

// ReadQueue method returns a list of messages from MSMQ
func ReadQueue(server string, queuetype string, queuename string, search string) []QueueMsg {

	var list []QueueMsg
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("MSMQ.MSMQQueueInfo")
	msmq, _ := unknown.QueryInterface(ole.IID_IDispatch)

	var formatName string
	if queuetype == "queue" {
		formatName = "direct=os:" + server + "\\private$\\" + queuename
	} else {
		formatName = "direct=os:" + server + "\\private$\\" + queuename + ";Journal"
	}
	oleutil.PutProperty(msmq, "FormatName", formatName)

	MSMQqueue := oleutil.MustCallMethod(msmq, "Open", 32, 0).ToIDispatch()

	// fmt.Println(oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value())
	isOpen := oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value().(int16)

	if isOpen == 1 {
		fmt.Println("Queue is open now....")
	}
	for {
		msg := oleutil.MustCallMethod(MSMQqueue, "PeekCurrent", false, true, 2000, false).ToIDispatch()
		if msg != nil {
			msgBody := oleutil.MustGetProperty(msg, "Body").Value()
			msgBodyStr := "nil"
			if msgBody != nil {
				msgBodyStr = msgBody.(string)
			}
			msgLabel := oleutil.MustGetProperty(msg, "Label").Value()
			msgLabelStr := "nil"
			if msgLabel != nil {
				msgLabelStr = msgLabel.(string)
			}
			//subStr := search
			// if strings.Contains(msgBody, subStr) {
			// 	list = append(list, QueueMsg{Label: msgLabelStr, Body: msgBodyStr})
			// }
			list = append(list, QueueMsg{Label: msgLabelStr, Body: msgBodyStr})
			oleutil.MustCallMethod(MSMQqueue, "PeekNext", 0, true, 1000, 0).ToIDispatch()
		} else {
			oleutil.MustCallMethod(MSMQqueue, "Close")
			fmt.Println("Queue is closed....")
			break
		}
	}

	return list
}

// WriteQueue method adding message to queue
func WriteQueue(server string, queuetype string, queuename string, label string, body string) {

	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("MSMQ.MSMQQueueInfo")
	msmq, _ := unknown.QueryInterface(ole.IID_IDispatch)

	var formatName string
	if queuetype == "queue" {
		formatName = "direct=os:" + server + "\\private$\\" + queuename
	} else {
		formatName = "direct=os:" + server + "\\private$\\" + queuename + ";Journal"
	}
	oleutil.PutProperty(msmq, "FormatName", formatName)

	MSMQqueue := oleutil.MustCallMethod(msmq, "Open", 32, 0).ToIDispatch()

	// fmt.Println(oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value())
	isOpen := oleutil.MustGetProperty(MSMQqueue, "IsOpen").Value().(int16)

	if isOpen == 1 {
		fmt.Println("Queue is open now....")
	}
	unk, _ := oleutil.CreateObject("MSMQ.MSMQMessage")
	MSMQmessage, _ := unk.QueryInterface(ole.IID_IDispatch)
	MSMQmessage.PutProperty("Label", label)
	MSMQmessage.PutProperty("Body", body)

	variant := ole.VARIANT{}
	variant.Val = 3
	msg, err := MSMQmessage.CallMethod("Send", msmq, "3")
	//msg := oleutil.MustCallMethod(MSMQmessage, "Send", &msmq, &variant).ToIDispatch()
	if err != nil {
		println(err)
	}
	if msg != nil {
		println(msg)
	}
	oleutil.MustCallMethod(MSMQqueue, "Close")
	fmt.Println("Queue is closed....")
}
