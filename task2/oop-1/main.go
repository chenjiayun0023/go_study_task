package main

import (
	"fmt"
	"math"
)

//面向对象

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	rectangle := Rectangle{width: 5.0, height: 3.0}
	fmt.Printf("rectangle的面积: %.2f, 周长：%.2f \n", rectangle.Area(), rectangle.Perimeter())
	var shape Shape = rectangle
	fmt.Printf("rectangle的面积: %.2f, 周长：%.2f \n", shape.Area(), shape.Perimeter())
	var shape1 Shape = &rectangle
	fmt.Printf("rectangle的面积: %.2f, 周长：%.2f \n", shape1.Area(), shape1.Perimeter())

	circle := Circle{radius: 2.0}
	fmt.Printf("circle的面积: %.2f, 周长：%.2f \n", circle.Area(), circle.Perimeter())
	var shape2 Shape = &circle
	fmt.Printf("circle的面积: %.2f, 周长：%.2f \n", shape2.Area(), shape2.Perimeter())
}
