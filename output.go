package main

import (
	"fmt"
	"strings"
)

func Output(str string, args []string) {
	if len(args) == 0 {
		fmt.Println(str)
	} else {
		as := strings.Join(args, " ")
		index := strings.Index(as, "{}")

		if index > 0 {
			as = as[:index] + str + as[index+2:]
			args = strings.Split(as, " ")
		} else {
			args = append(args, str)
		}

		fmt.Println(strings.Join(args, " "))
	}
}
