package main

import "fmt"

const (
	Weight = 60
	Height = 160
	Age    = 30
)

const (
	count0 = iota
	count1
	count2
	count3
)

func main() {
	fmt.Println(Weight)
	fmt.Println(Height)
	fmt.Println(Age)
	fmt.Println(count0)
	fmt.Println(count1)
	fmt.Println(count2)
	fmt.Println(count3)

}
