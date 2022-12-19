package model

type Bunker struct {
	ID       string `json:"id"`
	NAME     string `json:"name"`
	CATEGORY string `json:"category"`
	PLAQUE   string `json:"plaque"`
}

var InitialData = []Bunker{
	{
		ID:       "1",
		NAME:     "God of War",
		CATEGORY: "Open World",
		PLAQUE:   "34 ZPA 45",
	},
	{
		ID:       "2",
		NAME:     "World War",
		CATEGORY: "Fantasy",
		PLAQUE:   "59 KPL 45",
	},
	{
		ID:       "3",
		NAME:     "BattleFiled",
		CATEGORY: "Shooter",
		PLAQUE:   "59 KMT 67",
	},
}
