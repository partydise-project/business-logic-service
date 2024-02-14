package database

import "gorm.io/gorm"

type PlanContratado struct {
	gorm.Model
	PlanEventoID uint   `json:"plan_evento_ID" gorm:"not null"`
	ClientID     uint   `json:"client_ID" gorm:"not null"`
	Pagada       string `json:"pagada"`
	EventStatus  string `json:"event_status"`
	ClientRating string `json:"client_rating"`
}

type EventoTrabajado struct {
	gorm.Model
	PlanContratadoID uint           `json:"plan_contratado_ID" gorm:"not null"`
	RecreadorID      uint           `json:"recreador_ID" gorm:"not null"`
	ExecutionDate    string         `json:"execution_date"`
	PlanContratado   PlanContratado `json:"plan_contratado"`
}
