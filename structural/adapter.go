package main

import "fmt"

//适配器模式

//适配的目标，依赖v5接口
type v5 interface {
	Usev5()
}

//业务类
type Phone struct {
	//需要5v电压进行充电
	v v5
}

//创建phone
func newPhone(v v5) *Phone {
	return &Phone{v}
}

//手机有充电方法
//phone类中需要有实现了v5接口的类才可以使用Usev5方法
func (p *Phone) charge() {
	fmt.Println("进行充电")
	p.v.Usev5()
}

//适配者（被适配的类）
//220v交流电源
type v220 struct {
}

func (*v220) charge() {
	fmt.Println("220v")
}

//适配器
//需要适配器将220v转换为5v,手机才能使用220v电压进行充电
type Adapter struct {
	v *v220
}

func newAdapter(v220 *v220) *Adapter {
	return &Adapter{v220}
}

//适配器实现了v5接口
//通过适配器将v220转换为v5
func (a *Adapter) Usev5() {
	fmt.Printf("适配器插入电源为： ")
	a.v.charge()
}

func main() {
	myphone := newPhone(newAdapter(new(v220)))
	myphone.charge()
}
