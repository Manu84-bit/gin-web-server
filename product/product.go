package product

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"nombre"`
	Price      float64 `json:"precio"`
	Stock      int     `json:"stock"`
	Code       string  `json:"código"`
	Expiration string  `json:"vencimiento"`
	Published  bool    `json:"publicado"`
}
