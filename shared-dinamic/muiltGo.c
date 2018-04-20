/**
*
* @description: 3 formas de criar bibliotecas em C / linux
* 
* Primeira: Criar o arquivo .h e usar incluir no main
* 
* Segunda:  Criar o arquivo.h e o arquivo.c da biblioteca e compilar para ficar dinamica
* copiar a lib para /usr/lib
* 
* Terceira: Criar o arquivo.h e o arquivo.c da biblioteca copilar
* depois fazer a chamada dinamicas no main
* 
* @abstract: Criação de Bibliotecas Dinâmicas em C
* Carregando em tempo de exeucao biblioteca Multi
*
* padrao lib linux / nome: lib<nome>.so.<versao>
* encontra-se em : /usr/lib ou /lib
* gcc -l => biblioteca dinâmica = passando nome da lib => -lm ou -lglut
*
* utilizar o comando ldd no executavel para visualizar as libs
*
* copilando para gerar biblioteca dinamica
* gcc -shared -fPIC Multi.c -o libMulti.so
* gcc -o mainload main1.c -L. -lMulti
* sudo cp libMulti.so /lib
* 
* agora fazendo um programa carregar em tempo de execusao a lib
* gcc -o main2 main2.c  -ldl
* 
* @name: @jeffotoni
* @date: 2016
*
* #compilando o codigo sem biblioteca
* gcc main.c -c
*
* #linkagem
* gcc -o main main.o
*
* main #compilando nossa biblioteca
* gcc -c Multi.c
* 
* # compilando e linkando nossa biblioteca no main
* gcc -o main main.c Multi.o
* 
*
*/

#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>

int main()
{
  
  ///declara ponteiro para receber lib
  void *handle;


  ///assinatura da funcao Multi
  unsigned long int (*Multi)(unsigned long int f);

  ///capturar o valor na tela..
  int  valorToGerar;
  int n , i = 0, c;

  ///erro para tratar dlopen
  char *error;
  
  ///dlopen => abrindo biblioteca
  handle = dlopen("./C/libmulti.so",RTLD_LAZY);
  
  if(!handle)
  {
    fprintf(stderr, "Erro: %s\n", dlerror());
    exit(1);
  }

  ///chando a funcao fibonancci  
  Multi  = dlsym(handle, "Multi");

  if((error = dlerror()) != NULL)
  {
    fprintf(stderr,"Erro: %s\n", error);
    exit(1);
  }
  
  printf("Quantidade de termos da sequencia Multi:\n");
  scanf("%d", &n);
  
  for ( c = 1 ; c <= n ; c++ )
  { 
    valorToGerar = c;
    printf("Sequencia: %d * 2 = %lu\n",valorToGerar,(*Multi)(c));
    i++; 
  }
  
  dlclose(handle);
  return 0;
}

