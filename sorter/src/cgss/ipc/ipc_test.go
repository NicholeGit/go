package ipc

import (
	"testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
	re := &Response{method, "handle" + params}
	return re
}
func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("From Client1", "1")
	resp2, _ := client1.Call("From Client2", "2")
	f1 := Response{"From Client1", "handle1"}
	f2 := Response{"From Client2", "handle2"}
	if resp1.Code != f1.Code || resp1.Body != f1.Body {
		t.Error("IpcClient.Call failed. resp1:", resp1)
	}
	if resp2.Code != f2.Code || resp2.Body != f2.Body {
		t.Error("IpcClient.Call failed. resp2:", resp2)
	}
	client1.Close()
	client2.Close()
}
