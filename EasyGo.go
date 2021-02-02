package EasyGo

import (
	"EasyGo/config"
	"EasyGo/httpMux"
)

type EasyGo struct {
	Mux    *httpMux.Mux
	Config *config.Config
}

var easyGo = &EasyGo{
	Mux:    httpMux.NewMux(),
	Config: config.GetCfg(),
}

//获取单例
func Default() *EasyGo {
	return easyGo
}

//设置配置文件路径
func (easy *EasyGo) SetConfigPath(path string) {
	easy.Config.FilePath = path
}
