package work

type ControllerInterface interface {
	Before(*Context)
	After(*Context)
	Logic(*Context)
}
