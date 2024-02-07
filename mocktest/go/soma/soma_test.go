package soma

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// CalculadoraMock é um mock para a interface Calculadora.
type CalculadoraMock struct {
	mock.Mock
}

// Soma é a implementação mockada de Soma.
func (m *CalculadoraMock) Soma(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0) // Retorna o primeiro argumento como int.
}

func TestSoma(t *testing.T) {
	// Cria uma nova instância do mock.
	mockCalc := new(CalculadoraMock)

	// Define o comportamento esperado.
	mockCalc.On("Soma", 5, 3).Return(8)

	// Testa a função Soma com os valores mockados.
	resultado := mockCalc.Soma(5, 3)

	// Verifica se o resultado é o esperado.
	mockCalc.AssertExpectations(t)
	if resultado != 8 {
		t.Errorf("Esperado 8, mas obteve %d", resultado)
		return
	}
	t.Log("success....")
}
