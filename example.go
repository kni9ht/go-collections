package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kni9ht/go-collections/Collection/HashSet"
	List "github.com/kni9ht/go-collections/Collection/Linklist"
)

func getRandom() int32 {
	num := int32(rand.Int())
	if num < 0 {
		num *= -1
	}
	return num
}

func main() {
	t := time.Now()
	mp := HashSet.New()
	for i := 0; i < 10000000; i++ {
		mp.Push(getRandom())
	}
	fmt.Printf("Time required for inserting elements in Map : %d Sec \n", int(time.Since(t).Seconds()))

	t = time.Now()
	for i := 0; i < 100000; i++ {
		mp.Find(getRandom())
	}
	fmt.Printf("Time required for searching elements in Map : %d Sec \n", int(time.Since(t).Seconds()))

	t = time.Now()
	ls := List.New()
	for i := 0; i < 10000000; i++ {
		ls.Push(getRandom())
	}
	fmt.Printf("Time required for inserting elements in List : %d Sec \n", int(time.Since(t).Seconds()))

	t = time.Now()
	for i := 0; i < 100000; i++ {
		ls.Search(getRandom())
	}
	fmt.Printf("Time required for searching elements in List : %d Sec \n", int(time.Since(t).Seconds()))
}
