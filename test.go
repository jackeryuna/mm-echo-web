package main

import (
	"fmt"
	"math"
)

const g string = "const hello world"

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 9, 12
}
func main() {
	fmt.Println("Hello World")

	var a string = "hello world2"

	var b, c int = 1, 2

	var d = true

	var e int

	f := "hello world3"

	fmt.Println(a, b, c, d, e, f, g, math.Sin(200))

	h := make([]string, 3)

	h[0] = "hello world4"
	h[1] = "hello world5"
	h[2] = "hello world6"

	h = append(h, "hello world7")
	fmt.Println(h)

	i := make(map[string]string)

	i["a"] = "hello world8"
	i["b"] = "hello world9"

	fmt.Println(i)

	j := map[string]string{"a": "hello world10", "b": "hello world11"}

	for k, v := range j {
		fmt.Println(k, v)
	}

	fmt.Println(plus(1, 2))

	_, v := vals()
	fmt.Println(v)
}