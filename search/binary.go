package main

import "fmt"

func find_index(length uint) uint {
	var i uint

	if (length != 1) {
		if (length % 2 == 0) {
			i = (length)/2 - 1
		} else {
			i = length/2
		}
	} else {
		i = 0
	}
	return i;
}

func find_number(numb int, arr []int, i uint, prev uint) int {
	var ix, length uint
	var result int

	if arr[i] == numb {
		return (1)
	} else if numb < arr[i] {
		arr = arr[:(i + 1)]
	} else {
		arr = arr[(i + 1):]
	}
	length = uint(len(arr))
	ix = find_index(length)
	if (ix != 0) {
		prev = ix
		result = find_number(numb, arr, ix, prev)
	} else if (ix == 0 && length > 1) {
		result = find_number(numb, arr, ix, prev)
	} else {
		return (-1)
	}
	return result
}

func main() {
	var arr = []int{ 2, 5, 8, 12, 16, 23, 38, 56, 72, 91 }
	var length uint = uint(len(arr))
	var i int

	i = find_number(4, arr, find_index(length), 0)
	fmt.Println(i)
}
