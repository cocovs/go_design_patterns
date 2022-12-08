package main

import "fmt"

//外观模式
//其实就是通过提供一个接口将多个类整合在一起

//举例子：喝茶需要烧水、茶具、不同的茶叶

//这里是是外观角色
type DrinkTea struct {
	gt GreenTea
	bt BlackTea
	bw Water
	ts TeaSet
}

//通过不同的方法来组合不同的类
func (dt *DrinkTea) DrinkBlackTea() {
	fmt.Println("来喝红茶")
	dt.bw.BoilWater()
	dt.ts.PutCup()
	dt.bt.tea()
}

func (dt *DrinkTea) DrinkGreenTea() {
	fmt.Println("来喝绿茶")
	dt.bw.BoilWater()
	dt.ts.PutCup()
	dt.gt.tea()
}

//以下是子系统类
//绿茶
type GreenTea struct {
}

func (tea *GreenTea) tea() {
	fmt.Println("放入绿茶")
}

//红茶
type BlackTea struct {
}

func (tea *BlackTea) tea() {
	fmt.Println("放入红茶")
}

type Water struct {
}

func (water *Water) BoilWater() {
	fmt.Println("起锅烧水")
}

type TeaSet struct {
}

func (ts *TeaSet) PutCup() {
	fmt.Println("放好茶具")
}

func main() {
	dt := new(DrinkTea)
	dt.DrinkBlackTea()
	fmt.Println()
	dt.DrinkGreenTea()
}
