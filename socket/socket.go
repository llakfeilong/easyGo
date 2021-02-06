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
	ReadTimeout    time.Duration //服务端读取不到数据超时时间
	BufferSize     int           //缓冲区大小
	ChannelHandler channelHandler
}

//通道入站处理程序
type ChannelInboundHandler struct {
	ByteBuffer []byte
	Conn       net.Conn
}

type channelHandler func(c *ChannelInboundHandler)

func (s *Socket) MakeTcpChannel(port string, bufferSize int, readTimeout time.Duration, handler channelHandler) *ChannelFuture {
	return &ChannelFuture{
		Port:           port,
		ReadTimeout:    readTimeout,
		BufferSize:     bufferSize,
		ChannelHandler: handler,
	}
}

func newChannelInboundHandler(buf []byte, conn net.Conn) *ChannelInboundHandler {
	return &ChannelInboundHandler{
		ByteBuffer: buf,
		Conn:       conn,
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
	log.Println("###开始监听")
	for _, channel := range s.ChannelFutures {
		go func(channel *ChannelFuture) {
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
		}(channel)
	}
	select {}
}

//处理客户端消息
func handlerClient(conn net.Conn, c *ChannelFuture) {

	buf := make([]byte, c.BufferSize) //设置读取数据缓冲区
	defer conn.Close()                // close connection before exit
	for {
		if c.ReadTimeout != 0 {
			conn.SetReadDeadline(time.Now().Add(c.ReadTimeout)) //设置连接超时时间
		}
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
		c.ChannelHandler(newChannelInboundHandler(buf[:read_len], conn))
		////转成16进制string
		//result := hex.EncodeToString(buf[:read_len])
		//log.Println(result) //打印数据
		//conn.Write(buf[:read_len])
	}
}
