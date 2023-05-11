package task

type Task interface {
	Run()
}

var Tasks []Task

func Register(t ...Task) {
	Tasks = append(Tasks, t...)
}
