package main

import "fmt"

func demo2() {
	var name string
	var age int
	n, err := fmt.Scanf("%s - %d", &name, &age)
	fmt.Println(n, err, name, age)
}
