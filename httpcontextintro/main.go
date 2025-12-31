package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)

	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		println("Excedeu o tempo limite antes de fazer a requisição")
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Requisição bem sucedida")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
