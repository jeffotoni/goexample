#include <iostream>
#include <cmath>
#include <chrono>

using namespace std;

int main() {
  int n;
  cin >> n;

  auto start = chrono::high_resolution_clock::now();
  // Cria um vetor de booleanos, onde cada índice representa um número.
  // O vetor é inicializado com todos os valores `true`.

  bool is_prime[n + 1];
  for (int i = 0; i <= n; i++) {
    is_prime[i] = true;
  }

  // Começa no número 2.

  for (int i = 2; i <= n; i++) {
    // Verifica se o número é divisível por qualquer número menor ou igual à sua raiz quadrada.
    for (int j = 2; j <= sqrt(i); j++) {
      if (is_prime[i] && i % j == 0) {
        is_prime[i] = false;
        break;
      }
    }
  }

  // Finaliza o cronômetro
  auto end = chrono::high_resolution_clock::now();

  // Calcula o tempo de execução em milissegundos
  auto duration = chrono::duration_cast<chrono::milliseconds>(end - start);

  // Retorna uma lista de números primos.

  for (int i = 2; i <= n; i++) {
    if (is_prime[i]) {
      cout << "\n" << i << " ";
    }
  }

  cout << endl;

  // Imprime o tempo de execução
  cout << "\nTempo de execução: " << duration.count() << " milissegundos" << endl;


  return 0;
}

