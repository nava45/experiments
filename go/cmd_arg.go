package main

import (
	"fmt";
	"os";
)

type Profile struct {
	Name string
	Age int
}

type Marks struct {
	Tamil int
	English int
	Maths int
	Science int
	SST int
}

func (marks *Marks) average() int {
	return (marks.Tamil + marks.English + marks.Maths + marks.Science + marks.SST) / 5
	
}

func (prof *Profile) PrintDetails() {
	fmt.Println("Name: %s, Age: %d", prof.Name, prof.Age)
	prof.Age += 2
}

func add(a int, b int) int {
	return a+b
}

func main(){
	mark := 100
	fmt.Println("OS Input",  os.Args[1], mark)
	mark = 99;
	fmt.Println("New Mark:", mark)
	res := add(mark, mark)
	fmt.Println("Add Result: %d", res)
	nava := &Profile{Name: "Navaneethan", Age: 27}
	nava.PrintDetails()
	Super(nava)
	fmt.Println("My Profile Structure is:", nava)

	m := &Marks{Tamil:90, English:50, Maths:99, Science:96, SST:90}
	fmt.Println(m)
	avg := m.average()
	fmt.Println("Total Marks", avg)
}


func Super(mys *Profile){
	//mys.Age = 20
	fmt.Println("Inside Super", mys)
	mys = &Profile{Name:"Saravana", Age:20}
	fmt.Println("Inside Super", mys)
}
