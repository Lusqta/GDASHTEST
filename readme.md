````markdown
# 🌦️ GDASH Weather Monitor 2025

> Solução Full-Stack Distribuída para Monitoramento Climático e Análise de Dados via IA.

Este repositório contém a implementação completa do desafio técnico GDASH 2025/02. O sistema utiliza uma **Arquitetura Orientada a Eventos (EDA)** para garantir alta disponibilidade, resiliência e desacoplamento entre a coleta, processamento e exibição de dados meteorológicos.

## 🏛️ Arquitetura da Solução

O sistema foi desenhado para suportar picos de carga e garantir a integridade dos dados através de filas de mensagens.

```mermaid
graph LR
    A[🐍 Python Collector] -->|JSON Metrics| B(🐰 RabbitMQ)
    B -->|Consumo| C[🐹 Go Worker]
    C -->|HTTP Post| D[🦁 NestJS API]
    D <-->|Leitura/Escrita| E[(🍃 MongoDB)]
    F[⚛️ React Dashboard] <-->|REST API| D

    style A fill:#3776ab,color:#fff
    style B fill:#ff6600,color:#fff
    style C fill:#00add8,color:#fff
    style D fill:#e0234e,color:#fff
    style E fill:#47a248,color:#fff
    style F fill:#61dafb,color:#000
````

### Fluxo de Dados

1.  **Coleta (Python):** Busca dados meteorológicos (Open-Meteo) periodicamente e publica na fila.
2.  **Mensageria (RabbitMQ):** Garante a persistência e entrega assíncrona das mensagens.
3.  **Processamento (Go):** Worker de alta performance consome a fila e normaliza os dados.
4.  **Core (NestJS):** Gerencia regras de negócio, autenticação (JWT), persistência e insights de IA.
5.  **Apresentação (React):** Dashboard interativo com gráficos, alertas e exportação de dados.

## 🛠️ Stack Tecnológica

| Camada | Tecnologias |
| :--- | :--- |
| **Frontend** | React, Vite, TailwindCSS, shadcn/ui, Lucide Icons |
| **Backend** | NestJS, TypeScript, Mongoose, JWT |
| **Worker** | Golang 1.24, AMQP |
| **Ingestão** | Python 3.14, Requests, Pika |
| **Banco de Dados** | MongoDB (NoSQL) |
| **Infraestrutura** | Docker, Docker Compose |

## ✨ Funcionalidades Principais

  * ✅ **Monitoramento em Tempo Real:** Exibição de temperatura, umidade e vento com atualização automática.
  * ✅ **Insights de IA:** Análise automática de tendências climáticas e alertas.
  * ✅ **Gestão de Usuários:** Sistema de login seguro e CRUD de usuários.
  * ✅ **Exportação de Dados:** Geração de relatórios em CSV e XLSX.
  * ✅ **Resiliência:** Sistema de retry automático no Worker Go e DLQ no RabbitMQ.

## 📂 Estrutura do Projeto

```bash
/
├── backend/            # API NestJS (Core do sistema)
├── frontend/           # Aplicação React (Dashboard)
├── queue-worker/       # Microsserviço em Go (Consumidor)
├── weather-collector/  # Script Python (Produtor)
└── docker-compose.yml  # Orquestração completa
```

## 🚀 Execução

Todo o ambiente é conteinerizado. Para subir a stack completa:

```bash
docker-compose up -d
```

### Acessos

  * **Dashboard:** `http://localhost:5173`
  * **API:** `http://localhost:3000`
  * **RabbitMQ Manager:** `http://localhost:15672`

-----

*Projeto desenvolvido para o Processo Seletivo GDASH 2025/02.*

```
```