package main

import "testing"

func BenchmarkProtoMarshalMsg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProtoMarshalMsg()
	}
}

func BenchmarkJsonMarshalMsg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonMarshalMsg()
	}
}

func BenchmarkProtoMarshalMsgList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProtoMarshalMsgList(10)
	}
}

func BenchmarkJsonMarshalMsgList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonMarshalMsgList(10)
	}
}

func BenchmarkProtoMarshalMsgList100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProtoMarshalMsgList(100)
	}
}

func BenchmarkJsonMarshalMsgList100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonMarshalMsgList(100)
	}
}

func BenchmarkProtoUnmarshalMsg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProtoUnmarshalMsg()
	}
}

func BenchmarkJsonUnmarshalMsg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonUnmarshalMsg()
	}
}

func BenchmarkProtoUnmarshalMsgList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProtoUnmarshalMsgList10()
	}
}

func BenchmarkJsonUnmarshalMsgList10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonUnmarshalMsgList10()
	}
}
