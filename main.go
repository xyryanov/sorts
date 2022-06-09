package main

import (
	"fmt"
	"math/rand"
	"time"
	"00.id.lv/sorts/bubble"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr []int

	for i := 0; i < 1000; i++ {
		n := rand.Intn(1000)
		arr = append(arr, n)
	}


	bubble.Sort(arr)	
	fmt.Println(arr)
}


