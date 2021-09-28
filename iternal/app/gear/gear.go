package gear

import "github.com/DmiAS/iternal/app/config"

type Gear struct {
	mapper  config.Letters
	current int // текущая буква в ротере, представляет собой индекс в массиве
}

func NewGear(mapper config.Letters, initLetter int) *Gear {
	return &Gear{mapper: mapper, current: initLetter}
}

func (r *Gear) Spin() {
	r.current = (r.current + 1) % config.AlphSize
}

func (r *Gear) GetLetter() int {
	return r.current
}

func (r *Gear) MapReversed(index int) int {
	for i, char := range r.mapper {
		if index == char {
			return i
		}
	}
	return 0
}

func (r *Gear) Map(index int) int {
	return r.mapper[index]
}
