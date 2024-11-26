package scenarios

type IScenario interface {
	Run() (interface{}, error)
	Check() error
}
