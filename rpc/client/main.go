package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
	"studygo/learngo/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 2}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, err)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result, err)
	}
}

