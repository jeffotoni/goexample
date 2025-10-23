# Example Upload - Quick Framework

Este exemplo demonstra como criar um servidor de upload de imagens usando o framework [Quick](https://github.com/jeffotoni/quick) do jeffotoni.

## Descrição

API simples para upload de avatar de usuário que:
- Recebe imagens através do endpoint `/v1/user/avatar`
- Suporta dois métodos: JSON com base64 e multipart/form-data
- Valida o formato do arquivo (jpg, jpeg, png, gif, webp, bmp)
- Salva a imagem no diretório `uploads/`
- Retorna JSON com confirmação e informações do arquivo

## Como executar

```bash
# Entre no diretório
cd example.upload

# Execute o servidor
go run main.go
```

O servidor estará rodando em `http://localhost:8080`

## Como testar

### Método 1: JSON com Base64 (Recomendado)

Este é o método recomendado para usar com o framework Quick:

```bash
# Primeiro, codifique sua imagem em base64
BASE64_IMG=$(base64 -w 0 sua_imagem.png)

# Faça o upload
curl -X POST http://localhost:8080/v1/user/avatar \
  -H "Content-Type: application/json" \
  -d "{\"filename\":\"avatar.png\",\"content\":\"$BASE64_IMG\"}"
```

### Método 2: Multipart Form Data (Alternativo)

```bash
# Upload usando multipart (endpoint alternativo)
curl -X POST http://localhost:8080/v1/user/avatar/multipart \
  -F "avatar=@sua_imagem.png"
```

**Nota**: O endpoint multipart pode ter limitações devido ao funcionamento interno do Quick framework.

### Resposta de sucesso

```json
{
  "success": true,
  "message": "Avatar salvo com sucesso!",
  "filename": "avatar_1234567890.png",
  "size": 245678,
  "path": "uploads/avatar_1234567890.png",
  "uploaded_at": "2025-10-23T10:30:45Z",
  "original_name": "avatar.png"
}
```

### Resposta de erro

```json
{
  "success": false,
  "message": "Campos 'filename' e 'content' são obrigatórios"
}
```

## Endpoints

### POST /v1/user/avatar
Upload de imagem usando JSON com base64

**Request:**
```json
{
  "filename": "avatar.png",
  "content": "iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJ..."
}
```

**Response:**
```json
{
  "success": true,
  "message": "Avatar salvo com sucesso!",
  "filename": "avatar_1761227642.png",
  "size": 67,
  "path": "uploads/avatar_1761227642.png",
  "uploaded_at": "2025-10-23T13:54:02Z",
  "original_name": "avatar.png"
}
```

### POST /v1/user/avatar/multipart
Upload de imagem usando multipart/form-data (alternativo)

**Request:**
```bash
curl -F "avatar=@imagem.png" http://localhost:8080/v1/user/avatar/multipart
```

### GET /
Página inicial com informações e exemplos de uso

## Formatos suportados

- JPG / JPEG
- PNG
- GIF
- WEBP
- BMP

## Estrutura do projeto

```
example.upload/
├── main.go          # Código principal do servidor
├── go.mod           # Dependências do projeto
├── go.sum           # Checksums das dependências
├── uploads/         # Diretório onde os arquivos são salvos
└── README.md        # Este arquivo
```

## Features

- ✅ Upload de imagens via JSON com base64
- ✅ Upload alternativo via multipart/form-data
- ✅ Validação de formato
- ✅ Nome único para cada arquivo (timestamp)
- ✅ Resposta JSON estruturada
- ✅ Tratamento de erros
- ✅ Criação automática do diretório uploads
- ✅ Dois endpoints para flexibilidade

## Exemplo completo de uso

```bash
# 1. Inicie o servidor
go run main.go

# 2. Em outro terminal, faça o upload
# Método JSON (recomendado):
curl -X POST http://localhost:8080/v1/user/avatar \
  -H "Content-Type: application/json" \
  -d '{"filename":"avatar.png","content":"iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAACklEQVR4nGMAAQAABQABDQottAAAAABJRU5ErkJggg=="}'

# Método multipart (alternativo):
curl -X POST http://localhost:8080/v1/user/avatar/multipart \
  -F "avatar=@test_image.png"

# 3. Verifique o arquivo salvo
ls -l uploads/
```

## Sobre o Quick Framework

Quick é um framework web minimalista e rápido para Go, criado por [@jeffotoni](https://github.com/jeffotoni). Ele oferece uma API simples e intuitiva para criar aplicações web.

Neste exemplo, demonstramos como lidar com uploads de arquivos usando duas abordagens:
1. **JSON com base64**: Método mais compatível com o Quick, usando o `BodyParser` nativo
2. **Multipart**: Método alternativo que pode ter limitações dependendo da versão do Quick

## Licença

MIT
