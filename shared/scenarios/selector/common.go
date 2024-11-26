package selector

import (
	shared_dto "shared/dto"
	"shared/scenarios/scenarios"
)

type Selector struct {
	Scenario scenarios.IScenario
}

func NewSelector(scenarioNum int, data interface{}) *Selector {
	var scenario scenarios.IScenario
	switch int(scenarioNum) {
	case 1:
		dataToAnalyse := data.(shared_dto.Event).Content.PortTestScenario
		scenario = scenarios.NewScan(dataToAnalyse)
	}
	return &Selector{
		Scenario: scenario,
	}
}
