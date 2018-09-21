/*
* Example Singleton C
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
*/

#include <stdlib.h>
#include <stdio.h>

struct Conn
{
    int conn_pg;
};

struct Conn* Connect()
{
    static struct Conn *instance = NULL;

    // entra somente
    // uma unica vez
    if(instance == NULL)
    {
        instance = malloc(sizeof(*instance));
        instance->conn_pg = 34498;
    }
    
    return instance;
};

int main() {

    printf("Connect em C: %d\n", Connect()->conn_pg);
    printf("Connect em C: %d\n", Connect()->conn_pg);

    return 0;
}