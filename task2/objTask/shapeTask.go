package main

import "fmt"

//定义一个形状的接口
type Shape interface {
	//面积的方法
	Area() float64
	//周长的方法
	Perimeter() float64
}

//长方形的对象 就是结构体
type Rectangle struct {
	//长和宽
	length, width float64
}

//圆形的结构体
type Circle struct {
	//半径
	radii float64
}

// Circle 实现 Shape 接口 面积
func (c Circle) Area() float64 {
	return 3.14 * c.radii * c.radii
}

// Circle 实现 Shape 接口 周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radii
}

// Rectangle 实现 Shape 接口 面积 长方形 面积
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Rectangle 实现 Shape 接口 周长
func (r Rectangle) Perimeter() float64 {
	return r.length*2 + r.width*2
}
func main() {
	//创建长方形对象
	r := Rectangle{10, 20}
	//创建圆形对象 半径为5
	c := Circle{5}
	var r1 Shape = r
	var c1 Shape = c

	//调用长方形的面积和周长方法
	fmt.Println("长方形的面积为：", r1.Area())
	fmt.Println("长方形的周长为：", r1.Perimeter())

	//调用圆形的面积和周长方法
	fmt.Println("圆形的面积为：", c1.Area())
	fmt.Println("圆形的的周长为：", c1.Perimeter())

}
