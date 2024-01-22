
// C++ program to print all primes smaller than or equal to
// n using Sieve of Eratosthenes
#include <bits/stdc++.h>
#include <chrono>

using namespace std;

int main()
{
    int n = 1000000;

    cout << "Following are the prime numbers smaller "

         << " than or equal to " << n << endl;
    
    auto start = chrono::high_resolution_clock::now();
    
    //SieveOfEratosthenes(n);

    // Create a boolean array "prime[0..n]" and initialize

    // all entries it as true. A value in prime[i] will

    // finally be false if i is Not a prime, else true.

    bool prime[n + 1];

    memset(prime, true, sizeof(prime));
 

    for (int p = 2; p * p <= n; p++) {

        // If prime[p] is not changed, then it is a prime

        if (prime[p] == true) {

            // Update all multiples of p

            for (int i = p * p; i <= n; i += p)

                prime[i] = false;

        }

    }
 
    // Finaliza o cronômetro
    auto end = chrono::high_resolution_clock::now();

    // Calcula o tempo de execução em milissegundos
    auto duration = chrono::duration_cast<chrono::milliseconds>(end - start);

    // // Print all prime numbers
    for (int p = 2; p <= n; p++)
        if (prime[p])
            cout << "\n" << p << " ";
   
    // Imprime o tempo de execução
    cout << "\nTempo de execução: " << duration.count() << " milissegundos" << endl;

    return 0;
}

