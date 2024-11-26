package shared_dto

import (
	"time"
)

type Event struct {
	Id        int64                   `json:"id" validate:"required"`
	CreatedAt time.Time               `json:"createdAt" validate:"required"`
	Scenario  int                     `json:"scenario" validate:"required"`
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
	IpMin  string  `json:"ip_min" validate:"required"`
	IpMax  *string `json:"ip_max,omitempty"`
	Unique bool    `json:"unique" validate:"boolean"`
}

type PortRange struct {
	Min int32  `json:"min" validate:"required,number"`
	Max *int32 `json:"max,omitempty" `
}
