package main

import "fmt"

type Stack struct {
	capacity, size int
	nums []int
	
}

func (this *Stack) Push(num int) bool {
	if this.capacity == this.size {
		return false
	}

	this.nums[this.size] = num
	this.size++

	return true
}

func (this *Stack) Pop() (int, bool) {
	if this.size == 0 {
		return 0, false
	}

	this.size--
	return this.nums[this.size], true
}

type Tower struct {
	id int
	st Stack
}

func (this *Tower) moveDisks(n int, dest, buff *Tower) {
	if n <= 0 {
		return
	}

	this.moveDisks(n-1, buff, dest)

	this.moveTopDisk(dest)
	printTowers(this, dest, buff)

	buff.moveDisks(n-1, dest, this)
}

func (this *Tower) moveTopDisk(dest *Tower) {
	topDisk, _ := this.st.Pop()
	dest.st.Push(topDisk)
}

func (this *Tower) initialize(id, disks int) {
	this.id = id
	this.st.nums = make([]int, disks)
	this.st.capacity = disks
	this.st.size = 0
}

func (this *Tower) stackDisks() {
	for i := this.st.capacity; i > 0; i-- {
		this.st.Push(i)
	}
}

func printTowers(towers... *Tower) {
	for start, end := 0, len(towers); start != end; start++ {
		for i := 0; i < end; i++ {
			if towers[i].id == start {
				fmt.Print(towers[i].st, " ")
				break
			}
		}
	}
	fmt.Println()
}

func main() {
	numTowers, numDisks := 3, 5

	towers := make([]Tower, numTowers)
	for i := 0; i < numTowers; i++ {
		towers[i].initialize(i, numDisks)
	}

	towers[0].stackDisks()
	printTowers(&towers[0], &towers[1], &towers[2])
	towers[0].moveDisks(numDisks, &towers[1], &towers[2])
}
