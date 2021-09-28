package enigma

import (
	"github.com/DmiAS/iternal/app/config"
)

func getLetterIndex(letter byte) int {
	return int(letter)
}

func indexToChar(index int) byte {
	return byte(index)
}

func keepRing(a int) int {
	return (config.AlphSize + a) % config.AlphSize
}

func (e *Enigma) Process(ch chan byte) {
	inLeft := make(chan int)
	outLeft := e.startProcessLeft(inLeft)
	inRight := make(chan int)
	outRight := e.startProcessRight(inRight)

	for char := range ch {
		//e.spin()
		// буква проходит справа налево
		inLeft <- getLetterIndex(char)
		resL := <-outLeft
		// отражаем через рефлектор
		resRef := e.processReflector(resL)
		// буква проходит слева направо
		inRight <- resRef
		resR := <-outRight
		// формируем результирующую строку
		ch <- indexToChar(resR)
	}
}

func (e *Enigma) startProcessLeft(in chan int) chan int {
	funcs := []pipeFunc{
		e.processFirstL,
		e.processSecondL,
		e.processThirdL,
	}
	out := startPipelines(in, funcs)
	return out
}

func (e *Enigma) startProcessRight(in chan int) chan int {
	funcs := []pipeFunc{
		e.processThirdR,
		e.processSecondR,
		e.processFirstR,
	}
	out := startPipelines(in, funcs)
	return out
}

func (e *Enigma) processFirstL(letter int) int {
	curr := e.rotter1.GetLetter()
	sum := keepRing(curr + letter)
	return e.rotter1.Map(sum)
}
func (e *Enigma) processFirstR(letter int) int {
	mapped := e.rotter1.MapReversed(letter)
	curr := e.rotter1.GetLetter()
	newLetter := keepRing(mapped - curr)
	return newLetter
}

func (e *Enigma) processSecondL(letter int) int {
	curr := e.rotter2.GetLetter()
	currRotter1 := e.rotter1.GetLetter()
	diff := keepRing(curr - currRotter1)
	sum := keepRing(letter + diff)
	return e.rotter2.Map(sum)
}
func (e *Enigma) processSecondR(letter int) int {
	mapped := e.rotter2.MapReversed(letter)
	curr := e.rotter2.GetLetter()
	currRotter1 := e.rotter1.GetLetter()
	diff := keepRing(curr - currRotter1)
	newLetter := keepRing(mapped - diff)
	return newLetter
}

func (e *Enigma) processThirdL(letter int) int {
	curr := e.rotter3.GetLetter()
	currRotter2 := e.rotter2.GetLetter()
	diff := keepRing(curr - currRotter2)
	sum := keepRing(letter + diff)
	return e.rotter3.Map(sum)
}

func (e *Enigma) processThirdR(letter int) int {
	mapped := e.rotter3.MapReversed(letter)
	curr := e.rotter3.GetLetter()
	currRotter2 := e.rotter2.GetLetter()
	diff := keepRing(curr - currRotter2)
	newLetter := keepRing(mapped - diff)
	return newLetter
}

func (e Enigma) processReflector(letter int) int {
	currRotter3 := e.rotter3.GetLetter()
	// буква только пришла на рефлектор
	diff := keepRing(letter - currRotter3)
	mapped := e.reflector.MapReversed(diff)

	// выводим букву с рефлектора
	res := keepRing(mapped + currRotter3)
	return res
}
