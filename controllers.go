package main

import (
        "encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
        "math/rand"
        "time"

	"github.com/zenazn/goji/web"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./assets/index.html")
}
func GetGames(w http.ResponseWriter, r *http.Request) {
        games, _ := json.Marshal(Games)
	io.WriteString(w, string(games))
}
func GetGamesEvents(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

        games, _ := json.Marshal(Games)
        fmt.Fprintf(w, "data: %s\n\n", string(games))
}
func GetGame(c web.C, w http.ResponseWriter, r *http.Request) {
        game, ok := Games[c.URLParams["game_id"]]
	if !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}
        result, _ := json.Marshal(game)
	io.WriteString(w, string(result))
}
func GetGameEvents(c web.C, w http.ResponseWriter, r *http.Request) {
        game, ok := Games[c.URLParams["game_id"]]
	if !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}

        w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

        result, _ := json.Marshal(game)
        fmt.Fprintf(w, "data: %s\n\n", string(result))
}
func PostGame(c web.C, w http.ResponseWriter, r *http.Request) {
        game_id := strconv.Itoa(rand.Int())
        Games[game_id] = Game{Name: string(c.URLParams["game_name"]), Players: map[string]string{}}
        io.WriteString(w, game_id)
}
func PutGame(c web.C, w http.ResponseWriter, r *http.Request) {
	if _, ok := Games[c.URLParams["game_id"]]; !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}
        Games[c.URLParams["game_id"]].Players[string(c.URLParams["player"])] = c.URLParams["vote"]

        result, _ := json.Marshal(Games[c.URLParams["game_id"]])
	io.WriteString(w, string(result))
}
func DeletePlayer(c web.C, w http.ResponseWriter, r *http.Request) {
	if _, ok := Games[c.URLParams["game_id"]]; !ok {
		http.Error(w, http.StatusText(404), 404)
		return
	}
        delete(Games[c.URLParams["game_id"]].Players, string(c.URLParams["player"]))

        time.Sleep(time.Second * 10)
	if len(Games[c.URLParams["game_id"]].Players) == 0 {
                delete(Games, c.URLParams["game_id"])
	}
}
