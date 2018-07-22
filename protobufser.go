package main

//客户机
import (
	"fmt"
	"log"
	"net"
	"example"
	"github.com/golang/protobuf/proto"
)

//连接服务器
func connectServer() {
	//接通
	conn, err := net.Dial("tcp", "127.0.0.1:2110")
	checkError(err)
	fmt.Println("连接成功！\n")
	//输入
	body2 := example.Body2{
			Name:string("allen"),
	}


	data1, err := proto.Marshal(&body2)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}

	// 为 AllPerson 填充数据
	msg1 := example.Message{
		Id:*proto.Int32(2),
		MessageData:data1,
	}

	msg2, err := proto.Marshal(&msg1)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}



	conn.Write([]byte(msg2))



}

//检查错误
func checkError(err error) {
	if err != nil {
		log.Fatal("an error!", err.Error())
	}
}

//主函数
func main() {
	//连接servser
	connectServer()
}

