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

func (p person) pSpeak() {
	fmt.Println("Hello my name is", p.fname, p.lname)
}

func (s secretagent) saSpeak() {
	fmt.Printf("Hello my name is %v %v and I have a liscense to kill? Tell me is that true or false: %v", s.fname, s.lname, s.ltk)
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
	v.pSpeak()
	fmt.Println(s.ltk)
	s.saSpeak()
	s.pSpeak()

}
