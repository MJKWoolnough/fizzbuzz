package main // import "vimagination.zapto.org/fizzbuzz"

type fb struct {
	Name         string
	Count, Every uint
}

func newFB(name string, every uint) fb {
	return fb{
		Name:  name,
		Count: every,
		Every: every,
	}
}

func (f *fb) Next() bool {
	f.Count--
	if f.Count == 0 {
		f.Count = f.Every
		return true
	}
	return false
}

func main() {
	fbs := [...]fb{
		newFB("Fizz", 3),
		newFB("Buzz", 5),
	}
	for i := 1; i < 1000; i++ {
		number := true
		for n := range fbs {
			if fbs[n].Next() {
				print(fbs[n].Name)
				number = false
			}
		}
		if number {
			print(i)
		}
		print("\n")
	}
}
