package product

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"nombre"`
	Price        float64 `json:"precio"`
	Stock        int     `json:"stock"`
	Code         int     `json:"código"`
	CreationDate string  `json:"fecha_creación"`
	Published    bool    `json:"publicado"`
}
