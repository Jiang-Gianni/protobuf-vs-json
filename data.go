package main

import (
	"github.com/Jiang-Gianni/protobuf-vs-json/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var Msg = proto.MyMessage{
	MyInt:       1,
	MyBool:      true,
	MyBytes:     []byte("Hello"),
	MyString:    "World",
	MyTimestamp: timestamppb.Now(),
	MyMyStatus:  proto.MyStatus_STATUS_ONE,
}

func MsgList(n int) *proto.MyMessageList {
	values := make([]*proto.MyMessage, n)
	for i := 0; i < n; i++ {
		values[i] = &Msg
	}
	return &proto.MyMessageList{
		Values: values,
	}
}

var err error
