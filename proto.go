package main

import (
	my "github.com/Jiang-Gianni/protobuf-vs-json/proto"
	"google.golang.org/protobuf/proto"
)

var (
	MsgProtoBytes       []byte
	MsgListProtoBytes10 []byte
)

func init() {
	MsgProtoBytes, err = proto.Marshal(&Msg)
	if err != nil {
		panic(err)
	}
	MsgListProtoBytes10, err = proto.Marshal(MsgList(10))
	if err != nil {
		panic(err)
	}
}

func ProtoMarshalMsg() {
	_, err := proto.Marshal(&Msg)
	if err != nil {
		panic(err)
	}
}

func ProtoMarshalMsgList(n int) {
	_, err := proto.Marshal(MsgList(n))
	if err != nil {
		panic(err)
	}
}

func ProtoUnmarshalMsg() {
	err := proto.Unmarshal(MsgProtoBytes, &Msg)
	if err != nil {
		panic(err)
	}
}

func ProtoUnmarshalMsgList10() {
	var list my.MyMessageList
	err := proto.Unmarshal(MsgListProtoBytes10, &list)
	if err != nil {
		panic(err)
	}
}
