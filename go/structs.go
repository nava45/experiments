package main

import (
	"fmt"
)

type person struct{
	name string
	age int
}

//Variadic function example
func calcTotal(nums ...int) int {
	fmt.Print(nums, " ")
	total := 0
	for _,v := range nums{
		total += v
	}
	return total
}

//closures
func closExam(i int) func() int {
	return func() int {
		i += 1
		return i
	}
}

func main(){
	s1 := person{name:"Nava", age:20}
	s1.name = "python"
	fmt.Println(s1.name)

	fmt.Println(calcTotal(1,2,3,4))
	fmt.Println(calcTotal([]int{1,2,3,4}...))

	clf := closExam(1)
	fmt.Println(clf())
	fmt.Println(clf())
	fmt.Println(clf())
	fmt.Println(clf())
}
