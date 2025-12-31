package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.WriteString("Olá, Mundo!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso! : Tamanho: %d bytes", tamanho)

	f.Close()

	/// leitura do arquivo
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("\nConteúdo do arquivo:", string(arquivo))

	/// leitura de pouco em pouco, abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 12)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Printf("Lido %d bytes: %s\n", n, string(buffer[:n]))
	}
}
