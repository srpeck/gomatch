// Gomatch is a RESTful web service built with Goji
package main

import (
	"net/http"
        "math/rand"
        "time"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
        rand.Seed(time.Now().UnixNano())

	static := web.New()
	static.Get("/assets/*", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.Handle("/assets/", static)

	goji.Get("/favicon.ico", http.FileServer(http.Dir("./assets/img/")))
	goji.Get("/games", ServeIndex)
	goji.Get("/game/:id", ServeIndex)
	goji.Get("/", ServeIndex)

	goji.Get("/api/games", GetGames)
	goji.Get("/api/events/games", GetGamesEvents)
	goji.Get("/api/game/:game_id", GetGame)
	goji.Get("/api/events/game/:game_id", GetGameEvents)
	goji.Post("/api/game/:game_name", PostGame)
	goji.Put("/api/game/:game_id/player/:player/vote/:vote", PutGame)
	goji.Delete("/api/game/:game_id/player/:player", DeletePlayer)

	goji.Serve()
}
