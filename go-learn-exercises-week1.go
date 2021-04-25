package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect" //Package reflect implements run-time reflection, allowing a program to manipulate objects with arbitrary types.
)

func exercise1() {
	// TODO: task #1 - Why is it not working?

	// Only numeric types, strings and bools can be constants.
	// task1 is declared as const, but os.Getenv is a function.
	// Correct way is to either use the short hand way or long hand way to declare task1.
	// shorthand way: task1 := os.Getenv("task_1")
	// longhand way: var task1 string = os.Getenv("task_1")
	var task1 string = os.Getenv("task_1")

	// DeepEqual reports whether x and y are “deeply equal,” deeply equal being defined depending on the types of the params.
	if !reflect.DeepEqual("", task1) {
		fmt.Println("[Exercise 1]: Task #1 failed!")
		return
	}

	// TODO: Challenge #1 - How to convert from any type into string
	var challenge1 interface{} // DO NOT CHANGE THE DATATYPE

	challenge1 = 1

	// Seems like for interfaces, we use something called a Stringer to convert them to strings, these are defined by the fmt package.
	challenge1 = fmt.Sprintf("%v", challenge1)

	if !reflect.DeepEqual("1", challenge1) {
		fmt.Println("[Exercise 1]: Challenge #1 failed!")
		return
	}
	fmt.Println("[Exercise 1]: All passed!")
}

type Domicile struct {
	Country  string `json:"country"`
	IsRemote bool   `json:"is_remote"`
}

type Employee struct {
	Name           string   `json:"name"`
	Entity         string   `json:"entity"`
	EmployeeNumber int      `json:"employee_number"`
	Salary         float64  `json:"salary"`
	Domicile       Domicile `json:"domicile"`
}

func exercise2() {
	// TODO: task #1 - give me a skeleton!
	// Missing comma near salary, fixed it.
	data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5,
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
	var employee Employee

	if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println("[Exercise 2]: Task #1 failed!")
		return
	}

	// TODO: task #2 - I am a legal employee, include me into your database!
	var database map[string]Employee

	// Initialize database first so we can add entries
	database = make(map[string]Employee)

	// Add entry Golang
	database["Golang"] = employee

	if !reflect.DeepEqual(database["Golang"], employee) {
		fmt.Println("[Exercise 2]: Task #2 failed!")
		return
	}

	fmt.Println("[Exercise 2]: All passed!")
}

func main() {
	exercise1()
	exercise2()
}
