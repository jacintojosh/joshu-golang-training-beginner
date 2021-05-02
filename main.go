// package main

// import (
// 	weekzero "github.com/jacintojosh/weekzero"
// )

// func main() {
// 	weekzero.Exercise1()
// }

package main

import (
	"fmt"
	"strconv"
)

func xenDit(i int) string {
		// TODO: write down the logic with if statement
		if i < 0 {
			return "error, number must not be negative."
		}
		remainder := i % 3
		if remainder > 0 {
			return strconv.Itoa(i)
		}

		quotient := i / 3

		if quotient % 10 == 0 {
			return "Xendit"
		} else if quotient % 2 == 0 {
			return "Dit"
		} else if quotient % 3 == 0 {
			return "Xen"
		}

		return "Unexpected error"
}

func main() {
    tests := []struct {
        name string
        in   int
        want string
    }{
        {
            name: "Should return number 1",
            in:   1,
            want: "1",
        }, {
            name: "Should return Xen",
            in:   9,
            want: "Xen",
        }, {
            name: "Should return Dit",
            in:   24,
            want: "Dit",
        }, {
            name: "Should return Xendit",
            in:   30,
            want: "Xendit",
        },
		}
		for _, tt := range tests {
        result := xenDit(tt.in)
        if result != tt.want {
            fmt.Println("Task #1 - Failed! ", tt.name)
            panic(1)
        }
		}
		fmt.Println("Task #1 - Passed!")
}