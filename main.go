package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	stdout := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(stdin)
	for scanner.Scan() {
		unquoted, err := strconv.Unquote(scanner.Text())
		if err != nil {
			unquoted = scanner.Text()
		}
		_, err = io.WriteString(stdout, unquoted)
		if err != nil {
			panic(err)
		}
		err = stdout.Flush()
		if err != nil {
			panic(err)
		}
	}
}
