package main

import "fmt"

//template method 模板方法模式
//通过定义一个接口，在接口内定义好一个算法不变的部分，然后可变的行为交给子类来实现

//举个例子 喝茶 和 喝咖啡

//首先定义一个抽象类，包括了饮料的全部制作方法
type Beverage interface {
	//烧水
	BoilWater()
	//放杯子
	PutCup()
	//加料
	AddThings()
}

//模板 封装
type template struct {
	b Beverage
}

//封装的固定模板
func (t *template) MakeBeverage() {
	t.b.BoilWater()
	t.b.PutCup()
	t.b.AddThings()
}

//具体实现类（子类）
type MakeTea struct {
	template
}

func newMakeTea() *MakeTea {
	//MakeTea匿名嵌套了template，实际上就是继承了template
	//因此MakeTea内部就有一个接口b(也就是MakeBeverage)，但是这个接口还没有实现
	//通过将maketea赋给b，使得maketea实现了MakeBeverage接口（MakeTea实现了MakeBeverage接口）
	maketea := &MakeTea{}
	maketea.b = maketea
	return maketea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("烧水")
}

func (mt *MakeTea) PutCup() {
	fmt.Println("放杯子")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("放入茶叶")
}

func main() {
	makeTea := newMakeTea()
	makeTea.MakeBeverage()
}
