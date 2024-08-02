// server.go
package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	apiURL     = "http://economia.awesomeapi.com.br/json/last/USD-BRL"
	dbName     = "cotacoes.db"
	serverPort = ":8080"
)

type Cotacao struct {
	Bid string `json:"bid"`
}

type Response struct {
	USDBRL Cotacao `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/cotacao", cotacaoHandler)
	log.Printf("Servidor rodando na porta %s\n", serverPort)
	log.Fatal(http.ListenAndServe(serverPort, nil))
}

func cotacaoHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 200*time.Millisecond)
	defer cancel()

	cotacao, err := obterCotacao(ctx)
	if err != nil {
		log.Printf("Erro ao obter cotação: %v", err)
		http.Error(w, "Erro ao obter cotação", http.StatusInternalServerError)
		return
	}

	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	if err := salvarCotacao(ctxDB, cotacao.Bid); err != nil {
		log.Printf("Erro ao salvar cotação no banco: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cotacao)
}

func obterCotacao(ctx context.Context) (*Cotacao, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", apiURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}

	return &response.USDBRL, nil
}

func salvarCotacao(ctx context.Context, bid string) error {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return err
	}
	defer db.Close()

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		_, err = db.ExecContext(ctx, "CREATE TABLE IF NOT EXISTS cotacao (id INTEGER PRIMARY KEY, bid TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, "INSERT INTO cotacao (bid) VALUES (?)", bid)
		if err != nil {
			return err
		}
	}
	return nil
}
