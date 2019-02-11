package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretagent struct {
	person
	ltk bool
}

func (p person) Speak() {
	fmt.Println("Hello my name is", p.fname, p.lname)
}

func (s secretagent) Speak() {
	fmt.Printf("Hello my name is %v %v and I have a liscense to kill? Tell me is that true or false: %v", s.fname, s.lname, s.ltk)
}

type human interface {
	Speak()
}

func main() {
	v := person{
		"Robbie",
		"Heeger",
	}
	s := secretagent{
		v,
		true,
	}

	fmt.Println(v.lname)
	v.Speak()
	fmt.Println(s.ltk)
	s.Speak()
	s.person.Speak()

}
