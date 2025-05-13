package server

import (
	"backend/models"
	"context"
	"fmt"
	"sync"
	"time"
)

type TaskQueue struct {
	mu    sync.Mutex
	tasks map[int]context.CancelFunc
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks: make(map[int]context.CancelFunc),
	}
}

func (tq *TaskQueue) StartTask(id int, duration time.Duration, task func(k *models.Kill) error, k *models.Kill) {
	ctx, cancel := context.WithCancel(context.Background())

	tq.mu.Lock()
	tq.tasks[id] = cancel
	tq.mu.Unlock()

	go func() {
		defer func() {
			tq.mu.Lock()
			delete(tq.tasks, id)
			tq.mu.Unlock()
		}()

		select {
		case <-ctx.Done():
			fmt.Printf("La tarea con ID %d fue cancelada.\n", id)
		case <-time.After(duration):
			fmt.Printf("Iniciando tarea asíncrona con id %d...\n", id)
			err := task(k)
			if err != nil {
				fmt.Printf("Error en tarea asíncrona: %v\n", err)
			}
			fmt.Printf("La tarea con ID %d fue completada tras %v.\n", id, duration)
		}
	}()
}

func (tq *TaskQueue) CancelTask(id int) bool {
	tq.mu.Lock()
	cancel, exists := tq.tasks[id]
	tq.mu.Unlock()

	if exists {
		cancel()
		return true
	}
	return false
}
