package models

type OsuMessage struct {
	State struct {
		Name string `json:"name"`
	} `json:"state"`
	ResultsScreen ResultsScreen `json:"resultsScreen"`
	Beatmap       Beatmap       `json:"beatmap"`
}

type ResultsScreen struct {
	Mode struct {
		Name string `json:"name"`
	} `json:"mode"`
	Accuracy float64 `json:"accuracy"`
	Mods     struct {
		Name string `json:"name"`
	} `json:"mods"`
	Rank string `json:"rank"`
	PP   struct {
		Current float64 `json:"current"`
	} `json:"pp"`
	CreatedAt string `json:"createdAt"`
}

type Beatmap struct {
	ID      int    `json:"id"`
	Artist  string `json:"artist"`
	Title   string `json:"title"`
	Version string `json:"version"`
	Mapper  string `json:"mapper"`
}
