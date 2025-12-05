# 📝 Go To-Do List App

![Go Version](https://img.shields.io/badge/Go-1.25-blue)
![Docker](https://img.shields.io/badge/Docker-Enabled-blue)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

A lightweight, containerized To-Do List application built with **Golang**. It features a fully automated CI/CD pipeline that builds the Docker image and deploys it to an **AWS EC2** instance.

## 🚀 Tech Stack

* **Language:** Golang (1.25.1)
* **Containerization:** Docker (Multi-stage builds)
* **CI/CD:** GitHub Actions
* **Registry:** Docker Hub
* **Deployment:** AWS EC2 (Ubuntu 22.04)

## 📂 Project Structure

```bash
├── .github/workflows/   # CI/CD Configuration
├── views/               # HTML Templates
├── Dockerfile           # Multi-stage Docker build
├── go.mod               # Go Module definitions
└── main.go              # Application entry point
```

## 🛠️ How to Run Locally

### Method 1: Standard Go
If you have Go installed on your machine:

1.  Clone the repository:
    ```bash
    git clone https://github.com/AnlongZhou/automation_fullstack.git
    ```
2.  Install dependencies:
    ```bash
    go mod download
    ```
3.  Run the application:
    ```bash
    go run cmd/main.go
    ```
4.  Visit `http://localhost:3000` in your browser.

### Method 2: Using Docker (Recommended)
You don't need Go installed if you use Docker:

1.  Build the image:
    ```bash
    docker build -t todo-app .
    ```
2.  Run the container:
    ```bash
    docker run -p 3000:3000 todo-app
    ```

## ⚙️ CI/CD Pipeline

This project uses **GitHub Actions** for automation.

1.  **Continuous Integration (CI):**
    * Triggers on every push to `main`.
    * Runs `go test` to ensure code quality.
2.  **Continuous Delivery (CD):**
    * Logs into Docker Hub.
    * Builds the Docker Image.
    * Pushes the image to Docker Hub.
3.  **Deployment:**
    * SSHs into the AWS EC2 instance.
    * Pulls the latest image.
    * Restarts the container with the new updates.

## 🔐 Secrets Configuration

To replicate this pipeline, you need the following **GitHub Secrets**:

| Secret Name | Description |
| :--- | :--- |
| `DOCKERHUB_USERNAME` | Your Docker Hub ID |
| `DOCKERHUB_TOKEN` | Docker Hub Access Token |
| `SERVER_HOST` | AWS EC2 Public IP |
| `SERVER_USER` | `ubuntu` |
| `SSH_PRIVATE_KEY` | Content of your `.pem` key file |

## 🛡️ AWS Security Setup

The EC2 instance is configured with the following **Security Group** rules:

* **Port 22 (SSH):** Restricted to specific IP (Admin access).
* **Port 3000 (App):** Open to `0.0.0.0/0` (Application access).
* **Port 80 (HTTP):** Open to `0.0.0.0/0` (Future reverse proxy).
