package socket

import (
	"log"
	"net"
	"time"
)

type Socket struct {
	ChannelFutures []*ChannelFuture //socket通道切片
	Address        string           //监听IP地址
}

func NewSocket() *Socket {
	return &Socket{
		ChannelFutures: make([]*ChannelFuture, 0),
	}
}

//通道
type ChannelFuture struct {
	Port           string        //监听端口
	Timeout        time.Duration //没有数据交互的超时时间
	BufferSize     int           //缓冲区大小
	ChannelHandler ChannelHandler
}

//缓冲区数据结构体
type ByteBuf struct {
	ByteBuffer []byte
}

type ChannelHandler func(b *ByteBuf)

func (s *Socket) MakeTcpChannel(port string, bufferSize int, timeout time.Duration, handler ChannelHandler) *ChannelFuture {
	return &ChannelFuture{
		Port:           port,
		Timeout:        timeout,
		BufferSize:     bufferSize,
		ChannelHandler: handler,
	}
}

func newByteBuf(buf []byte) *ByteBuf {
	return &ByteBuf{
		ByteBuffer: buf,
	}
}

//绑定服务
func (s *Socket) BindServer(channels ...*ChannelFuture) {
	for _, channel := range channels {
		s.ChannelFutures = append(s.ChannelFutures, channel)
	}

}

//设置TCP监听地址
func (s *Socket) ListenTcp() {
	for _, channel := range s.ChannelFutures {
		log.Println("###监听端口:" + channel.Port)
		addr, _ := net.ResolveTCPAddr("tcp4", s.Address+":"+channel.Port)
		listener, _ := net.ListenTCP("tcp", addr)
		for {
			conn, err := listener.Accept()
			if err != nil {
				continue
			}
			go handlerClient(conn, channel)
		}
	}

}

//处理客户端消息
func handlerClient(conn net.Conn, c *ChannelFuture) {
	if c.Timeout != 0 {
		conn.SetReadDeadline(time.Now().Add(c.Timeout)) //设置连接超时时间
	}
	buf := make([]byte, c.BufferSize) //设置读取数据缓冲区
	defer conn.Close()                // close connection before exit
	for {
		read_len, err := conn.Read(buf) //读取客户端发来的数据
		if err != nil {
			log.Println(err)
			break
		}
		if read_len == 0 {
			//数据长度为空
			break
		}
		//把数据回调到handler作处理
		c.ChannelHandler(newByteBuf(buf[:read_len]))
		////转成16进制string
		//result := hex.EncodeToString(buf[:read_len])
		//log.Println(result) //打印数据
		//conn.Write(buf[:read_len])
	}
}
