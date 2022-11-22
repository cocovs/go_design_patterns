package main

import "fmt"

//简单工厂模式主要实现了 通过工厂类来进行对象的创建，通过传入参数的不同创建不同具体产品类的实例。
//使创建和使用实例的工作分开，使用者不必关心类对象如何创建，实现了解耦。
//更符合面向对象的原则和面向接口编程

type car interface {
	say(name string) string
}

type car1 struct {
}

func (*car1) say(name string) string {
	return fmt.Sprint("car1 say : ", name)
}

type car2 struct {
}

func (*car2) say(name string) string {
	return fmt.Sprint("car2 say : ", name)
}

//SimpleFactory对外可见，封装了实现细节
func SimpleFactory(name int) car {
	if name == 1 {
		return &car1{}
	} else if name == 2 {
		return &car2{}
	}
	return nil
}

func main() {
	myCar := SimpleFactory(1)
	fmt.Println(myCar.say("car1"))

	myCar = SimpleFactory(2)
	fmt.Println(myCar.say("car2"))

}
