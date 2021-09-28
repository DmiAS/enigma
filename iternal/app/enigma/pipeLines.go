package enigma

func startPipelines(in chan int, funcs []pipeFunc) chan int{
	for _, f := range funcs{
		in = func(in chan int, f pipeFunc)chan int{
			out := make(chan int)
			go startPipe(in, out, f)
			return out
		}(in, f)
	}
	return in
}

func startPipe(in, out chan int, f pipeFunc){
	for char := range in{
		out <- f(char)
	}
	close(out)
}
