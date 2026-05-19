package main

import "fmt"

type Animal struct {
	Name string
}

type Cat struct {
	Animal
}

func main() {
	animal := Animal{
		Name: "猫",
	}
	cat := Cat{
		Animal: animal,
	}
	fmt.Println(cat)

}
