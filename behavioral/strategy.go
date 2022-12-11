package main

import "fmt"

//策略模式

//首先定义一个策略接口
type Weapon interface {
	UseWeapon()
}

//具体的策略
type Gun struct{}

func (*Gun) UseWeapon() {
	fmt.Println("使用枪")
}

type knife struct{}

func (*knife) UseWeapon() {
	fmt.Println("使用刀")
}

//环境类
//环境类拥有一个抽象策略字段，可以通过传入不同的策略去执行不同的算法
type hero struct {
	we Weapon
}

//设置策略（替换策略）
func (h *hero) SetWeaponStrategy(we Weapon) {
	h.we = we
}

//调用策略
func (h *hero) Attack() {
	h.we.UseWeapon()
}

func main() {
	he := &hero{}

	he.SetWeaponStrategy(new(knife))
	he.Attack()

	he.SetWeaponStrategy(new(Gun))
	he.Attack()
}
