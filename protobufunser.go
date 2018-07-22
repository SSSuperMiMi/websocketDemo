package main

import (
	"fmt"
	"net"
	"os"
	"example"
	"github.com/golang/protobuf/proto"
)

func main() {

	//println( newHub())

	fmt.Printf("Started ProtoBuf Server")

	//Listen to the TCP port
	listener, err := net.Listen("tcp", "127.0.0.1:2110")
	checkError(err)
	for{
		if conn, err := listener.Accept(); err == nil{
			//If err is nil then that means that data is available for us so we take up this data and pass it to a new goroutine
			go handleProtoClient(conn)
		} else{
			continue
		}
	}
}

func handleProtoClient(conn net.Conn){
	fmt.Println("Connection established")

	//Close the connection when the function exits
	defer conn.Close()
	//Create a data buffer of type byte slice with capacity of 4096
	data := make([]byte, 4096)
	//Read the data waiting on the connection and put it in the data buffer
	n,err:= conn.Read(data)
	checkError(err)
	fmt.Println("Decoding Protobuf message")
	//Create an struct pointer of type ProtobufTest.TestMessage struct
	protodata := new(example.Message)
	//Convert all the data retrieved into the ProtobufTest.TestMessage struct type
	err = proto.Unmarshal(data[0:n], protodata)
	checkError(err)

	if protodata.Id == 1 {

		var data1 example.Body1
		err = proto.Unmarshal(protodata.MessageData, &data1)

		println(data1.Id)


		println("sssssss") // 打印第一个 person Name 的值进行反序列化验证
	} else {

		var data1 example.Body2
		err = proto.Unmarshal(protodata.MessageData, &data1)

		println(data1.Name)

		println("ppppppppp") // 打印第一个 person Name 的值进行反序列化验证
	}
}

func checkError(err error){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}