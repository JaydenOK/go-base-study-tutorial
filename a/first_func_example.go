package a

import (
	"fmt"
)

type student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

// 类似于 php array_filter(array $input, $callback = null, $flag = 0) { }
func filter(s []student, f func(student) bool) []student {
	var r []student
	for _, v := range s {
		if f(v) == true {
			r = append(r, v)
		}
	}
	return r
}

func main() {
	s1 := student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}
	s := []student{s1, s2}
	f := filter(s, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

	f2 := filter(s, func(s student) bool {
		if s.grade == "BA" {
			return true
		}
		return false
	})
	fmt.Println(f2)
}
