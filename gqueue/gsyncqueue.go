package gqueue

type Job interface {
	Run() Job
}

type syncQueue struct {
	Queue chan Job
}

func NewSyncQueue(m int) *syncQueue {
	return &syncQueue{
		Queue: make(chan Job, m),
	}
}
func (p *syncQueue) AddJob(job Job) {
	select {
	case p.Queue <- job:
	}
}

func (p *syncQueue) Run() {
	go func() {
		for task := range p.Queue {
			task.Run()
		}
	}()
}
