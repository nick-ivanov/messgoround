package main

import "testing"

func TestQuickSort(t *testing.T) {
	var a qsintarray
	a = append(a, 1, 3, 2, 10, 11, -16, 15, 2)
	a.nnQuickSortInt()

	if !(a[0] == -16 && a[1] == 1 && a[2] == 2 && a[3] == 2 && a[4] == 3 && a[5] == 10 && a[6] == 11 && a[7] == 15) {
		t.Errorf("Expected sorted sequence to be [-16, 1, 2, 2, 3, 10, 11, 15], but instead got %v", a)
	}
}
