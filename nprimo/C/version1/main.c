#include <stdio.h>
#include <math.h>

int main() {
  int i, j;
  int primes[100000];
  int num_primes = 0;

  for (i = 2; i <= 100000; i++) {
    if (i % 2 == 0) {
      if (i == 2) {
        primes[num_primes++] = i;
      }
    } else {
      for (j = 3; j * j <= i; j++) {
        if (i % j == 0) {
          break;
        }
      }
      if (j * j > i) {
        primes[num_primes++] = i;
      }
    }
  }

  for (i = 0; i < num_primes; i++) {
    printf("%d\n", primes[i]);
  }

  return 0;
}

