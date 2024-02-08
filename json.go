package main

import (
	"encoding/json"

	pb "github.com/Jiang-Gianni/protobuf-vs-json/proto"
)

var (
	MsgJsonBytes       []byte
	MsgListJsonBytes10 []byte
)

func init() {
	MsgJsonBytes, err = json.Marshal(&Msg)
	if err != nil {
		panic(err)
	}
	MsgListJsonBytes10, err = json.Marshal(MsgList(10))
	if err != nil {
		panic(err)
	}
}

func JsonMarshalMsg() {
	_, err := json.Marshal(&Msg)
	if err != nil {
		panic(err)
	}
}

func JsonMarshalMsgList(n int) {
	_, err := json.Marshal(MsgList(n))
	if err != nil {
		panic(err)
	}
}

func JsonUnmarshalMsg() {
	err := json.Unmarshal(MsgJsonBytes, &Msg)
	if err != nil {
		panic(err)
	}
}

func JsonUnmarshalMsgList10() {
	var list pb.MyMessageList
	err := json.Unmarshal(MsgListJsonBytes10, &list)
	if err != nil {
		panic(err)
	}
}
