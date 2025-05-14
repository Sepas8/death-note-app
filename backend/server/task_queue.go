package server

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Sepas8/death-note-app/backend/models"
)

type TaskQueue struct {
	tasks map[int]context.CancelFunc
	mu    sync.Mutex
}

func (tq *TaskQueue) StartTask(id int, duration time.Duration, task func(k *models.Kill) error, k *models.Kill) {
	tq.mu.Lock()
	defer tq.mu.Unlock()

	// Inicializar el mapa si es nil
	if tq.tasks == nil {
		tq.tasks = make(map[int]context.CancelFunc)
	}

	// Cancelar tarea existente si hay una
	if cancel, exists := tq.tasks[id]; exists {
		cancel()
	}

	ctx, cancel := context.WithCancel(context.Background())
	tq.tasks[id] = cancel

	go func() {
		select {
		case <-ctx.Done():
			fmt.Printf("Tarea cancelada para persona %d\n", id)
		case <-time.After(duration):
			fmt.Printf("Ejecutando muerte para persona %d\n", id)
			if err := task(k); err != nil {
				fmt.Printf("Error ejecutando muerte: %v\n", err)
			}

			tq.mu.Lock()
			delete(tq.tasks, id)
			tq.mu.Unlock()
		}
	}()
}
