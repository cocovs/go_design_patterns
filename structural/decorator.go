package main

import "fmt"

type Phone interface {
	Show() //构件的功能
}

//装饰器
//包含具体被装饰的类
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

//实现层
type Man struct{}

func (*Man) Show() {
	fmt.Println("this is man")
}

//具体的装饰器
//继承装饰器基础类
type GunDecorator struct {
	Decorator
}

func (gun *GunDecorator) Show() {
	gun.phone.Show()
	fmt.Println("A man with a gun ")
}

func newGunDecorator(ph Phone) *GunDecorator {
	return &GunDecorator{Decorator{ph}}
}

func main() {
	//
	var xiaoming Phone
	xiaoming = new(Man)
	xiaoming.Show()
	fmt.Println("--------")
	//加入装饰器
	var gunMan Phone
	gunMan = newGunDecorator(xiaoming)
	gunMan.Show()

}
