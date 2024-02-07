package main

import (
	"fmt"
)

func main() {
	fmt.Println("MsgProtoBytes len: ", len(MsgProtoBytes))
	fmt.Println("MsgJsonBytes len: ", len(MsgJsonBytes))
	fmt.Println("MsgListProtoBytes10 len: ", len(MsgListProtoBytes10))
	fmt.Println("MsgListJsonBytes10 len: ", len(MsgListJsonBytes10))

	fmt.Println("MsgProtoBytes string: ", string(MsgProtoBytes))
	fmt.Println("MsgJsonBytes string: ", string(MsgJsonBytes))
}
