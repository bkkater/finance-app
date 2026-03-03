# 🏦 F$ Compound Interest RPC System

Este projeto consiste em uma solução de **Chamada de Procedimento Remoto (RPC)** desenvolvida em **Go** para a empresa financeira hipotética **F$**. O sistema calcula juros compostos de forma distribuída, garantindo precisão matemática, transparência de auditoria (logs de IP e tempo) e alta performance em ambiente Windows.

## 🚀 Tecnologias e Recursos

- **Linguagem:** Go (Golang) v1.25.0
- **Protocolo de Comunicação:** TCP (Transmission Control Protocol)
- **Middleware:** `net/rpc` nativo do Go
- **Serialização:** `encoding/gob` (Binário nativo do Go)
- **Concorrência:** Goroutines para suporte a múltiplos clientes simultâneos
- **Arquitetura:** Baseada em princípios SOLID e Inversão de Dependência (DIP)

## 📂 Estrutura do Projeto

A organização modular do diretório separa as responsabilidades de rede, contrato e lógica de negócio:

```text
finance-app/
├── cmd/                # Pontos de entrada (Executáveis)
│   ├── client/         # Interface do Investidor (CLI)
│   └── server/         # Configuração de Rede e Handler RPC
├── pkg/                # Código compartilhado
│   └── rpctypes/       # Estruturas Request/Response (Gob)
├── services/           # Regras de Negócio (Domínio)
│   └── calc/           # Core: Cálculo de Juros Compostos
└── go.mod              # Definição do Módulo Go
```

## 🛠️ Como Executar (Windows)

Certifique-se de que o ambiente Go está configurado em seu terminal (PowerShell ou CMD). É recomendável utilizar dois terminais distintos para observar a comunicação em tempo real.

### 1. Iniciar o Servidor

O servidor registra o serviço e aguarda conexões na porta TCP `1234`. No diretório raiz do projeto, execute:

```powershell
go run ./cmd/server/
```

Requisito F$: O servidor exibirá no console o IP do cliente conectado e o tempo exato (em microssegundos) gasto para processar o cálculo, garantindo a transparência exigida pela auditoria financeira.

### 2. Executar o Cliente

Em um novo terminal, execute o cliente para realizar as consultas e inserção de dados:

```powershell
go run ./cmd/client/
```

## 📊 Exemplo de Validação

Ao interagir com o cliente, utilize os dados de teste abaixo para validar a precisão do sistema e a integridade da comunicação RPC via protocolo **Gob**:

| Campo               | Valor de Entrada | Descrição                                     |
| :------------------ | :--------------- | :-------------------------------------------- |
| **Capital Inicial** | 5000.00          | Valor base do investimento solicitado pela F$ |
| **Taxa Mensal**     | 0.01             | Representa a taxa de 1% ao mês                |
| **Meses**           | 5                | Período de capitalização do investimento      |

**Resultado esperado no Cliente:**

- Montante Final: **R$ 5.255,06**
- Detalhamento mês a mês de todos os juros auferidos, garantindo a transparência para o investidor.

## 📜 Histórico de Commits

O desenvolvimento seguiu o padrão **Conventional Commits** para garantir um histórico semântico e rastreabilidade total das funcionalidades implementadas:

- `feat(project): initialize go module`
- `feat(rpctypes): define request and response structs`
- `feat(services): implement compound interest logic`
- `feat(server): implement rpc handler`
- `feat(server): setup tcp listener and connection`
- `feat(client): implement cli interface and remote procedure call`

---

**Desenvolvido para o Projeto de Disciplina de Sistemas Distribuídos | 2026**
