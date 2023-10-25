package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.google.commm")
	if err != nil {
		log.Fatal(err.Error(), "Ocorreu um erro na requisição")
	}

	fmt.Printf("O tamanho do conteudo retornado foi de %d bytes", res.ContentLength)
}
