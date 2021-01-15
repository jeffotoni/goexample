//@autor jeff

#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void quicksort(int number[25],int first,int last){
   int i, j, pivot, temp;

   if(first<last){
      pivot=first;
      i=first;
      j=last;

      while(i<j){
         while(number[i]<=number[pivot]&&i<last)
            i++;
         while(number[j]>number[pivot])
            j--;
         if(i<j){
            temp=number[i];
            number[i]=number[j];
            number[j]=temp;
         }
      }

      temp=number[pivot];
      number[pivot]=number[j];
      number[j]=temp;
      quicksort(number,first,j-1);
      quicksort(number,j+1,last);

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
      //scanf("%d",&number[i]);
      number[i] = rand();
   }
   clock_t toc1 = clock();
   printf("Vetor: %f seconds\n", (double)(toc1 - tic1) / CLOCKS_PER_SEC);

   clock_t tic = clock();
   quicksort(number,0,count-1);
   clock_t toc = clock();

   printf("Quicksort: %f seconds\n", (double)(toc - tic) / CLOCKS_PER_SEC);

   // for(i=0;i<count;i++){
   //    printf(" %d",number[i]);
   // }

   return 0;
}