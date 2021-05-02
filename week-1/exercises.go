package weekone

import (
	"fmt"
	"strconv"
)

func XenDit(i int) string {
		// TODO: write down the logic with if statement
		if i < 0 {
			return "error, number must not be negative."
		}

		if i == 0 {
			return "0"
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

func Test() {
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
        }, {
					name: "Should return Error",
					in:   -1,
					want: "error, number must not be negative.",
				}, {
					name: "Should return number 0",
					in:   0,
					want: "0",
				},
		}

		for _, tt := range tests {
        result := XenDit(tt.in)
        if result != tt.want {
            fmt.Println("Task #1 - Failed! ", tt.name)
            panic(1)
        }
		}
		
		fmt.Println("Task #1 - Passed!")
}