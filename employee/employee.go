package employee

type Employee struct {
	Id     int    `json:"id"`
	Nombre string `json:"name"`
	Activo bool   `json:"active"`
}