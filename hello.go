package learn_go_with_tests

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
