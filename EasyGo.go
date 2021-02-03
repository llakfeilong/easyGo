package EasyGo

import (
	"github.com/llakfeilong/EasyGo/config"
	"github.com/llakfeilong/EasyGo/httpMux"
)

//获取路由
func Default() *httpMux.Mux {
	return httpMux.NewMux()
}

//获取配置
func Config() *config.Config {
	return config.GetCfg()
}
