// Package greeting prints greeting messages to standard output
// excercises on fmt package
package greeting

import "fmt"

var mystring string

type guy struct {
	name string
	age  int
}

// gettingOld is the method of guy
// predicts the person's age after the years
func (s guy) gettingOld(years int) (age int) {
	age = s.age + years
	fmt.Printf("%v will be %v at that time\n", s.name, age)
	return age
}

// Hello prints three messages
func Hello(name string, age int) guy {
	fmt.Println("hello, world!")
	fmt.Printf("great to meet you %v!\n", name)
	return guy{name, age}
}

// MeetHim give you more info of the guy you just met
func MeetHim(guy guy, years int) {
	_ = guy.gettingOld(years)
}
