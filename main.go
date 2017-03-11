package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LINEONE   = "__________"
	LINETWO   = "|x _|  o |"
	LINETHREE = "| |   O  |"
	LINEFOUR  = "|_|_o__0_|"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(LINEONE)
	fmt.Println(LINETWO)
	fmt.Println(LINETHREE)
	fmt.Println(LINEFOUR)
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	return
}
