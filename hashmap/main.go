package main

import "fmt"

func iterate(student_grades map[string]int) {

	for key, value := range student_grades {
		fmt.Println(key, value)
	}

}
func main() {
	var student_grades map[string]int

	var def int = 9
	student_grades = make(map[string]int)

	student_grades["srinivas"] = 45
	student_grades["basha"] = 56
	student_grades["srinivas"] = 78
	student_grades["raghava"] = 90

	//iterating over keys
	iterate(student_grades)

	// checking for key
	if value, exists := student_grades["primary"]; exists == false {
		student_grades["primary"] = def
	} else {
		fmt.Println(value)
	}

	iterate(student_grades)

	delete(student_grades, "xyz")

}
