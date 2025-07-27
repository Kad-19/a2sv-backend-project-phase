package usecase

import (
	"context"
	"time"
	"task_manager/Domain")

type taskUsecase struct {
	taskRepository Domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository Domain.TaskRepository, timeout time.Duration) Domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tu *taskUsecase) Create(c context.Context, task *Domain.Task, userId string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Create(ctx, task, userId)
}

func (tu *taskUsecase) FetchByUserID(c context.Context, userID string) ([]Domain.Task, error) {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.FetchByUserID(ctx, userID)
}

func (tu *taskUsecase) Update(c context.Context, task *Domain.Task) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Update(ctx, task)
}

func (tu *taskUsecase) Delete(c context.Context, id string) error {
	ctx, cancel := context.WithTimeout(c, tu.contextTimeout)
	defer cancel()
	return tu.taskRepository.Delete(ctx, id)
}
