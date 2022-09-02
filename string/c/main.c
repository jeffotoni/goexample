#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#define MAXBUF  (50)

char * message( void )
{
    char * msg = (char*) malloc( (MAXBUF + 1) * sizeof(char) );
    strncpy( msg, "jeffotoni, let's Go!", MAXBUF );
    return msg;
}

int main( void)
{
    char * msg = message();
    printf("%s\n", msg );
    free(msg);
    return 0;
}
