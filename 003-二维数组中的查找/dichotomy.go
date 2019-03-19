package main

import "fmt"

func find(target int, array [][]int) bool {
	x := len(array) - 1
	y := 0
	for x >= 0 && y < len(array[0]) {
		if target == array[x][y] {
			return true
		} else if target > array[x][y] {
			y++
		} else {
			x--
		}
	}
	return false
}

func main() {
	a := [][]int{
		{1, 2, 8, 9},
		{2, 4, 9, 12},
		{4, 7, 10, 13},
		{6, 8, 11, 15},
	}
	fmt.Println(find(3, a))
}
