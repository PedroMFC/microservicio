package models

type Valoracion struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Asignatura  string `json:"asignatura"`
	Valoracion int `json:"valoracion"`
}
