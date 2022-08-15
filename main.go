package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"tictactoe-golang/AI"
)

var board map[string]string

type Move struct {
	IsTerminated bool   `json:"is_terminated"`
	Winner       string `json:"winner"`
	Player       string `json:"player"`
	Move         int    `json:"move"`
}

func updateBoard(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}

	board = map[string]string{}

	err = json.Unmarshal(body, &board)
	if err != nil {
		w.WriteHeader(500)
		log.Fatalln(err)
	}

	w.WriteHeader(200)
}

func getMove(w http.ResponseWriter, r *http.Request) {
	var res Move
	if AI.Terminal(board) {
		res.Move = -1
		if win := AI.Winner(board); win != "" {
			res.Winner = win
		} else {
			res.Winner = "draw"
		}
		res.IsTerminated = true

	} else {
		res.Move = int(AI.Minimax(board)[1] - '0')
		res.Player = AI.Player(board)

		if AI.Terminal(AI.Result(board, fmt.Sprintf("c%v", res.Move))) {
			res.IsTerminated = true
			if win := AI.Winner(AI.Result(board, fmt.Sprintf("c%v", res.Move))); win != "" {
				res.Winner = win
			} else {
				res.Winner = "draw"
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)

	err := enc.Encode(res)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./UI/")))
	http.HandleFunc("/updateBoard", updateBoard)
	http.HandleFunc("/getMove", getMove)

	fmt.Println("Game started at http://localhost:8080/")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
