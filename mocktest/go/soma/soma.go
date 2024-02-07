package soma

// Calculadora define a interface que queremos mockar.
type Calculadora interface {
	Soma(a, b int) int
}

// MinhaCalculadora é uma implementação concreta de Calculadora.
type MinhaCalculadora struct{}

// Soma implementa a operação de soma para MinhaCalculadora.
func (c MinhaCalculadora) Soma(a, b int) int {
	return a + b
}
