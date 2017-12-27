package main

type qsintarray []int

func (q qsintarray) nnQuickSortInt() {
	q.nnQuickSortIntWorker(0, len(q) - 1)
}

func (q qsintarray) nnQuickSortIntWorker(start int, end int) {
	if start >= end {
		return
	}

	midIndex := (end+start) / 2
	midValue := q[midIndex]
	west, east := start, end

	for west <= east {
		for q[west] < midValue {
			west++
		}

		for q[east] > midValue {
			east--
		}

		if west <= east {
			q[west], q[east] = q[east], q[west]
			west++
			east--
		}
	}

	if west < end {
		q.nnQuickSortIntWorker(west, end)
	}

	if east > start {
		q.nnQuickSortIntWorker(start, east)
	}
}
