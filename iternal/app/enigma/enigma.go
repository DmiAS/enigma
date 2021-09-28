package enigma

import (
	"github.com/DmiAS/iternal/app/config"
	"github.com/DmiAS/iternal/app/gear"
)

type pipeFunc func(letter int) int
type Enigma struct {
	rotter1   *gear.Gear
	rotter2   *gear.Gear
	rotter3   *gear.Gear
	reflector *gear.Gear
	triggers [2]int
}

func NewEnigma(cfg *config.Config) *Enigma {
	m := new(Enigma)
	m.rotter1 = gear.NewGear(cfg.Mappers[0], cfg.InitLetters[0])
	m.rotter2 = gear.NewGear(cfg.Mappers[1], cfg.InitLetters[1])
	m.rotter3 = gear.NewGear(cfg.Mappers[2], cfg.InitLetters[2])
	m.reflector = gear.NewGear(cfg.Mappers[3], 0)
	m.triggers[0] = cfg.SpinTriggers[0]
	m.triggers[1] = cfg.SpinTriggers[1]
	return m
}
