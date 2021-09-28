package command

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/DmiAS/iternal/app/config"
	"github.com/DmiAS/iternal/app/enigma"
)

func Run(cfgPath string, genPath bool, fileIn, fileOut string) {
	if genPath {
		if err := config.GenerateConfigFile(cfgPath); err != nil {
			log.Fatalln("невозможно сконфигурировать файл - ", err)
		}
	}

	cfg, err := config.NewConfig(cfgPath)
	if err != nil {
		log.Fatalf("ошибка при создании конфигурации = %s", err.Error())
	}

	m := enigma.NewEnigma(cfg)
	if err := processFile(fileIn, fileOut, m); err != nil {
		log.Fatalln("невозможно обработать файл - ", err)
	}
}

func processFile(in, out string, machine *enigma.Enigma) error {
	file, err := os.OpenFile(in, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	encryptedFile, err := os.OpenFile(out, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	defer encryptedFile.Close()

	ch := make(chan byte)
	defer close(ch)
	go machine.Process(ch)

	buf := bufio.NewReader(file)
	encBuf := bufio.NewWriter(encryptedFile)
	defer encBuf.Flush()
	for {
		char, err := buf.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		ch <- char
		encrypted := <-ch
		if err := encBuf.WriteByte(encrypted); err != nil {
			return err
		}
	}
	return nil
}
