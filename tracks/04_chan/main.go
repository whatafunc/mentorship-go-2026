package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Основные настройки
var (
	ctxTimeout = 3 * time.Second       // таймаут всей обработки
	taskDelay  = 10 * time.Millisecond // имитация времени обработки задачи
)

func main() {
	// Создаем контекст с таймаутом и не забываем отмену
	ctx, cancel := context.WithTimeout(context.Background(), ctxTimeout)
	defer cancel()

	tasks := make(chan int32, 10)   // входной канал задач
	results := make(chan int32, 10) // канал результатов
	var completeResult int32        // финальный результат (через atomic)

	// Обработка задач: читаем из tasks, пишем в results
	go func(ctx context.Context, tasks <-chan int32, results chan<- int32) {
		defer close(results) // закрываем канал, когда всё обработали
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context canceled:", ctx.Err())
				return
			case task, ok := <-tasks:
				if !ok {
					return // tasks закрыт — выходим
				}
				time.Sleep(taskDelay) // имитация работы
				results <- task       // отправляем результат
			}
		}
	}(ctx, tasks, results)

	// Отправка задач: запускаем 100 горутин, каждая пишет в канал
	go func(tasks chan<- int32) {
		var wg sync.WaitGroup
		for i := 0; i < 100; i++ {
			wg.Add(1)
			i := i // захватываем переменную, иначе race
			go func() {
				defer wg.Done()
				tasks <- int32(i)
			}()
		}
		wg.Wait()
		close(tasks) // закрываем tasks, когда все задачи отправлены
	}(tasks)

	fmt.Println("ready to read results channel:")
	// Читаем результаты и агрегируем атомарно
	for result := range results {
		atomic.AddInt32(&completeResult, result)
	}

	fmt.Println("Final result:", completeResult)
}
