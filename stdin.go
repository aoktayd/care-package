package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func intcodeStdin() []int {
	bytes, err := ioutil.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}

	input := strings.Split(string(bytes), ",")
	intcode := make([]int, 0, len(input))

	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intcode = append(intcode, i)
	}

	return intcode
}
