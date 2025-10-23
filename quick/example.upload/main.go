package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/jeffotoni/quick"
)

func main() {
	app := quick.New()

	// Endpoint para upload de avatar usando base64
	app.Post("/v1/user/avatar", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")

		// Estrutura para receber o upload
		type UploadRequest struct {
			Filename string `json:"filename"`
			Content  string `json:"content"` // base64 encoded
		}

		var req UploadRequest
		err := c.BodyParser(&req)
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{
				"success": false,
				"message": "Erro ao fazer parse do JSON",
				"error":   err.Error(),
			})
		}

		if req.Filename == "" || req.Content == "" {
			return c.Status(400).JSON(map[string]interface{}{
				"success": false,
				"message": "Campos 'filename' e 'content' são obrigatórios",
			})
		}

		// Decodifica o base64
		fileContent, err := base64.StdEncoding.DecodeString(req.Content)
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{
				"success": false,
				"message": "Erro ao decodificar base64",
				"error":   err.Error(),
			})
		}

		// Valida extensão
		ext := strings.ToLower(filepath.Ext(req.Filename))
		allowedExts := map[string]bool{
			".jpg": true, ".jpeg": true, ".png": true,
			".gif": true, ".webp": true, ".bmp": true,
		}

		if !allowedExts[ext] {
			return c.Status(400).JSON(map[string]interface{}{
				"success": false,
				"message": "Formato não suportado. Use: jpg, jpeg, png, gif, webp, bmp",
			})
		}

		// Cria diretório
		uploadsDir := "./uploads"
		os.MkdirAll(uploadsDir, 0755)

		// Nome único
		filename := fmt.Sprintf("avatar_%d%s", time.Now().Unix(), ext)
		filePath := filepath.Join(uploadsDir, filename)

		// Salva arquivo
		err = os.WriteFile(filePath, fileContent, 0644)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"success": false,
				"message": "Erro ao salvar arquivo",
				"error":   err.Error(),
			})
		}

		return c.Status(200).JSON(map[string]interface{}{
			"success":      true,
			"message":      "Avatar salvo com sucesso!",
			"filename":     filename,
			"size":         len(fileContent),
			"path":         filePath,
			"uploaded_at":  time.Now().Format(time.RFC3339),
			"original_name": req.Filename,
		})
	})

	// Endpoint alternativo para upload usando multipart (raw HTTP)
	app.Post("/v1/user/avatar/multipart", func(c *quick.Ctx) error {
		c.Set("Content-Type", "application/json")

		// Tenta ler o multipart diretamente
		mr, err := c.Request.MultipartReader()
		if err != nil {
			return c.Status(400).JSON(map[string]interface{}{
				"success": false,
				"message": "Não foi possível ler multipart form",
				"error":   err.Error(),
				"hint":    "Use o endpoint /v1/user/avatar com JSON {\"filename\": \"...\", \"content\": \"base64...\"}",
			})
		}

		for {
			part, err := mr.NextPart()
			if err == io.EOF {
				break
			}
			if err != nil {
				return c.Status(400).JSON(map[string]interface{}{
					"success": false,
					"message": "Erro ao ler parte do multipart",
					"error":   err.Error(),
				})
			}

			if part.FormName() == "avatar" && part.FileName() != "" {
				// Valida extensão
				ext := strings.ToLower(filepath.Ext(part.FileName()))
				allowedExts := map[string]bool{
					".jpg": true, ".jpeg": true, ".png": true,
					".gif": true, ".webp": true, ".bmp": true,
				}

				if !allowedExts[ext] {
					return c.Status(400).JSON(map[string]interface{}{
						"success": false,
						"message": "Formato não suportado",
					})
				}

				// Cria diretório
				uploadsDir := "./uploads"
				os.MkdirAll(uploadsDir, 0755)

				// Nome único
				filename := fmt.Sprintf("avatar_%d%s", time.Now().Unix(), ext)
				filePath := filepath.Join(uploadsDir, filename)

				// Salva arquivo
				dst, err := os.Create(filePath)
				if err != nil {
					return c.Status(500).JSON(map[string]interface{}{
						"success": false,
						"message": "Erro ao criar arquivo",
						"error":   err.Error(),
					})
				}
				defer dst.Close()

				size, err := io.Copy(dst, part)
				if err != nil {
					return c.Status(500).JSON(map[string]interface{}{
						"success": false,
						"message": "Erro ao salvar arquivo",
						"error":   err.Error(),
					})
				}

				return c.Status(200).JSON(map[string]interface{}{
					"success":     true,
					"message":     "Avatar salvo com sucesso!",
					"filename":    filename,
					"size":        size,
					"path":        filePath,
					"uploaded_at": time.Now().Format(time.RFC3339),
				})
			}
		}

		return c.Status(400).JSON(map[string]interface{}{
			"success": false,
			"message": "Campo 'avatar' não encontrado",
		})
	})

	// Rota de teste
	app.Get("/", func(c *quick.Ctx) error {
		return c.Status(200).SendString(`Upload API - Quick Framework

Endpoints:
- POST /v1/user/avatar (JSON com base64)
- POST /v1/user/avatar/multipart (multipart/form-data)

Exemplo de uso com JSON:
curl -X POST http://localhost:8080/v1/user/avatar \
  -H "Content-Type: application/json" \
  -d '{"filename":"avatar.png","content":"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mNk+M9QDwADhgGAWjR9awAAAABJRU5ErkJggg=="}'

Exemplo de uso com multipart:
curl -X POST http://localhost:8080/v1/user/avatar/multipart \
  -F "avatar=@imagem.png"`)
	})

	fmt.Println("🚀 Servidor rodando em http://localhost:8080")
	fmt.Println("📤 Endpoint principal: POST http://localhost:8080/v1/user/avatar")
	fmt.Println("📤 Endpoint alternativo: POST http://localhost:8080/v1/user/avatar/multipart")
	app.Listen(":8080")
}
