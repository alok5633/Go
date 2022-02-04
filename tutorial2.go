package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)


func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	inp := Scanner.Text()
	fmt.Println(inp)
	no, _ := strconv.ParseInt(inp, 10, 64)
	fmt.Println(1000 - no)
}