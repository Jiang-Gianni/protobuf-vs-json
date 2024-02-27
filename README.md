# Protobuf vs JSON

Small project to compare the performance and message size differences between [Protocol Buffers](https://protobuf.dev/) and [JSON](https://www.json.org/json-en.html)


The reference proto message is:

```proto
message MyMessage {
    int32 my_int = 1;
    bool my_bool = 2;
    bytes my_bytes = 3;
    string my_string = 4;
    google.protobuf.Timestamp my_timestamp = 5;
    MyStatus my_my_status = 6;
}
```

The reference data is:

```go
var Msg = pb.MyMessage{
	MyInt:       1,
	MyBool:      true,
	MyBytes:     []byte("Hello"),
	MyString:    "World",
	MyTimestamp: timestamppb.Now(),
	MyMyStatus:  pb.MyStatus_STATUS_ONE,
}
```



```bash
 go version
go version go1.21.2 linux/amd64
```


```bash
 go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/Jiang-Gianni/protobuf-vs-json
cpu: AMD Ryzen 5 5600G with Radeon Graphics
BenchmarkProtoMarshalMsg-12              7126260         168.3 ns/op            48 B/op            1 allocs/op
BenchmarkJsonMarshalMsg-12               3293227         362.0 ns/op           144 B/op            1 allocs/op
BenchmarkProtoMarshalMsgList10-12         688795        1466 ns/op             528 B/op            3 allocs/op
BenchmarkJsonMarshalMsgList10-12          363262        3051 ns/op            1681 B/op            3 allocs/op
BenchmarkProtoMarshalMsgList100-12         88579             14021 ns/op            5056 B/op           3 allocs/op
BenchmarkJsonMarshalMsgList100-12          38719             28450 ns/op           15320 B/op       3 allocs/op
BenchmarkProtoUnmarshalMsg-12            5034340               238.0 ns/op            80 B/op       3 allocs/op
BenchmarkJsonUnmarshalMsg-12              536793              2091 ns/op             304 B/op      10 allocs/op
BenchmarkProtoUnmarshalMsgList10-12       407450              2764 ns/op            2232 B/op      46 allocs/op
BenchmarkJsonUnmarshalMsgList10-12         56782             20421 ns/op            2760 B/op      74 allocs/op
PASS
ok      github.com/Jiang-Gianni/protobuf-vs-json        13.019s
```


```bash
 go run .
MsgProtoBytes len:  34
MsgJsonBytes len:  142
MsgListProtoBytes10 len:  360
MsgListJsonBytes10 len:  1442
MsgProtoBytes:  [8 1 16 1 26 5 72 101 108 108 111 34 5 87 111 114 108 100 42 12 8 185 238 148 174 6 16 251 139 242 209 1 48 1]
MsgJsonBytes:  [123 34 109 121 95 105 110 116 34 58 49 44 34 109 121 95 98 111 111 108 34 58 116 114 117 101 44 34 109 121 95 98 121 116 101 115 34 58 34 83 71 86 115 98 71 56 61 34 44 34 109 121 95 115 116 114 105 110 103 34 58 34 87 111 114 108 100 34 44 34 109 121 95 116 105 109 101 115 116 97 109 112 34 58 123 34 115 101 99 111 110 100 115 34 58 49 55 48 55 52 50 51 53 52 53 44 34 110 97 110 111 115 34 58 52 52 48 49 55 52 48 55 53 125 44 34 109 121 95 109 121 95 115 116 97 116 117 115 34 58 49 125]
MsgProtoBytes string: Hello"World*
                                 �����0
MsgJsonBytes string:  {"my_int":1,"my_bool":true,"my_bytes":"SGVsbG8=","my_string":"World","my_timestamp":{"seconds":1707423545,"nanos":440174075},"my_my_status":1}
```

Removing `my_timestamp` from the proto message brings the `len(MsgProtoBytes)` to 20 and `len(MsgJsonBytes)` to 86.

Using an `int64` for the `my_timestamp` field (convertion through `func (t Time) Unix()`) brings the `len(MsgProtoBytes)` to around 26 - 27 and `len(MsgJsonBytes)` to around 112 - 115.
