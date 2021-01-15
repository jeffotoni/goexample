//GitHub: HenriqueIni
//https://www.blogcyberini.com/

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

//função auxiliar para realizar as trocas de elementos
void swap(int A[], int i, int j){
    int temp = A[i];
    A[i] = A[j];
    A[j] = temp;
}

int partition(int A[], int inicio, int fim) {
    //procura a mediana entre inicio, meio e fim
    int meio = (inicio + fim) / 2;
    int a = A[inicio];
    int b = A[meio];
    int c = A[fim];
    int medianaIndice; //índice da mediana
    //A sequência de if...else a seguir verifica qual é a mediana
    if (a < b) {
        if (b < c) {
            //a < b && b < c
            medianaIndice = meio;
        } else {
            if (a < c) {
                //a < c && c <= b
                medianaIndice = fim;
            } else {
                //c <= a && a < b
                medianaIndice = inicio;
            }
        }
    } else {
        if (c < b) {
            //c < b && b <= a
            medianaIndice = meio;
        } else {
            if (c < a) {
                //b <= c && c < a
                medianaIndice = fim;
            } else {
                //b <= a && a <= c
                medianaIndice = inicio;
            }
        }
    }
    //coloca o elemento da mediana no fim para poder usar o Quicksort de Cormen
    swap(A, medianaIndice, fim);
        
    //*******************ALGORITMO DE PARTIÇÃO DE CORMEN*********************
    //o pivo é o elemento final
    int pivo = A[fim];
    int i = inicio - 1;
    int j;
    /*
     * Este laço irá varrer os vetores da esquerda para direira
     * procurando os elementos que são menores ou iguais ao pivô.
     * Esses elementos são colocados na partição esquerda.         
     */
    for (j = inicio; j <= fim - 1; j++) {
        if (A[j] <= pivo) {
            i = i + 1;
            swap(A, i, j);
        }
    }
    //coloca o pivô na posição de ordenação
    swap(A, i + 1, fim);
    return i + 1; //retorna a posição do pivô
}
//Quicksort mediana de três
void quicksortMedianaDeTres(int A[], int inicio, int fim) {
    if (inicio < fim) {
        //realiza a partição
        int q = partition(A, inicio, fim);
        //ordena a partição esquerda
        quicksortMedianaDeTres(A, inicio, q - 1);
        //ordena a partição direita
        quicksortMedianaDeTres(A, q + 1, fim);
    }
}

int main(int argc, char *argv[]){

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
      //printf("\n%d",number[i]);
      number[i] = rand();
   }
   clock_t toc1 = clock();
   printf("\nVetor: %f seconds\n", (double)(toc1 - tic1) / CLOCKS_PER_SEC);

   clock_t tic = clock();
   quicksortMedianaDeTres(number,0,count-1);
   clock_t toc = clock();

   printf("\nQuicksort: %f seconds\n", (double)(toc - tic) / CLOCKS_PER_SEC);

   // for(i=0;i<count;i++){
   //    printf("\n%d",number[i]);
   // }
}

