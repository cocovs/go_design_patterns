package main

import "fmt"

//建造者模式主要有主管来进行各个部分的管理，首先将建造者子类作为参数传递给构造函数，构造函数内将参数传递给主管类。
//获得主管类后，主管类通过建造者子类的不同来完后后续的创建工作
//建造者模式 builder 解决的问题是 构造和装配的解耦，用户无须担心复杂对象的创建过程，只需要指定复杂对象的类型即可得到该对象。

//建造者接口 规定建造者子类需要实现的方法
type Builder interface {
	part1()
	part2()
	part3()
}

//主管类
type Director struct {
	builder Builder
}

func (d *Director) Construct() {
	d.builder.part1()
	d.builder.part2()
	d.builder.part3()
}

//构造函数
//传入子类建造者类型 返回一个包含其的主管类
func newDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

//建造者子类
type builder1 struct {
}

//建造者需要实现的方法
func (*builder1) part1() {
	fmt.Println("hello, this is part1")
}
func (*builder1) part2() {
	fmt.Println("hello, this is part2")
}
func (*builder1) part3() {
	fmt.Println("hello, this is part3")
}

func main() {
	//获取建造者子类
	builder := &builder1{}
	dir := newDirector(builder)
	dir.Construct()
}
