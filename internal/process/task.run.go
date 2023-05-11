package process

import (
	"context"
	"github.com/lishimeng/track/internal/task"
)

func RunTask(_ context.Context) (err error) {

	for _, t := range task.Tasks {
		go t.Run()
	}

	return
}
