# 设计模式

## 创建

### 简单工厂模式

简单工厂模式主要实现了 **通过工厂类来进行对象的创建**，通过传入参数的不同创建不同具体产品类的实例。
使创建和使用实例的工作分开，使用者不必关心类对象如何创建，**实现了解耦**。
更符合面向对象的原则和面向接口编程。但**违背了开闭原则**，添加新产品必须要修改工厂类的逻辑。

（go中使用工厂方法创建产品类）

优点：

- 将创建实例的工作与使用实例的工作分开，使用者不必关心类对象如何创建，实现了解耦；
- 把初始化实例时的工作放到工厂里进行，使代码更容易维护。 更符合面向对象的原则 & 面向接口编程，而不是面向实现编程。

缺点：

- 工厂类集中了所有实例（产品）的创建逻辑，一旦这个工厂不能正常工作，整个系统都会受到影响；
- 违背“开放 - 关闭原则”，一旦添加新产品就不得不修改工厂类的逻辑，这样就会造成工厂逻辑过于复杂。
- 简单工厂模式由于使用了静态工厂方法，静态方法不能被继承和重写，会造成工厂角色无法形成基于继承的等级结构。

场景：

- 客户如果只知道传入工厂类的参数，对于如何创建对象的逻辑不关心时；
- 当工厂类负责创建的对象（具体产品）比较少时。 



<img src="https://test-1309023885.cos.ap-guangzhou.myqcloud.com/typora/image-20221109180708991.png" style="zoom:80%;" />



go 语言没有构造函数一说，所以一般会定义NewXXX函数来初始化相关类。
NewXXX 函数返回接口时就是简单工厂模式，也就是说Golang的一般推荐做法就是简单工厂。

```go
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

func simpleFactory(name int) car {
	if name == 1 {
		return &car1{}
	} else if name == 2 {
		return &car2{}
	}
	return nil
}

func main() {
	myCar := simpleFactory(1)
	fmt.Println(myCar.say("car1"))

	myCar = simpleFactory(2)
	fmt.Println(myCar.say("car2"))

}

```





### 工厂方法模式*

又称工厂模式、多态工厂模式和虚拟构造器模式。

通过定义**工厂父类**负责定义创建对象的公共接口，而子类则负责生成具体的对象。

工厂方法模式使得**工厂类不再负责所有产品的生产，**而是**定义所有子类工厂类必须实现的接口，**这样添加新产品时就不需要修改工厂逻辑而是添加新的工厂子类，符合开放封闭原则。



```go
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


//////////////////////
hello, i am productA
hello, i am productB
```

