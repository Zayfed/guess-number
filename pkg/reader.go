package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadNumber(message string) int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println(message)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("error: input must be an integer")
			continue
		}

		return num
	}
}
