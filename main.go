package main

import (
	doublelinklist "algorithmsAndDataStructure/double_link_list"
)

//"algorithmsAndDataStructure/sorting"
//"fmt"

//func main() {
//nums := []int{3, 2, 4, 1, 5, 7}
//ret := sorting.HeapSort(nums)
//fmt.Println(ret)

//}

func main() {
	obj := doublelinklist.Constructor()
	//param_1 := obj.Get(0)

	obj.AddAtHead(1)
	//param_1 = obj.Get(0)
	//fmt.Println(param_1)
	//obj.AddAtTail(2)
	//param_1 = obj.Get(2)
	//fmt.Println(param_1)
	//obj.AddAtIndex(0,3)
	//param_1 = obj.Get(2)
	//fmt.Println(param_1)
	obj.DeleteAtIndex(0)
	//param_1 = obj.Get(0)
	//fmt.Println(param_1)
}
