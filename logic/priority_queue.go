package logic

import (
	"container/heap"
)

type MyNum struct {
	Val   int
	Count int
}

type MyNums []MyNum

func (n MyNums) Len() int {
	return len(n)
}

func (n MyNums) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n MyNums) Less(i, j int) bool {
	return n[i].Count >= n[j].Count
}

func (n *MyNums) Push(num interface{}) {
	myNum := num.(MyNum)
	*n = append(*n, myNum)
}

func (n *MyNums) Pop() interface{} {
	tmp := *n
	l := len(tmp)
	var res interface{} = tmp[l-1]
	*n = tmp[:l-1]
	return res
}

var myNums MyNums

//Sumary Push and data from priority queue(min heap)
//Description When create parking lot that time we will insert all,bcoz all are vacant and on basis of pop we will assign that slot
// to park the car, when we will free the car, that time again insert in to heap
func InsertAllSlotWhileCreateingParkingLot(totalCount int) {
	i := 0
	count := totalCount
	myNums = make(MyNums, totalCount)
	for j := 1; j <= totalCount; j++ {
		myNums[i] = MyNum{Val: j, Count: count}
		count--
		i++
	}
	heap.Init(&myNums)
	// for i := 0; i < totalCount; i++ {
	// 	num := heap.Pop(&myNums).(MyNum)
	// 	fmt.Println(num.Val)
	// }
}
