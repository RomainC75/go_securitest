package scenarios

type IScenario interface {
	Check() error
	Run() interface{}
}
