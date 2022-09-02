#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define MAXBUF  (50)

char * message( char *s )
{
    char * msg = (char*) malloc( (MAXBUF + 1) * sizeof(char) );
    //strncpy( msg, "jeffotoni, let's Go!", MAXBUF );
    strncpy( msg, s, MAXBUF );
    return msg;
}

int main( void)
{
    //char s[MAXBUF] = "jeffotoni, let's Go!";
    char *s = "jeffotoni, let's Go!";
    //strncpy(s, "jeffotoni, let's Go!", MAXBUF);
    char *msg = message(s);
    printf("%s\n", msg );
    free(msg);
    return 0;
}
