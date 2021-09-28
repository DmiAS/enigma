package enigma


func (e *Enigma) spin(){
	e.rotter1.Spin()
	index := e.rotter1.GetLetter()
	if index == e.triggers[0]{
		e.rotter2.Spin()
		index := e.rotter2.GetLetter()
		if index == e.triggers[1]{
			e.rotter3.Spin()
		}
	}
}
