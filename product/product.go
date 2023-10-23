package product

type Product struct {
	Id         int     `json:"id"`
	Name       string  `json:"nombre" binding:"required`
	Price      float64 `json:"precio" binding:"required`
	Stock      int     `json:"stock" binding:"required`
	Code       string  `json:"código" binding:"required`
	Expiration string  `json:"vencimiento" binding:"required`
	Published  bool    `json:"publicado" binding:"required`
}
