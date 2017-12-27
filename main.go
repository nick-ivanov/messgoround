package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("== Welcome to Mess Go Round ===")

	var arr qsintarray
	fmt.Print("NUMBER OF INTEGER ITEMS: ")

	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Print("ARRAY[" + strconv.Itoa(i+1) + "]: ")
		var value int
		fmt.Scan(&value)
		arr = append(arr, value)
	}

	arr.nnQuickSortInt()

	fmt.Println("=== QUICKSORTED ARRAY ===")
	fmt.Println(arr)
}
