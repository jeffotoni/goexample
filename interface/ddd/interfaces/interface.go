package interfaces

type Interface interface {
	create() bool
}

type Interface2 interface {
	Interface
}

func Create(i Interface2) bool {
	return i.create()
}
