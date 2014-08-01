package main

type Game struct {
	Name    string              `json:"name"`
	Players map[string]string   `json:"players"`
}

var Games = map[string]Game{"1234": {"Test game", map[string]string{"Bob": "0", "Mary": "1"}}}
