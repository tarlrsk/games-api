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
	{
		Title:       "Minecraft",
		Genre:       "Sandbox",
		Image:       "https://upload.wikimedia.org/wikipedia/en/5/51/Minecraft_cover.png",
		Rating:      9,
		Description: "A blocky, creative sandbox game with endless possibilities.",
		Details: Details{
			Developer:   "Mojang Studios",
			Publisher:   "Mojang Studios",
			Platforms:   []string{"PC", "PlayStation", "Xbox", "Nintendo Switch", "Mobile"},
			ReleaseDate: "2011-11-18",
			Price: Price{
				Amount:   26.95,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Fortnite",
		Genre:       "Battle Royale",
		Image:       "https://assets.goal.com/v3/assets/bltcc7a7ffd2fbf71f5/bltdaa08bc17301576f/60dc27ce2e95e10f21f3193c/8e3054435647a98c29bd131600e394a887e22841.png",
		Rating:      8,
		Description: "A popular battle royale game with building mechanics.",
		Details: Details{
			Developer:   "Epic Games",
			Publisher:   "Epic Games",
			Platforms:   []string{"PC", "PlayStation", "Xbox", "Nintendo Switch", "Mobile"},
			ReleaseDate: "2017-07-25",
			Price: Price{
				Amount:   0.00,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Overwatch",
		Genre:       "First-person shooter",
		Image:       "https://upload.wikimedia.org/wikipedia/commons/thumb/5/55/Overwatch_circle_logo.svg/500px-Overwatch_circle_logo.svg.png",
		Rating:      9,
		Description: "A team-based multiplayer shooter with diverse characters.",
		Details: Details{
			Developer:   "Blizzard Entertainment",
			Publisher:   "Blizzard Entertainment",
			Platforms:   []string{"PC", "PlayStation", "Xbox", "Nintendo Switch"},
			ReleaseDate: "2016-05-24",
			Price: Price{
				Amount:   39.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Grand Theft Auto V",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/a/a5/Grand_Theft_Auto_V.png",
		Rating:      10,
		Description: "A satirical take on modern society set in Los Santos.",
		Details: Details{
			Developer:   "Rockstar North",
			Publisher:   "Rockstar Games",
			Platforms:   []string{"PlayStation 4", "Xbox One", "PC"},
			ReleaseDate: "2013-09-17",
			Price: Price{
				Amount:   29.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Dark Souls III",
		Genre:       "Action RPG",
		Image:       "https://upload.wikimedia.org/wikipedia/en/b/bb/Dark_souls_3_cover_art.jpg",
		Rating:      9,
		Description: "A challenging journey through a dark and atmospheric world.",
		Details: Details{
			Developer:   "FromSoftware",
			Publisher:   "Bandai Namco Entertainment",
			Platforms:   []string{"PlayStation 4", "Xbox One", "PC"},
			ReleaseDate: "2016-03-24",
			Price: Price{
				Amount:   59.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Sekiro: Shadows Die Twice",
		Genre:       "Action-adventure",
		Image:       "https://upload.wikimedia.org/wikipedia/en/6/6e/Sekiro_art.jpg",
		Rating:      9,
		Description: "A shinobi's quest for revenge in Sengoku period Japan.",
		Details: Details{
			Developer:   "FromSoftware",
			Publisher:   "Activision",
			Platforms:   []string{"PlayStation 4", "Xbox One", "PC"},
			ReleaseDate: "2019-03-22",
			Price: Price{
				Amount:   59.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Hades",
		Genre:       "Roguelike dungeon crawler",
		Image:       "https://upload.wikimedia.org/wikipedia/en/6/67/Hades_cover_art.jpg",
		Rating:      10,
		Description: "An action-packed rogue-like where you escape the Underworld.",
		Details: Details{
			Developer:   "Supergiant Games",
			Publisher:   "Supergiant Games",
			Platforms:   []string{"PC", "Nintendo Switch", "PlayStation", "Xbox"},
			ReleaseDate: "2020-09-17",
			Price: Price{
				Amount:   24.99,
				Currency: "USD",
			},
		},
	},
	{
		Title:       "Celeste",
		Genre:       "Platform",
		Image:       "https://upload.wikimedia.org/wikipedia/en/9/9c/Celeste_cover.png",
		Rating:      10,
		Description: "A platformer about climbing a mountain and overcoming challenges.",
		Details: Details{
			Developer:   "Maddy Makes Games",
			Publisher:   "Maddy Makes Games",
			Platforms:   []string{"PC", "PlayStation", "Xbox", "Nintendo Switch"},
			ReleaseDate: "2018-01-25",
			Price: Price{
				Amount:   19.99,
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
