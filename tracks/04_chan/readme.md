Key Points About Channels:
Writing to a closed channel causes a panic.

Reading from a closed channel:

-Returns the zero value immediately for unbuffered channels

-Returns remaining values then zero values for buffered channels

Closing a channel:

-Should only be done by the sender

-Should not be done more than once (also causes panic)

Best Practices:

-Use defer close() when you know you're the only sender

-Consider using a sync.Once to ensure channels are only closed once

-Use select with a done channel for graceful termination



Привет, коллеги! 👋  
Сегодня хочу поделиться шпаргалкой по работе с горутинами и каналами в Go. 

🔥 Основные концепции  

1. Горутины — легковесные потоки, которые работают конкурентно. Запускаются через go func().  
2. Каналы — безопасный способ обмена данными между горутинами. Бывают буферизированные и небуферизированные.  
3. Контекст — помогает управлять временем жизни операций (таймауты, отмены).  
4. WaitGroup — синхронизация группы горутин (ждём завершения всех).  
5. Atomic — атомарные операции для безопасного доступа к переменным из разных горутин.  


📌 Аксиомы работы с каналами  

✔️ Закрывать каналы — иначе можно получить deadlock или панику.  
✔️ Не блокировать главную горутину — иначе программа зависнет.  
✔️ Использовать `select` для мультиплексирования — обрабатывать несколько каналов или контекст.  
✔️ Избегать data races — использовать atomic, мьютексы или каналы.  


💡 Моя шпаргалка  

Вот пример кода, который объединяет все эти концепции:  
- Запуск горутин для обработки задач.  
- Использование каналов для передачи данных.  
- Контекст с таймаутом для graceful shutdown.  
- WaitGroup для синхронизации отправки задач.  
- Atomic для безопасного агрегирования результатов.  

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

  // Читаем результаты и агрегируем атомарно
  for result := range results {
    atomic.AddInt32(&completeResult, result)
  }

  fmt.Println("Final result:", completeResult)
}

Go делает конкурентность простой, но требует дисциплины:  
- Всегда думайте о блокировках и утечках горутин.  
- Используйте правильные примитивы синхронизации.  
- Тестируйте на race condition (go run -race).  

Пользуйтесь шпаргалкой, готовьтесь к собеседованиям и пишите надежный код! 💪

