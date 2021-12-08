package main

import "fmt"

type studentinfo interface {
	progress() (string, string)
}

type student1 struct {
	name   string
	report string
}
type student2 struct {
	name   string
	report string
}

func main() {
	rollno1 := student1{name: "rani", report: "good"}
	rollno2 := student2{name: "rasmi", report: "bad"}
	fmt.Println(rollno1.progress())
	fmt.Println(rollno2.progress())
}

func (s1 student1) progress() (string, string) {
	return s1.name, s1.report
}

func (s2 student2) progress() (string, string) {
	return s2.name, s2.report
}
