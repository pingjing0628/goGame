package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func random_num(length int) []int32 {
	a := []int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := make([]int32, length, length) //make array integer length
	if length > 10 || length < 0 {
		fmt.Println("invalid length:", length)
		os.Exit(-1)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < length; i++ {
		tmp := r.Intn(len(a) - i)
		// fmt.Println(r.Intn(len(a) - i))
		b[i] = a[tmp]
		a[tmp], a[len(a)-1-i] = a[len(a)-1-i], a[tmp]
	}
	return b
}

func check_num(a, b []int32) bool {
	var aa, bb int
	if len(a) != len(b) {
		return false
	}
	dict := make(map[int32]int)
	for i := 0; i < len(a); i++ {
		dict[a[i]] = 0
	}
	for i := 0; i < len(a); i++ {
		if _, ok := dict[b[i]]; ok {
			bb++
		}
	}
	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			aa++
		}
	}
	bb = bb - aa
	fmt.Printf("%dA%dB\n", aa, bb)
	if aa == len(a) {
		fmt.Println("OK!")
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Welcome to the Final Code Game!")
	// fmt.Println("Game Start!! Please enter your guess.")
	// var guess int64
	// fmt.Scan(&guess)

	var input int32
	var i int
	status := false
	length := 4
	src := random_num(length)
	tmp := make([]int32, length, length)
	for j := 0; j < 15 && !status; j++ {
		fmt.Printf("Please input %d numbers\n", length)
		fmt.Scanf("%d", &input)
		for i = 0; input > 0; i++ {
			tmp[i] = input % 10
			input = input / 10
		}
		if i < length {
			tmp[i] = 0
		}
		dest := make([]int32, length, length)
		for i := 0; i < len(tmp); i++ {
			dest[length-1-i] = tmp[i]
		}
		status = check_num(src, dest)
	}
}
