#include <stdio.h>      /* printf */
#include <stdlib.h>     /* qsort */
#include <time.h>       /* clock_t*/

// int values[] = { 40, 10, 100, 90, 20, 25 };

int compare (const void * a, const void * b)
{
  return ( *(int*)a - *(int*)b );
}

int main(int argc, char *argv[])
{

  int qtd;
  int *number;

  qtd = atoi(argv[1]);
  printf("Quant:%d\n", qtd);
  if (qtd <= 1) {
    printf("Error qtd tem que ser maior que 1!");
    return 0;
  }

  number = malloc (qtd * sizeof (int));

  int i, count;
  count = qtd;

  clock_t tic1 = clock();
   srand(time(0));
   for(i=0;i<count;i++) {
      number[i] = rand();
      //printf ("\n%d ",number[i]);
   }
   clock_t toc1 = clock();
   printf("\nVetor: %f seconds\n", (double)(toc1 - tic1) / CLOCKS_PER_SEC);


  clock_t tic = clock(); 
  qsort (number, qtd, sizeof(int), compare);
  clock_t toc = clock();
  printf("\nQuicksort: %f seconds\n", (double)(toc - tic) / CLOCKS_PER_SEC);

  // for (i=0; i<qtd; i++)
  //    printf ("\n%d ",number[i]);
  return 0;
}

