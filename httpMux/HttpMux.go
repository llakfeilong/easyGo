package httpMux

import (
	"EasyGo/errorType"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Mux struct {
	nodes []*Node
}

func NewMux() *Mux {
	return &Mux{
		nodes: make([]*Node, 0), //创建空切片
	}
}

type MuxHandler func(m *MuxContext)

func (m *Mux) PUT(path string, handler MuxHandler) {
	m.regHandler(http.MethodPut, path, handler)
}

func (m *Mux) GET(path string, handler MuxHandler) {
	m.regHandler(http.MethodGet, path, handler)
}

func (m *Mux) POST(path string, handler MuxHandler) {
	m.regHandler(http.MethodPost, path, handler)
}

//注册到map里面
func (m *Mux) regHandler(method, path string, handlerFunc MuxHandler) {

	if method == "" {
		panic("method must not be empty")
	}
	if len(path) < 1 || path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}
	if handlerFunc == nil {
		panic("handle must not be nil")
	}
	node := NewNode(path, method, handlerFunc)
	//给切片追加赋值
	m.nodes = append(m.nodes, node)

}

/**
* 拦截解析器
 */
func (m *Mux) httpInterceptor(handler MuxHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("拦截器" + " /request method =" + r.Method + " /request url=" + r.RequestURI + " url.path=" + r.URL.Path + " /addr:" + r.RemoteAddr)
		var isReg = false
		var context *MuxContext
		context = NewContext(w, r)
		for _, node := range m.nodes {
			if strings.EqualFold(node.requestMethod, r.Method) && node.path == r.URL.Path {
				//已经注册的方法
				isReg = true
				handler(context)
			}
		}
		if !isReg {
			context.WritedString(false, errorType.REQUESTMETHODFAIL, errorType.BussinessErrorType[errorType.REQUESTMETHODFAIL], errorType.BussinessErrorType[errorType.REQUESTMETHODFAIL])
		}
	}
}

/**
 * 启动服务
 */
func (m *Mux) Run(address string) {
	mux := http.NewServeMux()
	for _, node := range m.nodes {
		mux.HandleFunc(node.path, m.httpInterceptor(node.handler))
	}
	if err := http.ListenAndServe(address, mux); err != nil {
		fmt.Println("start http server fail:", err)
	}
}
