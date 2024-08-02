package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	serverURL = "http://localhost:8080/cotacao"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	cotacao, err := obterCotacao(ctx)
	if err != nil {
		log.Fatalf("Erro ao obter cotação: %v", err)
	}

	if err := salvarCotacao(cotacao.Bid); err != nil {
		log.Fatalf("Erro ao salvar cotação: %v", err)
	}

	fmt.Println("Cotação salva com sucesso.")
}

func obterCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", serverURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var cotacao Cotacao
	if err := json.NewDecoder(resp.Body).Decode(&cotacao); err != nil {
		return nil, err
	}

	return &cotacao, nil
}

func salvarCotacao(bid string) error {
	data := fmt.Sprintf("Dólar: %s\n", bid)
	return os.WriteFile("cotacao.txt", []byte(data), 0644)
}
