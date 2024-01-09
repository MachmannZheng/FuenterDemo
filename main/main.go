package main

import (
	"fmt"

	syncT "github.com/MachmannZheng/FuenterDemo/syncTest"
)

func main() {
	fmt.Println("-- start the main function --")
	syncT.SynctestRaceCondition()()
}
