package work_dto

type ScenarioBase struct {
	BasicData string `json:"basic_data" validate:"required"`
}

type NetworkDiscover struct {
	ScenarioBase
	Range IpRange `json:"ip_range" validate:"required"`
}

type FullPortTestScenario struct {
	ScenarioBase
	PortTestScenario PortTestScenario `json:"scenario" validate:"required"`
}

type PortTestScenario struct {
	UserId    int32   `json:"user_id", validate:"required"`
	IPRange   IpRange `json:"ip_range" validate:"required"`
	PortRange Range   `json:"range" validate:"required"`
}

type IpRange struct {
	IpMin  string `json:"ip_min" validate:"required"`
	IpMax  string `json:"ip_max" validate:"required,ip"`
	Unique bool   `json:"unique" validate:"boolean"`
}

type Range struct {
	Min int `json:"min" validate:"required"`
	Max int `json:"max" validate:"required"`
}
