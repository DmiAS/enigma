package main

import (
	"flag"

	"github.com/DmiAS/command"
)

func main(){
	flag.Parse()
	args := flag.Args()
	command.Run(args)
}