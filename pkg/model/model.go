package model

type Message interface{}

type Stage interface {
	Process(message Message) ([]Message, error)
}

type PipelineOpts struct {
	Concurrency int
}

type Pipeline interface {
	AddPipe(pipe Stage, opt *PipelineOpts)
	Start() error
	Stop() error
	Input() chan<- Message
	Output() <-chan Message
}
