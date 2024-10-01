package handler

import (
	"encoding/json"
	"net/http"
)

// Define the Game struct
type Game struct {
	Title       string  `json:"title"`
	Genre       string  `json:"genre"`
	Image       string  `json:"image"`
	Rating      int     `json:"rating"`
	Description string  `json:"description"`
	Details     Details `json:"details"`
}

type Details struct {
	Developer   string   `json:"developer"`
	Platforms   []string `json:"platforms"`
	ReleaseDate string   `json:"releaseDate"`
}

// Create a list of games
var games = []Game{
	{
		Title:       "The Legend of Zelda: Breath of the Wild",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/c/c6/The_Legend_of_Zelda_Breath_of_the_Wild.jpg",
		Rating:      10,
		Description: "An open-world adventure with Link in the land of Hyrule.",
		Details: Details{
			Developer:   "Nintendo",
			Platforms:   []string{"Nintendo Switch", "Wii U"},
			ReleaseDate: "2017-03-03",
		},
	},
	{
		Title:       "The Witcher 3: Wild Hunt",
		Genre:       "Action RPG",
		Image:       "https://upload.wikimedia.org/wikipedia/en/0/0c/Witcher_3_cover_art.jpg",
		Rating:      10,
		Description: "Geralt's epic journey in a rich fantasy world.",
		Details: Details{
			Developer:   "CD Projekt Red",
			Platforms:   []string{"PlayStation 4", "Xbox One", "Nintendo Switch", "PC"},
			ReleaseDate: "2015-05-19",
		},
	},
	// Add more games here...
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gamesJSON, err := json.Marshal(games)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(gamesJSON)
}
