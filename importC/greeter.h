


#include "greeter.h"
#include <stdio.h>

int greet(const char *name, int year, char *out) {
    int n;
    
    n = sprintf(out, "Greetings, %s from %d! We come in peace :)", name, year);

    return n;
}