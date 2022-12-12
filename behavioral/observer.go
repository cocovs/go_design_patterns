package main

import "fmt"

//观察者模式 由观察对象通知观察者

//抽象层
//定义抽象观察者
type Observer interface {
	Updata()
}

//定义抽象观察对象
type Subject interface {
	Add()
	Notify()
}

//实现层
type Soldier1 struct {
}

func (*Soldier1) Updata() {
	fmt.Println("1号准备战斗")
}

type Soldier2 struct {
}

func (*Soldier2) Updata() {
	fmt.Println("2号准备战斗")
}

//观察对象
type Sentinel struct {
	//内部维护一个列表
	SoldierList []Observer
}

func (S *Sentinel) Add(O Observer) {
	S.SoldierList = append(S.SoldierList, O)
}

func (S *Sentinel) Notify() {
	for _, v := range S.SoldierList {
		v.Updata()
	}
}

func main() {
	s := &Sentinel{}
	s.Add(&Soldier1{})
	s.Add(&Soldier2{})

	s.Notify()

}
