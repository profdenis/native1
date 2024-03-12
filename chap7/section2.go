package main

import (
	"fmt"
	"native1/chap7/list"
)

func section2a() {
	intList1 := list.NewIntList(7, 1, 8, 2, 7)
	fmt.Println(intList1)
	err := intList1.Insert(6, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intList1)
	err = intList1.Insert(6, 22)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(intList1)

	intList2 := list.NewIntList()
	intList2.AddEnd(5)
	intList2.AddEnd(8)
	intList2.AddEnd(3)
	fmt.Println(intList2)

	for !intList2.IsEmpty() {
		i, err := intList2.RemoveStart()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(i)
			fmt.Println(intList2)
		}
	}
}
