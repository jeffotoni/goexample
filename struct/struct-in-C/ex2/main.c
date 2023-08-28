#include<stdio.h>
#include<stdlib.h>
#include<string.h>

struct Cliente{
	char *Cpf;
};

struct Cliente cli;

typedef struct{
	char *Cnpj;
}Empresa;

typedef struct Endereco{
	char *Rua;
}End;

void main() {
	Empresa emp;
	typedef struct Cliente cli2;
	
	cli.Cpf = malloc(15);
	cli.Cpf = strcpy(cli.Cpf,"038.847.393-33");
	free(cli.Cpf);

	cli2 *myCli2 = malloc(sizeof(cli2));
	myCli2->Cpf = malloc(15);
	strcpy(myCli2->Cpf, "393.393.595-83");
	free(myCli2->Cpf);
	free(myCli2);

	emp.Cnpj = malloc(20);
	strcpy(emp.Cnpj, "09.939.393/0001-98");	
	free(emp.Cnpj);

	Empresa *emp2 = malloc(sizeof(Empresa));
	emp2->Cnpj = malloc(25);
	strcpy(emp2->Cnpj, "98.393.399/0001-39");
	printf("Cnpj: %s\n", emp2->Cnpj);
	free(emp2->Cnpj);
	free(emp2);
}


