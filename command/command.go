package command

import (
	"fmt"
	"log"

	"github.com/DmiAS/iternal/app/config"
	"github.com/DmiAS/iternal/app/enigma"
)

const (
	encrypt = "e"
	decrypt = "d"
)
func Run(args []string) {
	if len(args) < 0{
		log.Fatalln("переданы не все аргументы")
	}
	configName := args[0]
	str := args[1]

	cfg, err := config.NewConfig(configName)
	if err != nil{
		log.Fatalf("оишбка при создании конфигурации = %s", err.Error())
	}
	m := enigma.NewEnigma(cfg)

	in := []byte(str)
	out := string(m.Process(in))
	fmt.Println(out)
}