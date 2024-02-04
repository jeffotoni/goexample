	SECTION .data   ; Declarando a seção de dados
msg: db "Hello World, assembly is love!", 0xA  ; Declaração da string e do caractere de quebra de linha

SECTION .text   ; Declarando a seção de código
global _start  ; Marcando o local de início do programa

_start:         ; Etiqueta para o ponto de início do programa
    mov eax, 4   ; Codigo syscall para escrever em STDOUT (1)
    mov ebx, 1   ; Descritor de arquivo (STDOUT)
    lea ecx, [msg]    ; Endereço da string "Hello World!"
    mov edx, 13   ; Tamanho da string
    int 0x80     ; Chamada a sistema (kernel)

    mov eax, 1   ; Codigo syscall para encerrar o programa (0 = exit code)
    int 0x80     ; Chamada a sistema (kernel)
