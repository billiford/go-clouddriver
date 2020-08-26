package kubernetes

//go:generate counterfeiter . Action
type Action interface {
	Run() error
}
