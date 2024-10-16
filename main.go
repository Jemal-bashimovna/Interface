package main

import "fmt"

func main() {
	person1.say()
	person2.say()

	person1.work()
	person2.work()

	person1.think()
	person2.think()
	fmt.Println("Interface topManager:")
	person3.say()
	person3.think()
}

type worker interface {
	say()
	work()
	think()
}

type topManager interface {
	say()
	think()
}
type empFront struct{}

func (e empFront) say() {
	fmt.Println("Method say() of Frontend employee")
}
func (e empFront) work() {
	fmt.Println("Method work() of Frontend employee")
}
func (e empFront) think() {
	fmt.Println("Method think() of Frontend employee")
}

type empBack struct{}

func (e empBack) say() {
	fmt.Println("Method say() of Backend employee")
}
func (e empBack) work() {
	fmt.Println("Method work() of Backend employee")
}
func (e empBack) think() {
	fmt.Println("Method think() of Backend employee")
}

var person1 worker = empFront{}

var person2 worker = empBack{}

var person3 topManager = empFront{}
