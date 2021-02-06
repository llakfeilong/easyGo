# EasyGo
Fast and easy development go

如果大家觉得好用麻烦点个星星



一.web功能使用示例

//结构体

 type Tests struct {

	Id  string `json:"id" @NotNull:"DefaultMsg:id不能为空" @Length:"Value:2;DefaultMsg:id长度必须少于2位"`
	Age int    `json:"age" @NotNull:"DefaultMsg:Age不能为空" @Length:"Value:2;DefaultMsg:Age长度必须少于2位"`
 }

//controller

func Test (m *httpMux.MuxContext) {

	var test Tests

	//获取post请求JSON 并解析到相应结构体

	json.Unmarshal(m.GetPostJson(), &test)
	
	log.Println("###id =", test.Id, "age:", test.Age)
	
	//校验结构体
	validtor := m.Validator.VaildSturct(test)
	
	//json校验结果
	b, _ := json.Marshal(validtor.ErrorResults)
	
	//遍历失败结果集取出异常原因
	for _, value := range validtor.ErrorResults {
		log.Println(value)
	}
	
	//直接获取校验结构体结果
	if m.Validator.ValidResult {
		//成功
	} else {
		//失败
	}
	
	//输出
	m.WritedString(true, "1000", "sucess", validtor.ErrorResults)
	log.Println("返回结果:", string(b))
}

   //路由示例
    router:=easyGo.DefaultMux()//默认路由
   
	router.GET("/download/getFileByteArray",controller.GetFileByteArray)

	router.POST("/v1/uploadFile",controller.UploadFileToFileType)

	router.Run(":"+config.GetWebConfig().Port)


二. socket 使用示例

   s:=easyGo.Socket()//获取socket
   
   test1channel := s.MakeTcpChannel("8079", 1024, 2*time.Minute, tcphandler1) //绑定端口，缓冲区大小，处理器
   
   test2channel := s.MakeTcpChannel("8987", 1024, 0, tcphandler2)//创建TCP通道并绑定处理器
   
   s.BindServer(test1channel, test2channel) //绑定服务
   
   s.ListenTcp() //开始监听服务

//处理器1

func tcphandler1(buf *socket.ChannelInboundHandler) {
log.Println("接收到数据:", buf.ByteBuffer)

}

//处理器2

func tcphandler2(buf *socket.ChannelInboundHandler) {
log.Println("接收到数据2:", buf.ByteBuffer)
}


三 读取配置在viper基础上封装一些方法

config:=easyGo.Config()

config.SetConfigPath(“配置文件相对路径”)

config.LoadConfig() //加载配置

config.GetString(“key”)  //返回string类型

config.GetUInt64(“key”) //返回uInt64类型

config.GetInt64(“key”) //返回Int64类型


四. 当http和tcp一起使用

s:=socket.NewSocket()

r:=DefaultMux()

r.POST("/test",Test)

r.GET("/v1",Testv1)

r.Run(":8384")

test1channel := s.MakeTcpChannel("8079", 1024, 2*time.Minute, tcphandler1)

test2channel := s.MakeTcpChannel("8987", 1024, 0, tcphandler2)

s.BindServer(test1channel, test2channel)

s.ListenTcp()





