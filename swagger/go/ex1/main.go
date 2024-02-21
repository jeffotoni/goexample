package main

import (
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/jeffotoni/goexample/swagger/go/ex1/docs"
)

// @description estrutura de resposta para uma criação bem-sucedida de usuário
type UserResponse struct {
	Message string `json:"message"`
}

// User representa a estrutura dos dados do usuário recebidos na requisição
// @description estrutura de dados do usuário
type User struct {
	Nome  string `json:"nome"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	http.HandleFunc("/v1/user", userHandler)

	log.Println("Run Server 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// User handler
// @Summary Adiciona um novo usuário
// @Description adiciona um novo usuário ao sistema
// @ID user-create
// @Accept  json
// @Produce  json
// @Success 201 {object} UserResponse
// @Router /v1/user [post]
func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Aqui você pode processar os dados do usuário conforme necessário

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UserResponse{Message: "Usuário criado com sucesso"})
}
