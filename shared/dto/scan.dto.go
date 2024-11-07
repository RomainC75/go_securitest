package shared_dto

import (
	"shared/types"
	"time"
)

type Event struct {
	Id        int64                   `json:"id" validate:"required"`
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
	IPRange   IpRange   `json:"ip_range" validate:"required"`
	PortRange PortRange `json:"port_range" validate:"required"`
}

type IpRange struct {
	IpMin  string            `json:"ip_min" validate:"required"`
	IpMax  types.SNullString `json:"ip_max"`
	Unique bool              `json:"unique" validate:"boolean"`
}

type PortRange struct {
	Min int              `json:"min" validate:"required,number"`
	Max types.SNullInt32 `json:"max" `
}
