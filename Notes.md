Go Tour:

一. Basics

<1> Packages, variables and functions
1. 多值返回

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

2. 变量

var c, python, java bool

3. 变量初始化

var c, python, java = true, false, "no!"

c, python, java := true, false, "no!"

3. 类型转换

i := 42
f := float64(i)
u := uint(f)

4. 常量

常量不能使用 := 语法定义

const Pi = 3.14
const STR = "one string"

<2> Flow control statements

1. defer

defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}

输出：
hello
world

2. defer stack

延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

输出：
counting
done
9
8
7
6
5
4
3
2
1
0

<3> More types

1. 指针

Go 具有指针。 指针保存了变量的内存地址。

类型 *T 是指向类型 T 的值的指针。其零值是 nil 。
var p *int

& 符号会生成一个指向其作用对象的指针。
i := 42
p = &i

* 符号表示指针指向的底层的值。
fmt.Println(*p) // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i

package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

输出：
42
21
73

2. struct

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(a.X + a.Y)
	fmt.Println(Vertex{1, 2})
}

输出：
3
{1 2}

3. 结构体指针

结构体字段可以通过结构体指针来访问。

通过指针间接的访问是透明的。

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

输出：
{1000000000 2}

4. 结构体文法

结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 类型为 Vertex
	v2 = Vertex{X: 1}  // Y:0 被省略
	v3 = Vertex{}      // X:0 和 Y:0
	p  = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

输出：
{1 2} &{1 2} {1 0} {0 0}

5. array

类型 [n]T 是一个有 n 个类型为 T 的值的数组。

表达式

var a [10]int

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
}

6. slice

一个 slice 会指向一个序列的值，并且包含了长度信息。

[]T 是一个元素类型为 T 的 slice。

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("s ==", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] == %d\n", i, s[i])
	}
}

7. Slices are like references to arrays

A slice does not store any data, it just describes a section of an underlying array.

Changing the elements of a slice modifies the corresponding elements of its underlying array.

package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

输出：
[John Paul George Ringo]
[John Paul] [Paul George]
[John XXX] [XXX George]
[John XXX George Ringo]

8. 构造slice

slice 由函数 make 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组：

a := make([]int, 5)  // len(a)=5
为了指定容量，可传递第三个参数到 make：

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

输出：
a len=5 cap=5 [0 0 0 0 0]
b len=0 cap=5 []
c len=2 cap=5 [0 0]
d len=3 cap=3 [0 0 0]

9. append to slice

package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	// append works on nil slices.
	a = append(a, 0)
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1)
	printSlice("a", a)

	// we can add more than one element at a time.
	a = append(a, 2, 3, 4)
	printSlice("a", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

输出：
a len=0 cap=0 []
a len=1 cap=2 [0]
a len=2 cap=2 [0 1]
a len=5 cap=8 [0 1 2 3 4]