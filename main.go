package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	fmt.Println("Hello")
	pipeline := NewPipelineBuilder().
		WithSimpleStep(step{Name: "Step 1 "}).
		WithSimpleStep(concurrentStep{SubSteps: []Step{
			step{Name: "Step 2' "},
			step{Name: "Step 2'' "},
			step{Name: "Step 2''' "},
		}}).
		WithSimpleStep(step{Name: "Step 3 "}).Build()

	pipeline.Run()

}

type Pipeline struct {
	Steps []Step
}

func (p *Pipeline) Run() {
	for _, step := range p.Steps {
		step.Run()
	}
}

type step struct {
	Name string
}

func (s step) Run() {
	fmt.Printf("Running step %s \n", s.Name)
	time.Sleep(100 * time.Millisecond) // Simulate work being done, allow scheduler to switch

}

type concurrentStep struct {
	Name     string
	SubSteps []Step
}

func (s concurrentStep) Run() {
	fmt.Println("start")

	var wg sync.WaitGroup

	wg.Add(len(s.SubSteps))

	for _, step := range s.SubSteps {
		go func() {
			step.Run()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("end")
}

type Step interface {
	Run()
}

type PipelineBuilder struct {
	Steps []Step
}

func NewPipelineBuilder() PipelineBuilder {
	return PipelineBuilder{}
}

func (pb PipelineBuilder) WithSimpleStep(step Step) PipelineBuilder {
	pb.Steps = append(pb.Steps, step)
	return pb
}

func (pb PipelineBuilder) WithParallelStep(step Step) PipelineBuilder {
	pb.Steps = append(pb.Steps, step)
	return pb
}

func (pb PipelineBuilder) Build() *Pipeline {
	return &Pipeline{Steps: pb.Steps}
}
