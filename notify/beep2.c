#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/vt.h>
#include <sys/kd.h>
#include <unistd.h>
#include <sys/ioctl.h>
#include <errno.h>
#define RAPIDO 100000
#define LENTO 400000

void toca_som(char som);
char letras[38];
struct cod {
        char morse[9];
};
char tmp[1];
float vel;
int i,b,c;

int main() 
{
        letras[0] = 'm';
        letras[1] = ',';
        letras[2] = '.';
        letras[3] = '1';
        letras[4] = '0';
        letras[5] = '3';
        letras[6] = '2';
        letras[7] = '5';
        letras[8] = '4';
        letras[9] = '7';
        letras[10] = '6';
        letras[11] = '9';
        letras[12] = '8';
        letras[13] = '?';
        letras[14] = 'a';
        letras[15] = 'c';
        letras[16] = 'b';
        letras[17] = 'e';
        letras[18] = 'd';
        letras[19] = 'g';
        letras[20] = 'f';
        letras[21] = 'i';
        letras[22] = 'h';
        letras[23] = 'k';
        letras[24] = 'j';
        letras[25] = 'l';
        letras[26] = 'o';
        letras[27] = 'n';
        letras[28] = 'q';
        letras[29] = 'p';
        letras[30] = 's';
        letras[31] = 'r';
        letras[32] = 'u';
        letras[33] = 't';
        letras[34] = 'w';
        letras[35] = 'v';
        letras[36] = 'y';
        letras[37] = 'x';
        letras[38] = 'z';

        struct cod cod[38];
        strcpy(cod[0].morse,"--");
        strcpy(cod[1].morse, "--.--");
        strcpy(cod[2].morse,".-.-.-");
        strcpy(cod[3].morse,".----");
        strcpy(cod[4].morse,"-----");
        strcpy(cod[5].morse,"...--");
        strcpy(cod[6].morse,"..---");
        strcpy(cod[7].morse,".....");
        strcpy(cod[8].morse,"....-");
        strcpy(cod[9].morse,"--...");
        strcpy(cod[10].morse,"-....");
        strcpy(cod[11].morse,"----.");
        strcpy(cod[12].morse,"---..");
        strcpy(cod[13].morse,"..--..");
        strcpy(cod[14].morse,".-");
        strcpy(cod[15].morse,"-.-.");
        strcpy(cod[16].morse,"-...");
        strcpy(cod[17].morse,".");
        strcpy(cod[18].morse,"-..");
        strcpy(cod[19].morse,"--.");
        strcpy(cod[20].morse,"..-.");
        strcpy(cod[21].morse,"..");
        strcpy(cod[22].morse,"....");
        strcpy(cod[23].morse,"-.-");
        strcpy(cod[24].morse,".---");
        strcpy(cod[25].morse,".-..");
        strcpy(cod[26].morse,"---");
        strcpy(cod[27].morse,"-.");
        strcpy(cod[28].morse,"--.-");
        strcpy(cod[29].morse,".--.");
        strcpy(cod[30].morse,"...");
        strcpy(cod[31].morse,".-.");
        strcpy(cod[32].morse,"..-");
        strcpy(cod[33].morse,"-");
        strcpy(cod[34].morse,".--");
        strcpy(cod[35].morse,"...-");
        strcpy(cod[36].morse,"-.--");
        strcpy(cod[37].morse,"-..-");
        strcpy(cod[38].morse,"--..");

        char frase[30];
        printf("Escreva a frase: \n");
        while ( b != '\n' && i < 30)
         {
                         b = getchar();
                         frase[i] = tolower(b);
                         i++;
                 }
        frase[i -1] = '\0';
        printf("escolha a velocidade(1 para rápido): \n");
        b = getchar();
        tmp[0] = b;
        tmp[1] = '\0';
        if (tmp == "1")
         {
                         vel = RAPIDO;
                 } else {
                                 vel = LENTO;
                         } 
        for (i=0; i <= strlen(frase);i++)
        {
                        printf("letra atual: %c \n",frase[i]);
                        for (c=0; c<=38;c++)
                         {
                                                 if ( c <= strlen(frase))
                                                 {
                                                                                 if (frase[i]==letras[c])
                                                                                 {                
                                                                                                                         for (b=0; b<= strlen(cod[c].morse);b++)
                                                                                                                          {
                                                                                                                                                                          toca_som(cod[c].morse[b]);
                                                                                                                                                                  }
                                                                                                                 }
                                                                         }
                                 }       }
        return(0);
}

void toca_som(char som) 
{

        unsigned int count;
        count = 2450;

        if (som=='.')
         {
                         ioctl(2, KIOCSOUND, count);
                         usleep(50*vel);
                         ioctl(2, KIOCSOUND, 0);
                         usleep(0.1*vel);
                 }
        else if (som== '-')
         {
                         ioctl(2, KIOCSOUND, count);
                         usleep(150*vel);
                         ioctl(2, KIOCSOUND, 0);
                         usleep(0.1*vel);
                 } else {
                                 usleep(0.3*vel);
                         }

}
