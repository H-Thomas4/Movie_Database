package entities

import (
	"github.com/google/uuid"
)

type Movie struct {
	Id          string   `json:"Id"`
	Title       string   `json:"Title"`
	Genre       []string `json:"Genre"`
	Description string   `json:"Description"`
	Director    string   `json:"Director"`
	Actors      []string `json:"Actors"`
	Ratings     float32  `json:"Ratings"`
}

func (m *Movie) SetId() {
	m.Id = uuid.New().String()
}
