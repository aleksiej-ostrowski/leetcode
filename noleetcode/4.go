package main

import (
	"fmt"
)

type IVoiceable interface {
	Voice() string
}

type Cat struct {
	say string
}

type Cow struct {
	say string
}

type Dog struct {
	say string
}

func (r Cat) Voice() string {
	r.say = "Мяу"
	return r.say
}

func (r Cow) Voice() string {
	r.say = "Мууу"
	return r.say
}

func (r Dog) Voice() string {
	r.say = "Гав"
	return r.say
}

func main() {
	cat := Cat{}
	dog := Dog{}
	cow := Cow{}

	fmt.Println(cat.Voice()) // Мяу
	fmt.Println(dog.Voice()) // Гав
	fmt.Println(cow.Voice()) // Мууу
}
