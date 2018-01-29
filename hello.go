package main

import "fmt"

func getHello() string {
	return "Hello"
}

func getWorld() string {
  return "World"
}

func main() {
	fmt.Println(getHello())
      fmt.Println(getWorld())
}
