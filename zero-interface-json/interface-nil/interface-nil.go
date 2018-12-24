package main

type Named interface {
	Name() string
}

func greeting(thing Named) string {

	//if thing != nil {
	return "Hello " + thing.Name()
	//}

	//return ""
}

func main() {
	greeting(nil)
}
