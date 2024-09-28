package main

import "fmt"

func main() {

	fmt.Println("Hello")
	pipeline := NewPipelineBuilder().
		WithSimpleStep(step{Name: "Step 1 "}).
		WithSimpleStep(step{Name: "Step 2' "}).
		WithSimpleStep(step{Name: "Step 2'' "}).
		WithSimpleStep(step{Name: "Step 2''' "}).
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

func (pb PipelineBuilder) Build() *Pipeline {
	return &Pipeline{Steps: pb.Steps}
}
