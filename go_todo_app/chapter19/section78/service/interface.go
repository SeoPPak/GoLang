package service

import (
	"context"
	"section78/entity"
	"section78/store"
)

//go:generate moq -out task_adder_mock.go . TaskAdder TaskLister
type TaskAdder interface {
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface {
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}
