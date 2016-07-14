package main

import "fmt"

type fb struct {
	Name         string
	Count, Every uint
}

func NewFB(name string, every uint) fb {
	return fb{
		Name:  name,
		Count: every,
		Every: every,
	}
}

func (f *fb) Run() bool {
	f.Count--
	if f.Count == 0 {
		f.Count = f.Every
		return true
	}
	return false
}

func main() {
	fbs := [...]fb{
		NewFB("Fizz", 3),
		NewFB("Buzz", 5),
	}
	for i := 1; i < 1000; i++ {
		number := true
		for n := range fbs {
			if fbs[n].Run() {
				fmt.Print(fbs[n].Name)
				number = false
			}
		}
		if number {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
