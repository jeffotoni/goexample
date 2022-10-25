import (
	"fmt"
	"sync"
	"time"
)

var pool *sync.Pool

type A struct {
	Name string
}

func initPool() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(A)
		},
	}
}

func main() {
	initPool()
	one := pool.Get().(*A)
	one.Name = "jeffotoni"
	fmt.Printf("one.Name = %s\n", one.Name)
	pool.Put(one)
}

// Agora, a mesma instância se torna utilizável por outra rotina sem alocá-la novamente