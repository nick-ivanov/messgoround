package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("== Welcome to Mess Go Round ===")

	fmt.Println("Truly random integer between 0 and 10:", trulyRandomInt(10))

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
