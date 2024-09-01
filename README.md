# Stark Bank Hackathon Project

### Frontend Code 
- https://github.com/carlosdoki/starkbank
- https://github.com/carlosdoki/stark_atendimento

## Context

Welcome to our project for the Stark Bank Hackathon! This repository contains the code and documentation for our submission.

This project includes both the backend and the integration with AI, and it was developed using Golang. We have leveraged the power of Golang to build a robust and efficient backend system that seamlessly integrates with the postgre database. Additionally, we have incorporated AI capabilities using Google Cloud to enhance the functionality and intelligence of our solution.

Our Golang backend provides secure and efficient processing, ensuring the reliability and integrity of financial conversations. 

For the project was used Go version 1.22.

To get started with the backend and AI integration, follow these steps:

0. Install the gcloud CLI and configure google cloud credentials  

1. Clone the repository: `git clone https://github.com/julioc98/starkbank.git`
2. Install the required dependencies: `go mod download`
3. Run infra `docker compose up`
4. Run migrations `make migrate-up`
5. Build and run the application: `go run cmd/api/main.go`

If you have any questions or suggestions regarding this project, feel free to reach out to us.

## Usage

Once the application is up and running, you can access it at localhost:3000 (the end-point). Use the provided user interface to perform various financial conversations and explore the features of our project.

### Frontend Code 
- https://github.com/carlosdoki/starkbank
- https://github.com/carlosdoki/stark_atendimento

### endpoints

```
    h.r.Post("/analysts", h.CreateAnalyst)
	h.r.Get("/analysts", h.GetAllAnalysts)
	h.r.Get("/analysts/{id}", h.GetAnalystByID)
	h.r.Post("/msg", h.ResponseMsg)
```