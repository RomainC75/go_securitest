package events

type ScenarioSelector string

const (
	PortScanner   ScenarioSelector = "task:port_scanner"
	ScannerResult ScenarioSelector = "task:scanner_result"
)

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}
