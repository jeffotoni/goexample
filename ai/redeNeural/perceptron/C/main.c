//@Autor:jeffotoni
//////////////////

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

#define THRESHOLD 0
#define LEARNING_RATE 0.1
#define NUM_EPOCHS 100
#define NUM_FEATURES 2

// Função para calcular a saída do perceptron
int predict(int inputs[], float weights[], int numFeatures) {
    float activation = weights[0]; // Bias
    for (int i = 0; i < numFeatures; i++) {
        activation += inputs[i] * weights[i + 1];
    }
    return activation >= THRESHOLD ? 1 : 0;
}

// Função de treinamento do Perceptron
void trainPerceptron(int inputs[][NUM_FEATURES], int labels[], float weights[], int numSamples, int numFeatures) {
    for (int epoch = 0; epoch < NUM_EPOCHS; epoch++) {
        for (int i = 0; i < numSamples; i++) {
            int prediction = predict(inputs[i], weights, numFeatures);
            int error = labels[i] - prediction;
            weights[0] += LEARNING_RATE * error; // Ajustar o bias
            for (int j = 0; j < numFeatures; j++) {
                weights[j + 1] += LEARNING_RATE * error * inputs[i][j]; // Ajustar os pesos
            }
        }
    }
}

int main() {
    // Exemplo de dataset: [feature1, feature2], label
    int inputs[4][NUM_FEATURES] = {{0, 0}, {0, 1}, {1, 0}, {1, 1}};
    int labels[4] = {0, 0, 0, 1}; // Saídas esperadas

    // Inicializando pesos aleatoriamente
    float weights[NUM_FEATURES + 1]; // +1 para o bias
    srand(time(0));
    for (int i = 0; i < NUM_FEATURES + 1; i++) {
        weights[i] = (float)rand() / (float)RAND_MAX * 2.0 - 1.0; // Peso inicial entre -1 e 1
    }

    // Treinando o Perceptron
    trainPerceptron(inputs, labels, weights, 4, NUM_FEATURES);

    // Testando o Perceptron treinado
    for (int i = 0; i < 4; i++) {
        int prediction = predict(inputs[i], weights, NUM_FEATURES);
        printf("Input = [%d, %d], Predicted = %d, Expected = %d\n", inputs[i][0], inputs[i][1], prediction, labels[i]);
    }

    return 0;
}

