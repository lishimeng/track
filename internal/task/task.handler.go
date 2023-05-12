package task

import "context"

type Task interface {
	Run(ctx context.Context)
}

var Tasks []Task

func Register(t ...Task) {
	Tasks = append(Tasks, t...)
}
