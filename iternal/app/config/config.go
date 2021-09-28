package config

import (
	"github.com/spf13/viper"
)

const (
	AlphSize     = 26
	GearsCount   = 4
	RottersCount = 3
	Spinners     = 2
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

func createConfig() *Config{
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

func getInit(str string) int{
	char := str[0]
	return int(char - 'a')
}
func createMapper(str string) Letters{
	var l Letters
	for i := range str{
		l[i] = int(str[i] - 'a')
	}
	return l
}
