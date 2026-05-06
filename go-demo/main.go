package main

import (
	"fmt"
	"net/http"
)

//10 конкурентных запросов
//вывести в консоль 10 статусов

func main() {
	var url string = "https://google.com"
	code := make(chan int)
	for i := 0; i < 10; i++ {
		go getHttpCode(url, code)
		res := <-code
		fmt.Printf("Код: %d\n", res)
	}
	for res := range code {
		fmt.Printf("Получен ответ %d\n", res)
	}
}

func getHttpCode(url string, codeCh chan int) {
	//error
	//url := "https://google.com"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Ошибка %s при выполнении запроса", err.Error())
	}
	codeCh <- resp.StatusCode
}
