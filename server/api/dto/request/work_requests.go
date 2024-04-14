package request_dto

type ScenarioBase struct {
	BasicData string `json:"basic_data" validate:"required"`
}

type NetworkDiscover struct {
	ScenarioBase
	Range IpRange `json:"ip_range" validate:"required"`
}

type PortTestScenario struct {
	ScenarioBase
	Range IpRange `json:"ip_range" validate:"required"`
}

type IpRange struct {
	IpMin string `json:"ip_min" validate:"required,ip"`
	IpMax string `json:"ip_max" validate:"required,ip"`
}
