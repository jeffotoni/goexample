#include "soma.h"

// export LD_LIBRARY_PATH=.
// gcc -c -o libsoma.a soma.c
// ar rcs libsoma.a libsoma.o
int Soma(int x, int y) {
    return x + y;
}

