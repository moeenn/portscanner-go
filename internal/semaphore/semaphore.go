package semaphore

type unit struct{}

type Semaphore struct {
	c chan unit
}

func New(maxWorkers int) *Semaphore {
	return &Semaphore{
		c: make(chan unit, maxWorkers),
	}
}

func (s *Semaphore) Acquire() {
	s.c <- unit{}
}

func (s *Semaphore) Release() {
	<-s.c
}
