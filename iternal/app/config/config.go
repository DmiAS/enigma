package config

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

const (
	AlphSize     = 256
	GearsCount   = 4
	RottersCount = 3
	Spinners     = 2

	infoFirst  = "# Конфигурация роттера"
	infoSecond = "# Начальные значения роттера"
	infoThird  = `# При какой букве крутить роттер, например, если Spin2 = "r", значит когда первый роттер дойдет до этой
# буквы второй тоже повернется`

	firstMapper     = "RotterFirstMap"
	secondMapper    = "RotterSecondMap"
	thirdMapper     = "RotterThirdMap"
	reflectorMapper = "ReflectorMap"

	firstInit  = "RotterFirstInit"
	secondInit = "RotterSecondInit"
	thirdInit  = "RotterThirdInit"

	spinnerSecond = "Spin2"
	spinnerThird  = "Spin3"
)

type Letters = [AlphSize]int

type Config struct {
	Mappers [GearsCount]Letters
	// Начальные буквы в роттерах
	InitLetters [RottersCount]int
	// Триггеры прокрутки
	SpinTriggers [Spinners]int
}

func NewConfig(configName string) (*Config, error) {
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, err
	}

	return createConfig(), nil
}

func createConfig() *Config {
	cfg := new(Config)
	cfg.Mappers[0] = createMapper(viper.GetString("RotterFirstMap"))
	cfg.Mappers[1] = createMapper(viper.GetString("RotterSecondMap"))
	cfg.Mappers[2] = createMapper(viper.GetString("RotterThirdMap"))
	cfg.Mappers[3] = createMapper(viper.GetString("ReflectorMap"))

	cfg.InitLetters[0] = getInit(viper.GetString("RotterFirstInit"))
	cfg.InitLetters[1] = getInit(viper.GetString("RotterSecondInit"))
	cfg.InitLetters[2] = getInit(viper.GetString("RotterThirdInit"))

	cfg.SpinTriggers[0] = getInit(viper.GetString("Spin2"))
	cfg.SpinTriggers[1] = getInit(viper.GetString("Spin3"))
	return cfg
}

func GenerateConfigFile(fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := bufio.NewWriter(file)
	defer buf.Flush()

	mappers, err := genMappers()
	if err != nil {
		return err
	}
	if _, err := buf.Write(mappers); err != nil {
		return err
	}

	values, err := genInitValues()
	if err != nil {
		return err
	}
	if _, err := buf.Write(values); err != nil {
		return err
	}

	triggers, err := genTriggers()
	if err != nil {
		return err
	}
	if _, err := buf.Write(triggers); err != nil {
		return err
	}
	return nil
}

func genMappers() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteString(infoFirst + "\n")
	first, err := generateRotterMappers()
	if err != nil {
		return nil, err
	}

	second, err := generateRotterMappers()
	if err != nil {
		return nil, err
	}

	third, err := generateRotterMappers()
	if err != nil {
		return nil, err
	}

	reflector, err := generateRotterMappers()
	if err != nil {
		return nil, err
	}
	buf.WriteString(fmt.Sprintf("%s=\"%s\"\n", firstMapper, string(first)))
	buf.WriteString(fmt.Sprintf("%s=\"%s\"\n", secondMapper, string(second)))
	buf.WriteString(fmt.Sprintf("%s=\"%s\"\n", thirdMapper, string(third)))
	buf.WriteString(fmt.Sprintf("%s=\"%s\"\n", reflectorMapper, string(reflector)))
	return buf.Bytes(), nil

}

func generateRotterMappers() ([]byte, error) {
	shift := rand.Intn(AlphSize)
	b := new(bytes.Buffer)
	for i := 0; i < AlphSize; i++ {
		char := strconv.Itoa((i + shift) % AlphSize)
		if _, err := b.WriteString(char); err != nil {
			return nil, err
		}
		if i < AlphSize-1 {
			if err := b.WriteByte('-'); err != nil {
				return nil, err
			}
		}
	}
	return b.Bytes(), nil
}

func genInitValues() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteString(infoSecond + "\n")
	values := [...]int{
		rand.Intn(AlphSize),
		rand.Intn(AlphSize),
		rand.Intn(AlphSize),
	}
	buf.WriteString(fmt.Sprintf("%s=%d\n", firstInit, values[0]))
	buf.WriteString(fmt.Sprintf("%s=%d\n", secondInit, values[0]))
	buf.WriteString(fmt.Sprintf("%s=%d\n", thirdInit, values[0]))
	return buf.Bytes(), nil
}

func genTriggers() ([]byte, error) {
	buf := new(bytes.Buffer)
	buf.WriteString(infoThird + "\n")
	vals := [...]int{
		rand.Intn(AlphSize),
		rand.Intn(AlphSize),
	}
	buf.WriteString(fmt.Sprintf("%s=%d\n", spinnerSecond, vals[0]))
	buf.WriteString(fmt.Sprintf("%s=%d\n", spinnerThird, vals[1]))
	return buf.Bytes(), nil
}

func getInit(str string) int {
	char := str[0]
	return int(char - 'a')
}
func createMapper(str string) Letters {
	var l Letters
	for i := range str {
		l[i] = int(str[i] - 'a')
	}
	return l
}
