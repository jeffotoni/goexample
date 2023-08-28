/**
 * @autor jeffotoni
 * @year 2021
 */

#include <stdio.h>
#include <stdlib.h>

struct Login {
    char *User;
    char *Email;
    unsigned long Cpf;
};

/**
  // In Go would look like this
  type Login struct {
    User string
    Email string
    Cpf int
  }
 */

int main() {
/** 
 // In Go 
func main() {
*/
    
    struct Login login, *pointerToLogin;

    /**
    // In Go
    login := Login{User:"jeffotoni", Email:"jef@m.com"}
    */

    login.User = "jeffotoni";
    login.Email = "jef@m.com";
    login.Cpf = 3998756773;

    pointerToLogin = malloc(sizeof(struct Login));
    pointerToLogin->User = "pike";
    pointerToLogin->Email = "pike@g.com";
    pointerToLogin->Cpf = 3456789891;

    /** 
    // In Go
    var pointerToLogin  = new(Login) // or var pointerToLogin = &Login{}
    pointerToLogin.User  = "pike"
    pointerToLogin.Email = "pike@g.com"
    pointerToLogin.Cpf = 03987669671
    */
    printf("login vals: %s %s %ld\n", login.User, login.Email, login.Cpf);
    printf("pointerToLogin: %p %s %s %ld\n", pointerToLogin, pointerToLogin->User, pointerToLogin->Email, pointerToLogin->Cpf);

    /**
     fmt.Printf("login vals: %s %s\n", login.User, login.Email)
     fmt.Printf("pointerToLogin: %v %s %s\n", pointerToLogin, pointerToLogin.User, pointerToLogin.Email);
    */

    free(pointerToLogin);
    return 0;
}

