package main

import (
	"context"
	"fmt"
	"sync"
)

func getSrvData(srv string) (string, error) {
	// например, можно использовать http.Get(srv) или что-то подобное
	return fmt.Sprintf("otvet %s", srv), nil
}

func main() {
	wg := &sync.WaitGroup{} // создаем группу ожидания для горутин
	servers := []string{"server1", "server2", "server3"}
	servers = append(servers, "server4", "server5")
	chanResp := make(chan string, len(servers))                // канал для получения данных от серверов
	ctx, cancelCtx := context.WithCancel(context.Background()) // контекст с таймаутом
	defer cancelCtx()                                          // не забываем отменить контекст, когда он больше не нужен
	for _, server := range servers {
		wg.Add(1) // увеличиваем счетчик горутин на 1
		go func(srv string) {
			defer wg.Done() // уменьшаем счетчик горутин, когда функция завершится
			res, err := getSrvData(srv)
			if err != nil {
				// обработка ошибки, например, логирование
				return
			}
			select {
			case chanResp <- res: // отправляем результат в канал
			case <-ctx.Done():
				return
			}
		}(server)
	}
	wg.Wait()       // ждем, пока все горутины завершатся
	close(chanResp) // закрываем канал, чтобы сигнализировать, что данные получены
	var dat string
	dat = <-chanResp
	println("Data received from server:", dat)
}
