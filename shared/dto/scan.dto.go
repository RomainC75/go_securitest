package shared_dto

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID               `json:"id" validate:"required"`
	CreatedAt time.Time               `json:"createdAt" validate:"required"`
	Content   FullPortTestScenarioReq `json:"content" validate:"required"`
}

type FullPortTestScenarioReq struct {
	ScenarioBase
	PortTestScenario PortTestScenario `json:"scenario" validate:"required"`
}
type ScenarioBase struct {
	BasicData string `json:"basic_data" validate:"required"`
}

type NetworkDiscover struct {
	ScenarioBase
	Range IpRange `json:"ip_range" validate:"required"`
}

type PortTestScenario struct {
	IPRange   IpRange `json:"ip_range" validate:"required"`
	PortRange Range   `json:"range" validate:"required"`
}

type IpRange struct {
	IpMin  string `json:"ip_min" validate:"required"`
	IpMax  string `json:"ip_max" validate:"required,ip"`
	Unique bool   `json:"unique" validate:"boolean"`
}

type Range struct {
	Min int `json:"min" validate:"required,number"`
	Max int `json:"max" validate:"required,number"`
}
