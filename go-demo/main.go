package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

//10 конкурентных запросов
//вывести в консоль 10 статусов

func main() {
	t := time.Now()
	url := "https://google.com"
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			getHttpCode(url)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(t))
}

func getHttpCode(url string) {
	//error
	//url := "https://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка %s при выполнении запроса", err.Error())
	}
	fmt.Printf("Получен ответ %d\n", resp.StatusCode)
}
