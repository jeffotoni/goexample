//@Autor:jeffotoni
//////////////////
#include <stdio.h>
#include <math.h>

// Define a função de ativação sigmoid
double sigmoid(double x) {
    return 1 / (1 + exp(-x));
}

// Calcula a saída de um neurônio
double calculateNeuronOutput(double weights[], double inputs[], int numberOfInputs) {
    double activation = weights[0]; // Bias
    for (int i = 0; i < numberOfInputs; i++) {
        activation += weights[i + 1] * inputs[i];
    }
    return sigmoid(activation);
}

// Exemplo de uma rede com uma camada oculta
int main() {
    // Exemplo de entrada
    double inputs[] = {1.0, 2.0, 3.0}; // 3 entradas
    int numberOfInputs = sizeof(inputs) / sizeof(double);

    // Pesos para a camada oculta (incluindo o bias como o primeiro peso)
    // Assumindo 2 neurônios na camada oculta
    double hiddenLayerWeights[2][4] = {
        {0.5, -0.6, 0.1, 0.4}, // Neurônio 1 da camada oculta com 3 pesos + bias
        {0.8, 0.2, 0.3, -0.5}  // Neurônio 2 da camada oculta com 3 pesos + bias
    };

    // Pesos para a camada de saída (assumindo 
    // 1 neurônio de saída e 2 entradas + bias)
    double outputLayerWeights[3] = {0.9, -1.2, 0.3};

    // Calcular saída da camada oculta
    double hiddenLayerOutputs[2];
    for (int i = 0; i < 2; i++) { // Para cada neurônio na camada oculta
        hiddenLayerOutputs[i] = calculateNeuronOutput(hiddenLayerWeights[i], inputs, numberOfInputs);
    }

    // Calcular saída da camada de saída
    double finalOutput = calculateNeuronOutput(outputLayerWeights, hiddenLayerOutputs, 2); // 2 entradas da camada oculta

    printf("Saída da rede neural: %f\n", finalOutput);

    return 0;
}

