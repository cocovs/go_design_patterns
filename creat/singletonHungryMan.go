package main

//饿汉式单例 无论使用或者不使用，单例都会创建出来
import "fmt"

type singleton struct {
}

func (s *singleton) say() {
	fmt.Println("这里是单例的方法")
}

//对外部私有，只有对外暴露的get方法能够获取该对象
var instance *singleton = new(singleton)

//对外提供一个方法获取这个对象

func GetInstance() *singleton {
	return instance
}

func main() {
	i := GetInstance()
	i.say()
}
