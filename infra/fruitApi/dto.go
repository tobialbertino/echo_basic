package fruitApi

type DataFruit struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	Family        string `json:"family"`
	Order         string `json:"order"`
	Genus         string `json:"genus"`
	DataNutrition `json:"nutritions"`
}

type DataNutrition struct {
	Calories      int     `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}
