package main

import "fmt"

//代理模式
//
type Goods struct {
	Kind string //商品种类
	Fact bool   //商品的真伪
}

//抽象层
type Shopping interface {
	Buy(good *Goods) //某任务
}

//海外代理、韩国购物、美国购物都实现了shopping
//他们都是Shopping

//实现层
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国购物，买了：", goods.Kind)
}

type AmericaShopping struct{}

func (As *AmericaShopping) Buy(goods *Goods) {
	fmt.Println("去美国购物，买了：", goods.Kind)
}

//代理类
type OverSeasProxy struct {
	Shopping Shopping //代理某个主题 ，这里是抽象的数据类型
}

//new
func NewProxy(shopping Shopping) Shopping {
	return &OverSeasProxy{Shopping: shopping}
}

func (op *OverSeasProxy) Buy(goods *Goods) {
	//1 辨别真伪
	if op.distinguish(goods) == true {
		//2 调用具体要被代理的购物buy方法
		op.Shopping.Buy(goods)
		//3 海关安检
		op.check(goods)
	}

}

func (op *OverSeasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对:", goods.Kind, "进行了辨别真伪。")
	if goods.Fact == false {
		fmt.Println("发现假货:", goods.Kind, " ,不应该购买")
	}

	return goods.Fact
}

func (op *OverSeasProxy) check(goods *Goods) {
	fmt.Println("对", goods.Kind, " 进行了海关检查")
}

func main() {
	//根据对象类型的不同 代理方法中的某个子方法也会不同
	g1 := Goods{
		Kind: "韩国面膜",
		Fact: true,
	}
	g2 := Goods{
		Kind: "四级证书",
		Fact: false,
	}
	//创建一个shopping对象
	var KShopping Shopping
	KShopping = new(KoreaShopping)
	//传入具体的shopping对象，得到代理
	var k_proxy Shopping
	//将具体的类转换为代理类
	k_proxy = NewProxy(KShopping)
	//通过代理模式进行动作
	k_proxy.Buy(&g1)
	k_proxy.Buy(&g2)

	var AShopping Shopping
	AShopping = new(AmericaShopping)

	var A_proxy Shopping
	A_proxy = NewProxy(AShopping)

	A_proxy.Buy(&g1)
}
