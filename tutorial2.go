package main

import (
	"fmt"
	"bufio"
	"strconv"
	"os"
)


func main() {
	//Reading input
	Scanner := bufio.NewScanner(os.Stdin)
	Scanner.Scan()
	inp := Scanner.Text()
	fmt.Println(inp)
	//Conversion from string to int
	no, _ := strconv.ParseInt(inp, 10, 64)
	fmt.Println(1000 - no)
}