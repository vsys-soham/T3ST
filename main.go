package main

import "log"

func add(x, y int) int {
	return x + y
}

func main() {

	log.Println(add(2, 5))
	log.Println(add(-2, 5))
	log.Println(add(2, -5))
	log.Println(add(-2, -5))
}
