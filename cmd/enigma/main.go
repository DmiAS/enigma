package main

import (
	"fmt"

	"github.com/DmiAS/iternal/app/config"
)

func main() {
	err := config.GenerateConfigFile("C:\\fourth_course\\ib\\labs\\second_lab\\config\\config.toml")
	fmt.Println(err)
	cfg, err := config.NewConfig("config")
	fmt.Println(cfg.Mappers[0], err)
	//flag.Parse()
	//args := flag.Args()
	//command.Run(args)
}
