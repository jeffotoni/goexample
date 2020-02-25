#include <stdio.h>

int main() {
    unsigned char c;

    puts("Escreva o que desejar. Ao terminar, inclua um ponto (.) para fechar o programa.");

    do {
        c = getchar();
        putchar(c);
    } while(c != '.');

    return 0;
}

