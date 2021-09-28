package command

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/DmiAS/iternal/app/config"
	"github.com/DmiAS/iternal/app/enigma"
)

func Run(args []string) {
	if len(args) < 0 {
		log.Fatalln("переданы не все аргументы")
	}
	configName := args[0]
	fileIn := args[1]
	fileOut := args[2]

	cfg, err := config.NewConfig(configName)
	if err != nil {
		log.Fatalf("оишбка при создании конфигурации = %s", err.Error())
	}
	m := enigma.NewEnigma(cfg)

	if err := proccessFile(fileIn, fileOut, m); err != nil {
		log.Fatalln("невозможно обработать файл - ", err)
	}
}

func proccessFile(in, out string, machine *enigma.Enigma) error {
	file, err := os.OpenFile(in, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	encryptedFile, err := os.OpenFile(out, os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	ch := make(chan byte)
	defer close(ch)
	machine.Process(ch)

	buf := bufio.NewReader(file)
	encBuf := bufio.NewWriter(encryptedFile)
	for char, err := buf.ReadByte(); ; {
		if err != nil && err != io.EOF {
			return err
		}
		ch <- char
		encrypted := <-ch
		if err := encBuf.WriteByte(encrypted); err != nil {
			return err
		}
	}
	return nil
}
