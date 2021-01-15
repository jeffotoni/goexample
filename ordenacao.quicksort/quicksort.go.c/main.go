package main

// go build -o quicksort.so -buildmode=c-shared quicksort.go
// gcc -o quicksort quicksort.c ./quicksort.so &&
// ./quicksort

// #include <stdio.h>
// #include <stdlib.h>
// #include <time.h>       /* clock_t*/
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
//
// int compare (const void * a, const void * b)
// {
//   return ( *(int*)a - *(int*)b );
// }
//
// static int quickC(int qtd)
// {
//    int *number;
//    printf("Quant:%d\n", qtd);
//    if (qtd <= 1) {
//      printf("Error qtd tem que ser maior que 1!");
// 	    return 0;
//    }
//    number = malloc (qtd * sizeof (int));
//    int i, count;
//    count = qtd;
//    clock_t tic1 = clock();
//    srand(time(0));
//    for(i=0;i<count;i++) {
//       number[i] = rand();
//       //printf ("\n%d ",number[i]);
//    }
//
//    clock_t toc1 = clock();
//    printf("\nVetor: %f seconds\n", (double)(toc1 - tic1)/CLOCKS_PER_SEC);
//    clock_t tic = clock();
//    qsort (number, qtd, sizeof(int), compare);
//    clock_t toc = clock();
//    printf("\nQuicksort: %f seconds\n", (double)(toc - tic)/CLOCKS_PER_SEC);
//    //for (i=0; i<count; i++)
//       //printf ("\n%d ",number[i]);
// 	   return 0;
// }
import "C"
import (
	"os"
	"strconv"
)

func main() {
	// cs := C.CString("Hello from stdio")
	// C.myprint(cs)
	// C.free(unsafe.Pointer(cs))
	qtd, _ := strconv.Atoi(os.Args[1])
	if qtd <= 1 {
		println("O valor tem que ser maior que 1!")
		return
	}

	C.quickC(C.int(qtd))
	//C.quickC(C.int(qtd))
	println("done")
}
