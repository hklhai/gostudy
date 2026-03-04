package main

import "fmt"

type S struct {
	a string
}

func (s *S) set(b string) {
	s.a = b
}

func (s S) setNoPointer(b string) {
	s.a = b
}

func main() {
	s := S{}
	s.set("1")
	fmt.Println(s.a)
	s.setNoPointer("2")
	fmt.Println(s.a)

	ss := &S{}
	ss.set("3")
	fmt.Println(ss.a)
	ss.setNoPointer("4")
	fmt.Println(ss.a)
}