![image-20221109233627976](https://test-1309023885.cos.ap-guangzhou.myqcloud.com/typora/image-20221109233627976.png)





### 抽象工厂模式*







### 建造者模式

通过主管类来管理各个步骤，将具体的建造者作为参数传递给构造函数，构造函数再将参数传递给对应的主管类，最后由主管类完成后续建造任务。

建造者模式隐藏了复合对象的创建过程，不同的创建者builder有着不同的创建方法。

建造者模式解决的问题是**构建和组装的解耦**，用户无需关注**复杂对象的创建过程**，只需要指定复杂对象的类型就可以得到该对象。

<img src="https://test-1309023885.cos.ap-guangzhou.myqcloud.com/typora/image-20221110135429259.png" alt="image-20221110135429259" style="zoom:80%;" />

```go
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

```







### 单例模式singleton*

**解决的问题：**整个运行时域，一个类只有一个实例对象，并且该对象的功能依旧能被其他模块使用。

内部生成的单例对外部私有，只有通过对外部暴露的get方法才能获取该实例。



由于有的类比较庞大，频繁的销毁和创建将会造成不必要的性能浪费。（比如数据库链接对象）因此需要单例模式，在系统中只存在一个可控对象，从而节约系统资源。



优点:单例模式提供了对唯一实例的受控访问。在系统内存中只存在一个对象，节约了系统资源。

缺点：单例类的职责过重，拓展略难









> 饿汉式：

初始化单例唯一指针时，就已经提前提前开辟好对象申请了内存。

**好处**：不会出现多线程并发创建，导致多个单例的出现。

**缺点**：无论该单例对象是否被使用，都会创建该单例对象。

```go
package main

//饿汉式单例 无论使用或者不使用，单例都会创建出来
import "fmt"

type singleton struct {
}

func (s *singleton) say() {
	fmt.Println("这里是单例的方法")
}

//对外部私有，只有对外暴露的get方法能够获取该对象
var instance *singleton = new(singleton)

//对外提供一个方法获取这个对象
//去掉了写权限，只暴露读方法
func GetInstance() *singleton {
	return instance
}

func main() {
	i := GetInstance()
	i.say()
}

```



> 懒汉式：

通过get方法获取单例对象，对get方法加锁从而避免并发下多次创建单例对象，第一次使用get方法会开辟对象并申请内存，之后使用将直接返回单例对象。

```go
package main

import (
	"fmt"
	"sync"
)

//定义一个锁
//var lock sync.Mutex
var once sync.Once

var instance *singleton

type singleton struct {
}

//once 是线程安全的
func GetInstance() *singleton {
	//once.Do() 内的函数只能执行一次
	//只有第一次才会执行创建单例
	once.Do(func() {
		instance = new(singleton)
	})
	//之后的都会直接返回单例
	return instance
}


func (s *singleton) Say() {
	fmt.Println("hello")
}

func main() {
	in := GetInstance()
	in.Say()
}
```



Once.Do()方法的源代码：

```go
func (o *Once) Do(f func()) {　　　//判断是否执行过该方法，如果执行过则不执行
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    // Slow-path.
    o.m.Lock()
    defer o.m.Unlock()　　
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
```









## 结构型

### 代理模式

代理模式为某个目标对象提供一个代理对象，并且由代理对象控制对目标对象的访问，代理模式用于**延迟处理操作或者在进行实际操作前后进行其它处理**。

> 作用：

1. 代理模式在客户端和目标对象之间起到中介的作用和保护目标对象的作用
2. 代理对象可以**拓展目标对象的作用**，只需要修改代理类不需要修改目标对象，符合开闭原则。并且如果需要修改目标对象，因为实现了接口，不需要修改代理类，同样符合开闭原则。
3. 代理模式可以**将客户端与目标对象分离，降低了系统耦合**





抽象主题类：真实主题和代理主题的共同接口

真实主题类：代理对象所代表的真实对象

proxy 代理类：含有真实主题的引用，代理角色通常在客户端调用传递给真实主题对象之前或者之后执行某些操作，可以访问、控制、拓展真实主题的功能。

![image-20221112141138797](https://test-1309023885.cos.ap-guangzhou.myqcloud.com/typora/image-20221112141138797.png)

```go
package main

import "fmt"

//代理模式

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

//海外代理
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
    //将具体的类转换为代理类
	var k_proxy Shopping
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


////////////////////
对: 韩国面膜 进行了辨别真伪。
去韩国购物，买了： 韩国面膜
对 韩国面膜  进行了海关检查

对: 四级证书 进行了辨别真伪。
发现假货: 四级证书  ,不应该购买

对: 韩国面膜 进行了辨别真伪。
去美国购物，买了： 韩国面膜
对 韩国面膜  进行了海关检查
```



### 装饰器

装饰器模式关注于在一个对象上动态的添加方法，然而代理模式关注于控制对对象的访问。

当使用代理模式的时候，我们常常在一个代理类中创建一个对象的实例。并且，当我们使用装饰器模式的时候，我们通常的做法是将原始对象作为一个参数传给装饰者的构造器。

优点：

(1) 对于扩展一个对象的功能，装饰模式比继承更加灵活性，不会导致类的个数急剧增加。

(2) 可以通过一种动态的方式来扩展一个对象的功能，从而实现不同的行为。

(3) 可以对一个对象进行多次装饰。

(4) 具体构件类与具体装饰类可以独立变化，用户可以根据需要增加新的具体构件类和具体装饰类，原有类库代码无须改变，符合“开闭原则”。

缺点：

(1) 使用装饰模式进行系统设计时将产生很多小对象，大量小对象的产生势必会占用更多的系统资源，影响程序的性能。

(2) 装饰模式提供了一种比继承更加灵活机动的解决方案，但同时也意味着比继承更加易于出错，排错也很困难，对于多次装饰的对象，调试时寻找错误可能需要逐级排查，较为繁琐。

```go
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

```













### 适配器模式







### 外观模式*