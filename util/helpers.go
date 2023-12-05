package util

import (
	"fmt"
	"strconv"
)

func MustAtoi(v string) int {
	num, err := strconv.Atoi(v)
	if err != nil {
		panic(fmt.Errorf("parse int %s: %w", v, err))
	}
	return num
}
