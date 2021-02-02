# EasyGo
Fast and easy development go


如果大家有什么好建议欢迎留言



1.结构体校验使用示例 



type Tests struct {

	Id  string `json:"id" @NotNull:"DefaultMsg:id不能为空" @Length:"Value:2;DefaultMsg:id长度必须少于2位"`
	Age int    `json:"age" @NotNull:"DefaultMsg:Age不能为空" @Length:"Value:2;DefaultMsg:Age长度必须少于2位"`
	
}



2.httpcontroller示例

func Test(m *httpMux.MuxContext) {

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




3. 路由示例

        router:=httpMux.NewMux()
	router.GET("/download/getFileByteArray",controller.GetFileByteArray)
	
	router.GET("/v1/download/getFileBase64",controller.GetFileBase64)
	
	router.POST("/uploadFile",controller.UploadFile)
	
	router.POST("/v1/uploadFile",controller.UploadFileToFileType)
	
	router.POST("/v1/batchDeleteFiles",controller.BatchDeleteFile)
	
	router.POST("/fileManger/addFileType",controller.AddFileType)
	
	router.POST("/fileManger/updateFileType",controller.UpdateFileType)
	
	router.POST("/fileManger/pageList",controller.PageList)
	
	router.Run(":"+config.GetWebConfig().Port)
