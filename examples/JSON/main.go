package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct { // (1)!
	Page   int // (2)!
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"` // (3)!
	Fruits []string `json:"fruits"`
}

func main() {

	bolB, _ := json.Marshal(true) // (4)!
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"} // (5)!
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &response1{ // (6)!
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`) // (7)!

	var dat map[string]interface{} // (8)!

	if err := json.Unmarshal(byt, &dat); err != nil { // (9)!
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // (10)!
	fmt.Println(num)

	strs := dat["strs"].([]interface{}) // (11)!
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}` // (12)!
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout) // (13)!
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
