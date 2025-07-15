# Dynamic NFT Metadata AVS (Othentic Example)

## Overview
This project implements an **Othentic AVS (Actively Validated Service)** for dynamic NFT metadata, where off-chain achievements (such as GitHub contributions or PlayStation trophies) are verified and used to update on-chain NFT metadata. The AVS consists of:
- **Execution Service:** Fetches off-chain data, generates a proof, and submits it to the AVS network.
- **Validation Service:** Independently fetches the same off-chain data, verifies the proof, and attests to its validity.
- **DummyAPI:** Provides static/mock off-chain data for MVP/testing.
- **Othentic Node Stack:** Aggregator and Attester nodes for consensus and on-chain finality.

This stack is containerized and orchestrated using Docker Compose for easy local development and testing.

---

## Disclaimer
This project was made for learning purposes, is not production ready, and a UI is currently in progress.

---

## What is Othentic?
[Othentic](https://docs.othentic.xyz/) is a modular, programmable AVS framework built on EigenLayer. It allows developers to create custom off-chain validation and execution logic, enabling new decentralized services and protocols with robust security and consensus.

- **Execution Service:** Runs on Performer nodes, executes tasks, and generates Proof-of-Task (PoT).
- **Validation Service:** Runs on Attester nodes, verifies PoT, and casts attestations.
- **Aggregator:** Collects attestations and submits results on-chain.

---

## Project Structure

```
📂 assignment-2
├── 📂 Execution_Service         # Task execution logic (Performer)
│   ├── 📂 config/
│   │   └── config.go            # Loads environment variables and config.
│   ├── 📂 service/
│   │   ├── metadata.go          # Fetches off-chain user data from DummyAPI.
│   │   └── processor.go         # Handles proof generation and sendTask RPC.
│   ├── 📂 utils/
│   │   └── handler.go           # HTTP handler for /task/execute endpoint.
│   ├── Dockerfile               # Dockerfile for Execution Service.
│   ├── go.mod                   # Go module definition.
│   └── main.go                  # Entry point, sets up HTTP server.
│
├── 📂 Validation_Service        # Task validation logic (Attester)
│   ├── 📂 services/
│   │   ├── validation.go        # Validates proof by fetching and hashing off-chain data.
│   │   ├── data.go              # Helper to fetch user data from DummyAPI.
│   │   ├── resp.go              # Response helpers for API.
│   │   └── error.go             # Error helpers for API.
│   ├── 📂 handler/
│   │   └── handler.go           # HTTP handler for /task/validate endpoint.
│   ├── Dockerfile               # Dockerfile for Validation Service.
│   ├── go.mod                   # Go module definition.
│   └── main.go                  # Entry point, sets up HTTP server.
│
├── 📂 Dummy_API                 # Dummy API for static/mock off-chain data
│   ├── main.go                  # Simple Go server for /api/github and /api/psn
│   └── ...                      # (other files as needed)
│
├── docker-compose.yml           # Orchestrates all services and Othentic nodes
├── .env.example                 # Example environment variables
└── README.md                    # Project documentation (this file)
```

---

## Installation & Setup

### 1. **Install Othentic CLI and Node**
```sh
npm i -g @othentic/cli
npm i -g @othentic/node
```

### 2. **Clone and Prepare the Project**
```sh
git clone https://github.com/naman1402/eigen-bootcamp-assignment-2
cd assignment-2
cp .env.example .env
# Edit .env as needed (set private keys, DUMMY_API_URL, etc.)
```

### 3. **Build and Run with Docker Compose**
```sh
docker-compose up --build
```
This will start:
- Othentic Aggregator and Attester nodes
- Execution Service (Performer)
- Validation Service (Attester)
- DummyAPI (off-chain data source)

---

## Usage

### **Execution Service**
- Exposes `POST /task/execute` (default port 4003)
- Request body:
  ```json
  {
    "taskDefinitionId": 1,
    "address": "0x123",
    "achievementType": "github"
  }
  ```
- Returns a proof of task and off-chain data.

### **Validation Service**
- Exposes `POST /task/validate` (default port 4004)
- Request body:
  ```json
  {
    "proofOfTask": "<hash>",
    "address": "0x123",
    "achievementType": "github"
  }
  ```
- Returns validation result (`data: true/false`), error, and message.

### **DummyAPI**
- Exposes `/api/github?address=...` and `/api/psn?address=...` for static user data.
- Used by both Execution and Validation services for off-chain data simulation.

---

## How It Works
1. **Task Triggered:** Performer node (Execution Service) receives a task to check an off-chain achievement.
2. **Off-chain Data Fetch:** Execution Service queries DummyAPI for user data, generates a proof (hash of data), and submits it to the AVS network.
3. **Validation:** Attester node (Validation Service) independently fetches the same data, hashes it, and compares to the submitted proof.
4. **Consensus:** If the proof matches, the task is approved and the result is finalized on-chain by the Aggregator.

---

## References
- [Othentic Documentation](https://docs.othentic.xyz/)
- [Othentic Execution Service](https://docs.othentic.xyz/main/learn/core-concepts/execution-service)
- [Othentic Validation Service](https://docs.othentic.xyz/main/learn/core-concepts/validation-service)

