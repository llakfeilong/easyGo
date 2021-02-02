package httpMux

type Node struct {
	path          string
	requestMethod string
	handler       MuxHandler
}

func NewNode(path string, method string, handler MuxHandler) *Node {
	return &Node{
		path:          path,
		requestMethod: method,
		handler:       handler,
	}
}
