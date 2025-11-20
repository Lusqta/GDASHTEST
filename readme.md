# 🌦️ GDASH Weather Monitor 2025

![Arquitetura](https://img.shields.io/badge/Arquitetura-EDA-blue)
![Stack](https://img.shields.io/badge/Stack-Full--Stack-purple)
![IA](https://img.shields.io/badge/Feature-IA-green)
![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-orange)
![License](https://img.shields.io/badge/Licen%C3%A7a-MIT-green)

---

Solução **Full-Stack Distribuída para Monitoramento Climático e Análise de Dados via IA**, desenvolvida para o desafio técnico **GDASH 2025/02**.

O sistema utiliza uma **Arquitetura Orientada a Eventos (EDA)** para garantir:

- Alta disponibilidade  
- Resiliência  
- Escalabilidade  
- Desacoplamento entre serviços  

Toda comunicação entre os módulos ocorre via mensageria.

---

## 🚀 Proposta do Sistema

Este projeto tem como objetivo criar um ecossistema distribuído e escalável para **monitoramento climático em tempo real**, integrando:

- Coleta automatizada de dados meteorológicos  
- Processamento assíncrono via mensageria  
- Persistência em banco NoSQL  
- Visualização em dashboard web  
- Análise com suporte de IA  

---

## 🏛️ Arquitetura da Solução

A arquitetura foi desenhada para suportar **picos de carga** e garantir a **integridade dos dados meteorológicos**.

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
```

---

## 🧩 Componentes do Sistema

| Serviço                 | Função                                 |
| ----------------------- | -------------------------------------- |
| --------                | --------                               |
| 🐍 **Python Collector** | Coleta dados climáticos em tempo real  |
| 🐰 **RabbitMQ**         | Gerencia a troca de mensagens          |
| 🐹 **Go Worker**        | Processa, valida e normaliza os dados  |
| 🦁 **NestJS API**       | API principal e camada de persistência |
| 🍃 **MongoDB**          | Banco de dados NoSQL                   |
| ⚛️ **React Dashboard**  | Interface visual para monitoramento    |

---

## 📘 Fundamentos e Padrões Utilizados

O projeto aplica conceitos modernos de arquitetura e desenvolvimento:

* Arquitetura Orientada a Eventos (EDA)
* Microserviços desacoplados
* Comunicação assíncrona via Message Broker
* API RESTful
* Armazenamento NoSQL (MongoDB)
* Processamento distribuído de dados
* Integração com IA para análise climática

---

## 🛠️ Tecnologias

* **Python**
* **Go**
* **NestJS**
* **React**
* **RabbitMQ**
* **MongoDB**
# 🌦️ GDASH Weather Monitor 2025

![Arquitetura](https://img.shields.io/badge/Arquitetura-EDA-blue)
![Stack](https://img.shields.io/badge/Stack-Full--Stack-purple)
![IA](https://img.shields.io/badge/Feature-IA-green)
![Status](https://img.shields.io/badge/Status-Em%20Desenvolvimento-orange)
![License](https://img.shields.io/badge/Licen%C3%A7a-MIT-green)

---

Solução **Full-Stack Distribuída para Monitoramento Climático e Análise de Dados via IA**, desenvolvida para o desafio técnico **GDASH 2025/02**.

O sistema utiliza uma **Arquitetura Orientada a Eventos (EDA)** para garantir:

- Alta disponibilidade  
- Resiliência  
- Escalabilidade  
- Desacoplamento entre serviços  

Toda comunicação entre os módulos ocorre via mensageria.

---

## 🚀 Proposta do Sistema

Este projeto tem como objetivo criar um ecossistema distribuído e escalável para **monitoramento climático em tempo real**, integrando:

- Coleta automatizada de dados meteorológicos  
- Processamento assíncrono via mensageria  
- Persistência em banco NoSQL  
- Visualização em dashboard web  
- Análise com suporte de IA  

---

## 🏛️ Arquitetura da Solução

A arquitetura foi desenhada para suportar **picos de carga** e garantir a **integridade dos dados meteorológicos**.

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
```

---

## 🧩 Componentes do Sistema

| Serviço                 | Função                                 |
| ----------------------- | -------------------------------------- |
| --------                | --------                               |
| 🐍 **Python Collector** | Coleta dados climáticos em tempo real  |
| 🐰 **RabbitMQ**         | Gerencia a troca de mensagens          |
| 🐹 **Go Worker**        | Processa, valida e normaliza os dados  |
| 🦁 **NestJS API**       | API principal e camada de persistência |
| 🍃 **MongoDB**          | Banco de dados NoSQL                   |
| ⚛️ **React Dashboard**  | Interface visual para monitoramento    |

---

## 📘 Fundamentos e Padrões Utilizados

O projeto aplica conceitos modernos de arquitetura e desenvolvimento:

* Arquitetura Orientada a Eventos (EDA)
* Microserviços desacoplados
* Comunicação assíncrona via Message Broker
* API RESTful
* Armazenamento NoSQL (MongoDB)
* Processamento distribuído de dados
* Integração com IA para análise climática

---

## 🛠️ Tecnologias

* **Python**
* **Go**
* **NestJS**
* **React**
* **RabbitMQ**
* **MongoDB**

---

## 📂 Estrutura Geral

```text
📁 gdash-challenge
├── weather-collector/      # Serviço de coleta (Python)
├── queue-worker/           # Processamento de fila (Go)
├── backend/                # API Principal (NestJS)
├── frontend/               # Dashboard (React)
├── docker-compose.yml      # Orquestração
└── README.md
```

---

## 🎓 Contexto

Este projeto foi desenvolvido para o desafio técnico **GDASH 2025/02**, focando em:

*   **Monitoramento Climático:** Ingestão e normalização de dados em tempo real.
*   **Inteligência Artificial:** Análise de tendências e geração de insights.
*   **Arquitetura Robusta:** Sistema distribuído, escalável e resiliente a falhas.

---
