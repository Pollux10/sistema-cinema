# 🎬 Sistema de Reserva de Cinema (CLI) - Go

Sistema de simulação de compra de ingressos de cinema via terminal (CLI), com seleção de assentos e processamento de pagamento, desenvolvido em Go.

---

## 📄 Descrição

Este projeto simula o funcionamento básico de um sistema de cinema, permitindo que o usuário visualize filmes disponíveis, escolha assentos e finalize a compra de ingressos diretamente pelo terminal.

Os dados são armazenados em um arquivo JSON, garantindo persistência mesmo após o encerramento do programa.

---

## 🔄 Funcionamento

1. O usuário inicia o programa  
2. O sistema carrega os filmes e seus respectivos assentos do arquivo `data.json`  
3. Um menu interativo é exibido  

O usuário pode escolher entre as seguintes opções:

- **Ver filmes**: lista os filmes disponíveis  
- **Comprar ingresso**: permite selecionar um filme, escolher um assento disponível e realizar o pagamento  

### 🎟️ Fluxo de compra

1. O usuário escolhe um filme  
2. O sistema exibe a matriz de assentos  
   - `O` = assento livre  
   - `X` = assento ocupado  
3. O usuário seleciona linha e coluna  
4. O sistema verifica disponibilidade  
5. Caso o assento esteja livre:
   - Ele é reservado  
   - O usuário escolhe a forma de pagamento  
6. Se o pagamento for aprovado:
   - A compra é finalizada  
   - O assento é salvo como ocupado  

---

## 🚀 Funcionalidades

- Listagem de filmes  
- Visualização de assentos  
- Seleção de assentos  
- Validação de disponibilidade  
- Simulação de pagamento  
- Persistência de dados em JSON  

---

## 🛠 Tecnologias utilizadas

- Go (Golang)  
- JSON  

---

## ▶️ Como executar

```bash
go run sistema-cinema.go
