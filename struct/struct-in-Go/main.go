// @autor jeffotoni
// @year 2021
package main

// #include <stdio.h>
// #include <stdlib.h>

// struct Login {
//     char *User;
//     char *Email;
// };

//In Go would look like this
type Login struct {
	User  string
	Email string
	Cpf   int
}

// int main() {

// In Go
func main() {

	//struct Login login, *pointerToLogin;

	// In Go
	login := Login{User: "jeffotoni", Email: "jef@m.com"}

	login.User = "jeffotoni"
	login.Email = "jef@m.com"
	login.Cpf = 3987665487

	// pointerToLogin = malloc(sizeof(struct Login));
	// pointerToLogin->User = "pike";
	// pointerToLogin->Email = "pike@g.com";

	var pointerToLogin = new(Login)
	pointerToLogin.User = "pike"
	pointerToLogin.Email = "pike@g.com"
	pointerToLogin.Cpf = 3898798761

	//printf("login vals: %s %s\n", login.User, login.Email);
	//printf("pointerToLogin: %p %s %s\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email);

	//fmt.Printf("login vals: %s %s %d\n", login.User, login.Email, login.Cpf)
	println("login vals:\n", login.User, login.Email, login.Cpf)
	//fmt.Printf("pointerToLogin: %v %s %s %d\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email, pointerToLogin.Cpf)
	println("pointerToLogin:\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email, pointerToLogin.Cpf)

	//free(pointerToLogin);
	return
}
