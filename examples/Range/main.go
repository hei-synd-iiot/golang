package main

import "fmt"

func main() {
	nums := []int{2, 3, 4} // (1)!
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { // (2)!
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"} // (3)!
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs { // (4)!
		fmt.Println("key:", k)
	}

	for i, c := range "go" { // (5)!
		fmt.Println(i, c)
	}
}
