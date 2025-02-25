package model

type Product struct {
	Name      string  `json:"name"`
	Category  string  `json:"category"`
	Size      string  `json:"size"`
	Gender    string  `json:"gender"`
	IsForKids bool    `'json:"isForKids"`
	Price     float32 `json:"price"`
	Quantity  uint16  `json:"quantity"`
}
