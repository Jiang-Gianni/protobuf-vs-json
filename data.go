package main

import (
	pb "github.com/Jiang-Gianni/protobuf-vs-json/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var Msg = pb.MyMessage{
	MyInt:       1,
	MyBool:      true,
	MyBytes:     []byte("Hello"),
	MyString:    "World",
	MyTimestamp: timestamppb.Now(),
	// MyTimestamp: time.Now().Unix(),
	MyMyStatus: pb.MyStatus_STATUS_ONE,
}

func MsgList(n int) *pb.MyMessageList {
	values := make([]*pb.MyMessage, n)
	for i := 0; i < n; i++ {
		values[i] = &Msg
	}
	return &pb.MyMessageList{
		Values: values,
	}
}

var err error
