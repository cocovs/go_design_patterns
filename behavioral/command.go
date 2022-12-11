package main

import "fmt"

//比如说游戏中的操作命令,可以将客户端发送的命令被服务端存储起来，服务端有单线程处理这些请求。

//定义命令接口
type Command interface {
	Execute()
}

//移动
type Move struct {
}

func (m *Move) Execute() {
	fmt.Println("移动")
}

type Attack struct {
}

func (a *Attack) Execute() {
	fmt.Println("攻击")
}

func main() {
	CommandList := make([]Command, 0)

	CommandList = append(CommandList, new(Attack))
	CommandList = append(CommandList, new(Move))
	CommandList = append(CommandList, new(Move))
	CommandList = append(CommandList, new(Move))
	CommandList = append(CommandList, new(Attack))

	for _, i := range CommandList {
		i.Execute()
	}
}
