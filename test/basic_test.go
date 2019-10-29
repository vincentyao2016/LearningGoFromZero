package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strings"
	"testing"
)

// 声明变量
func TestVariables(t *testing.T) {
	var a string = "Runoob"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// 没有初始化就为零值
	var defaultInt int
	fmt.Println(defaultInt)

	// bool 零值为 false
	var defaultBool bool
	fmt.Println(defaultBool)
}

// 常量
const length int = 10

func TestConst(t *testing.T) {
	// 隐式声明
	const WIDTH = 5
	var area int
	const a, b, c = 1, false, "str" //多重赋值

	area = length * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	println(a, b, c)
}

// Operators 运算符
func TestOperators(t *testing.T) {
	//Arithmetic Operators 算术运算符
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)

	// Logical Operators 逻辑运算符
	var e bool = true
	var f bool = false
	if e && f {
		fmt.Printf("第一行 - 条件为 true\n")
	}
	if e || f {
		fmt.Printf("第二行 - 条件为 true\n")
	}
	/* 修改 a 和 b 的值 */
	e = false
	f = true
	if e && f {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(e && f) {
		fmt.Printf("第四行 - 条件为 true\n")
	}
}

// Decision Making 条件语句
func TestDecisionMaking(t *testing.T) {
	/* 局部变量定义 */
	a := 100

	// if 判断布尔表达式
	if a < 20 {
		/* 如果条件为 true 则执行以下语句 */
		fmt.Printf("a 小于 20\n")
	} else {
		/* 如果条件为 false 则执行以下语句 */
		fmt.Printf("a 不小于 20\n")
	}
	fmt.Printf("a 的值为 : %d\n", a)

	// switch 定义局部变量
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}

	fmt.Printf("你的等级是 %s\n", grade)
}

// Loop
func TestForLoop(t *testing.T) {
	// for loop
	b := 15
	var a int

	numbers := [6]int{1, 2, 3, 5}

	/* for 循环 */
	for a := 0; a < 10; a++ {
		fmt.Printf("a 的值为: %d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("a 的值为: %d\n", a)
	}

	for i, x := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

	// nested loop
	/* local variable definition */
	var i, j int

	for i = 2; i < 100; i++ {
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 {
				break // if factor found, not prime
			}
		}
		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

// func
/* function returning the max between two numbers */
func swap(x, y string) (string, string) {
	return y, x
}

func TestFunc(t *testing.T) {
	a, b := swap("Google", "Runoob")
	fmt.Println(a, b)
}

// String
func TestString(t *testing.T) {
	// String length
	var greeting = "Hello world!"

	fmt.Printf("String Length is: ")
	fmt.Println(len(greeting))

	// Concatenating Strings
	greetings := []string{"Hello", "world!"}
	fmt.Println(strings.Join(greetings, " "))
}

// Arrays
func TestArrays(t *testing.T) {
	// var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	// var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	var n [10]int /* n is an array of 10 integers */
	var i, j int

	/* initialize elements of array n to 0 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* set element at location i to i + 100 */
	}

	/* output each array element's value */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

// Pointers
func TestPointers(t *testing.T) {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

// Structures
type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func (b Books) GetName() string {
	fmt.Println("author is", b.author)
	return b.author
}

func TestStruc(t *testing.T) {
	b := Books{"New Book from Vincent", "Vincent", "Program", 1}
	fmt.Println("Test Struc", b.GetName())
}

// JSON convert Struct Adavanced
// https://www.jianshu.com/p/c7d5c1c8a2d8
func TestStructWithTag(t *testing.T) {
	person := Person{"tom", 12}
	if b, err := json.Marshal(person); err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("value: %s", b)
	}

	personWithTags := PersonWithoutTags{"tom", 12}
	if b, err := json.Marshal(personWithTags); err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Printf("value: %s", b)
	}
}

// 如果一个域不是以大写字母开头的，那么转换成json的时候，这个域是被忽略的。
type Person struct {
	Name string `json:"name"`
	age  int    `json:"age"`
}

type PersonWithoutTags struct {
	Name string
	Age  int
}

// Slice
func TestSlice(t *testing.T) {
	/* 创建切片 */
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	/* 打印原始切片 */
	fmt.Println("numbers ==", numbers)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers[1:4] ==", numbers[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers[:3] ==", numbers[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers[4:] ==", numbers[4:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number2 := numbers[:2]
	printSlice(number2)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number3 := numbers[2:5]
	printSlice(number3)

	numbers = []int{}

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 = make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

// Range
func TestRange(t *testing.T) {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)
	//在数组上使用range将传入index和值两个变量。上面那个例子我们不需要使用该元素的序号，所以我们使用空白符"_"省略了。有时侯我们确实需要知道它的索引。
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	//range也可以用在map的键值对上。
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

// Map
// https://medium.com/rungo/the-anatomy-of-maps-in-go-79b82836838b
// https://www.digitalocean.com/community/tutorials/understanding-maps-in-go
func TestMap(t *testing.T) {
	//var countryCapitalMap map[string]string /*创建集合 */
	// Create an empty map
	countryCapitalMap := make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India "] = "新德里"

	/*删除元素*/
	delete(countryCapitalMap, "France")

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap["American"] /*如果确定是真实的,则存在,否则不存在 */
	fmt.Println("检查是否存在：", countryCapitalMap["American"], ok)
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在", capital)
	}

	// Initializing a map
	/** Instead of creating empty map and assigning new data,
	we can create a map with some initial data, like array and slice.
	*/
	age := map[string]int{
		"mina": 28,
		"john": 32,
		"mike": 55,
	}

	fmt.Println(age)

}

// Type Casting
func TestTypeCasting(t *testing.T) {
	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count)
	fmt.Printf("Value of mean : %f\n", mean)
}

// Error Handling
func TestErrorHandling(t *testing.T) {
	result, err := Sqrt(-1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result, err = Sqrt(9)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

func Sqrt(value float64) (float64, error) {
	if value < 0 {
		return 0, errors.New("Math: negative number passed to Sqrt")
	}
	return math.Sqrt(value), nil
}
