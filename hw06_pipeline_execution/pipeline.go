package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	if in == nil || len(stages) == 0 {
		ch := make(Bi)
		close(ch)

		return ch
	}

	out := stageDone(in, done)

	for _, s := range stages {
		out = s(stageDone(out, done))
	}

	return out
}

func stageDone(in, done In) Out {
	out := make(Bi)

	go func() {
		defer close(out)

		for {
			select {
			case <-done:
				return
			case item, ok := <-in:
				if !ok {
					return
				}

				out <- item
			}
		}
	}()

	return out
}
