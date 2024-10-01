package handler

import (
	"encoding/json"
	"net/http"
)

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
	Publisher   string   `json:"publisher"`
	Platforms   []string `json:"platforms"`
	ReleaseDate string   `json:"releaseDate"`
	Price       Price    `json:"price"`
}

type Price struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

var games = []Game{
	{
		Title:       "The Legend of Zelda: Breath of the Wild",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/c/c6/The_Legend_of_Zelda_Breath_of_the_Wild.jpg",
		Rating:      10,
		Description: "An open-world adventure with Link in the land of Hyrule.",
		Details: Details{
			Developer:   "Nintendo",
			Publisher:   "Nintendo",
			Platforms:   []string{"Nintendo Switch", "Wii U"},
			ReleaseDate: "2017-03-03",
			Price: Price{
				Amount:   59.99,
				Currency: "USD",
			},
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
			Publisher:   "CD Projekt",
			Platforms:   []string{"PlayStation 4", "Xbox One", "Nintendo Switch", "PC"},
			ReleaseDate: "2015-05-19",
			Price: Price{
				Amount:   39.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Red Dead Redemption 2",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/4/44/Red_Dead_Redemption_II.jpg",
		Rating:      10,
		Description: "An immersive western experience following Arthur Morgan's story.",
		Details: Details{
			Developer:   "Rockstar Games",
			Publisher:   "Rockstar Games",
			Platforms:   []string{"PlayStation 4", "Xbox One", "PC"},
			ReleaseDate: "2018-10-26",
			Price: Price{
				Amount:   59.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Super Mario Odyssey",
		Genre:       "Platform",
		Image:       "https://upload.wikimedia.org/wikipedia/en/8/8d/Super_Mario_Odyssey.jpg",
		Rating:      10,
		Description: "Mario's globe-trotting adventure to rescue Princess Peach.",
		Details: Details{
			Developer:   "Nintendo",
			Publisher:   "Nintendo",
			Platforms:   []string{"Nintendo Switch"},
			ReleaseDate: "2017-10-27",
			Price: Price{
				Amount:   59.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "God of War",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/a/a7/God_of_War_4_cover.jpg",
		Rating:      10,
		Description: "Kratos and Atreus' journey through Norse mythology.",
		Details: Details{
			Developer:   "Santa Monica Studio",
			Publisher:   "Sony Interactive Entertainment",
			Platforms:   []string{"PlayStation 4"},
			ReleaseDate: "2018-04-20",
			Price: Price{
				Amount:   49.99,
				Currency: "USD",
			},
		},
	},
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
