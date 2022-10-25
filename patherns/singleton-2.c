/*
* Example Singleton C
* @package     main
* @author      @jeffotoni
* @size        10/09/2018
*/

#include <stdio.h>

// esta variavel
// nao sera acessada
// por outro modulo
// static global
// static int conn_pg2 = 0

int *Connect() {
	
	// variavel static local
	static int conn_pg = 0; // nosso ojbeto connexao

	 // existe connexao ?	
	 if (conn_pg == 0) {

	 	// Db.Connect
	 	conn_pg = 5454;	// atribuindo conexao postgres
	 }
	 	
	 // conn_pg++; // o valor ser atribuido sempre
	 // mantendo o valor, somente se for static

	 // retorna objeto
	 return &conn_pg;
 }

 int main() {

 	int Db = *Connect();

    printf("Connect em C: %d\n", Db);
    printf("Connect em C: %d\n", *Connect());
    printf("Connect em C: %d\n", *Connect());
    return 0;
}