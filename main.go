package main

import (
	"fmt"
	"math/rand"
)

func randomFillin(arr []int) {
	for index := range arr {
		arr[index] = rand.Int() % 100
	}
}

func findPos(val int, arr []int) int {
	i := 0
	for val < arr[i] {
		i++
	}
	return i
}

func insert(val int, arr []int, length *int) {
	size := *length
	pos := findPos(val, arr)
	for i := 0; i < (size - pos); i++ {
		arr[size-i] = arr[size-i-1]
	}
	arr[pos] = val
	*length++
}

func sort(arr []int) []int {
	result := make([]int, cap(arr), cap(arr))
	length := 0
	for _, val := range arr {
		insert(val, result, &length)
	}
	return result
}

func swap(arr []int, posA int, posB int) {
	tmp := arr[posA]
	arr[posA] = arr[posB]
	arr[posB] = tmp
}

func update(arr []int, pos int, val int) {
	origin := arr[pos]
	arr[pos] = val
	if val > origin {
		for i := pos - 1; i >= 0; i-- {
			if val > arr[i] {
				swap(arr, i, pos)
				pos = i
			} else {
				break
			}
		}
	} else if val < origin {
		for i := pos + 1; i <= len(arr); i++ {
			if val < arr[i] {
				swap(arr, i, pos)
				pos = i
			} else {
				break
			}
		}
	}
}

func main() {
	a := make([]int, 10)
	randomFillin(a)
	b := sort(a)
	fmt.Printf("before update:%v\n", b)
	update(b, 5, 10)
	fmt.Printf("after update:%v\n", b)
}
