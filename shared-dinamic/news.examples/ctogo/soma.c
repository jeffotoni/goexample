#include "soma.h"

// gcc -shared -o libsoma.so soma.c
// ar rcs libsoma.a libsoma.so
int Soma(int x, int y) {
    return x + y;
}

