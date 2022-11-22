package main

import "fmt"

//工厂方法模式使得工厂父类只用定义工厂子类需要实现的接口，
//添加新产品时只需要添加新的的工厂子类，不需要修改工厂父类的逻辑。
//解决了简单工厂模式不符合开闭原则的问题。

type factoryInterface interface {
	creatFactory(name string)
}

type productInterface interface {
	say()
}

//产品A
type productA struct {
}

//实现了方法say
func (*productA) say() {
	fmt.Println("hello, i am productA")
}

//工厂A
type factoryA struct {
}

//工厂实现了creat方法
func (*factoryA) creatFactory() productInterface {
	return &productA{}
}

//产品B
type productB struct {
}

//产品B实现了say方法
func (*productB) say() {
	fmt.Println("hello, i am productB")
}

//B工厂
type factoryB struct {
}

//产品B工厂实现了creat方法
func (*factoryB) creatFactory() productInterface {
	return &productB{}
}

func main() {
	//常见子类工厂实例
	fA := &factoryA{}
	//通过子类工厂创建产品
	proA := fA.creatFactory()
	//调用B的say方法
	proA.say()

	fB := &factoryB{}
	//通过子类工厂创建产品
	proB := fB.creatFactory()
	//调用B的say方法
	proB.say()

}
