package main

import (
	"fmt"

	"github.com/DmiAS/iternal/app/config"
)

func main() {
	err := config.GenerateConfigFile("C:\\fourth_course\\ib\\labs\\second_lab\\config\\config.toml")
	fmt.Println(err)
	//flag.Parse()
	//args := flag.Args()
	//command.Run(args)
}
