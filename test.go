package main

import "fmt"

// Golang Type Assertions
// эта функция (i... interface{}) работает для любого количества параметров.
func displayValue(i ...interface{}) {
	fmt.Println(i)
}

func main() {

	a := "Hello world"

	b := 90

	c := false

	displayValue(a)
	displayValue(a, b)
	displayValue(a, b, c)

	var p interface{}
	p = 20
	// Здесь (int)проверяется, равно ли значениеаявляется целым числом или нет. Если true, то значение будет присвоеноинтерфейсЗначение.
	//interfaceValue := i.(int)

	interfaceValue1, booleanValue1 := p.(int)
	fmt.Println("Interface Value: ", interfaceValue1)
	fmt.Println("Boolean Value: ", booleanValue1)

	var k interface{}
	k = "Hello World"
	interfaceValue2, booleanValue2 := k.(int)
	fmt.Println("Interface Value: ", interfaceValue2)
	fmt.Println("Boolean Value: ", booleanValue2)

	for i := 0; i < 5; i++ {
		result := 20 / i
		fmt.Println(result)
	}

}
