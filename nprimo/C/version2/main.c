#include <stdio.h>
#include <time.h>

int main() {
  int i, j;
  int primes[1000000];
  int num_primes = 0;

  struct timespec start, end;

  clock_gettime(CLOCK_MONOTONIC, &start);

  for (i = 2; i <= 1000000; i++) {
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

  clock_gettime(CLOCK_MONOTONIC, &end);

  long long time_taken = (end.tv_sec - start.tv_sec) * 1000000000 + (end.tv_nsec - start.tv_nsec);

  long long ms = time_taken/1000000;

  for (i = 0; i < num_primes; i++) {
    printf("%d\n", primes[i]);
  }
  
 // long long ms = time_taken/1000000;
  printf("Time: %lld ms\n", ms);

  //char buffer[100];
  //sprintf(buffer, "%.3f", ms);
  //printf("%s\n", buffer);

  return 0;
}

