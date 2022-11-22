package main

import (
	"fmt"
	"sync"
)

//定义一个锁
//var lock sync.Mutex
var once sync.Once

var instance *singleton

type singleton struct {
}

//once 是线程安全的
func GetInstance() *singleton {
	//once.Do() 内的函数只能执行一次
	//只有第一次才会执行创建单例
	once.Do(func() {
		instance = new(singleton)
	})
	//之后的都会直接返回单例
	return instance
}

// func GetInstance() *singleton {
// 	//为了线程安全，增加互斥
// 	lock.Lock()
// 	defer lock.Unlock()
// 	//首次GetInstance()方法被调用，才会生成单例对象
// 	if instance == nil {
// 		instance = new(singleton)
// 		return instance
// 	}
// 	return instance
// }

func (s *singleton) Say() {
	fmt.Println("hello")
}

func main() {
	in := GetInstance()
	in.Say()
}
