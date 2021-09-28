package main

import (
	"flag"

	"github.com/DmiAS/command"
)

func main() {
	var cfgPath, fileIn, fileOut string
	var shouldGen bool
	flag.StringVar(&cfgPath, "config_path", "C:\\fourth_course\\ib\\labs\\second_lab\\config\\config.toml",
		"путь к файлу, из которого загружается конфигурация")
	flag.BoolVar(&shouldGen, "gen", false,
		"путь к файлу, в котором будет сгенерирована конфигурация")
	flag.StringVar(&fileIn, "file_in", "C:\\fourth_course\\ib\\labs\\second_lab\\files\\hello.txt",
		"путь к файлу, который будет обработан")
	flag.StringVar(&fileOut, "file_out", "C:\\fourth_course\\ib\\labs\\second_lab\\files\\hello_dec.txt",
		"путь к файлу, в котором будет сохранен результат")
	flag.Parse()

	command.Run(cfgPath, shouldGen, fileIn, fileOut)
}
