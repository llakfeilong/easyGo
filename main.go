package easyGo

import (
	"github.com/llakfeilong/easyGo/socket"
	"log"
	"time"
)

func main() {
	s := Socket()
	test1channel := s.MakeTcpChannel("8079", 1024, 2*time.Minute, tcphandler1)
	test2channel := s.MakeTcpChannel("8987", 1024, 0, tcphandler2)
	s.BindServer(test1channel, test2channel)
	s.ListenTcp()
}

func tcphandler1(buf *socket.ByteBuf) {
	log.Println("接收到数据:", buf.ByteBuffer)
}

func tcphandler2(buf *socket.ByteBuf) {
	log.Println("接收到数据2:", buf.ByteBuffer)
}
