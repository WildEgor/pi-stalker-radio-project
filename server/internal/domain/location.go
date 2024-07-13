package domain

type Location struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	X        int       `json:"x"`
	Y        int       `json:"y"`
	Features []Feature `json:"features"`
}
