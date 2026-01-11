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
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) error {
	select {
	case <-ctx.Done():
		fmt.Println("Reserva de hotel cancelada ou expirou o tempo limite:", ctx.Err())
		return ctx.Err()

	case <-time.After(5 * time.Second):
		resp, err := http.Get("https://www.bookingholdings.com/modern-slavery-statement/")
		if err != nil {
			fmt.Println("Erro ao buscar URL:", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Erro ao ler resposta:", err)
		}
		println(string(body))
		fmt.Println("Hotel reservado com sucesso!")
		return nil
	}
}
