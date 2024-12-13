![](AcneScanFinal.png)
# AcneScan Backend

This repository is the server-side component of the AcneScan mobile application, which users can read the articles and receive personalized skincare recommendations.

---

## Overview

AcneScan backend is built using **Go (Golang)** and serves as provider personalized skincare recommendations, and managing data. The backend integrates the app for articles and product recommendation, and leverages **cloud computing** infrastructure on **Google Cloud Platform (GCP)** for scalability and efficiency.

## Features

- **Acne Classification**: Uses a CNN-based machine learning model to classify acne types from images uploaded by users.
- **Personalized Recommendations**: Based on the acne classification, the backend provides skincare product recommendations using an API.
- **Cloud Infrastructure**: Deployed on GCP using services like Cloud Run, App Engine, and Docker containers for scalability and reliability.
- **Database Management**: Data is stored in a relational SQL database for efficient user and recommendation management.

---


## Technology Stack

- **Backend Language**: Go (Golang)
- **Machine Learning Model**: Convolutional Neural Networks (CNN)
- **Cloud Provider**: Google Cloud Platform (GCP)
  - **Cloud Run**: For scalable containerized services.
  - **App Engine**: For backend deployment.
  - **SQL**: For data management and storage.
  - **Docker**: For containerization of the application.
  - **Artifact Registry**: For efficient code deployment and versioning.
- **API Framework**: Custom RESTful API developed in Go
- **CI/CD**: Automated deployment using GCP services.

---

## Installation

To run the AcneScan backend locally, follow these steps:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/acnescan-backend.git
   cd acne-scan-api
2. Install Dependency
   ```bash
   go mod tidy

3. set up environment database
   
5. Run the application locally:
      ```bash
      go run main.go

## Migration
```
migrate -database "mysql://root:@tcp(localhost:3306)/acne_scan" -path internal/infrastructure/database/migrations up
```
```
migrate create -ext sql -dir internal/infrastructure/database/migrations (namafile).sql
```

## API Docs
``` https://acne-scan-api.postman.co/workspace/acne-scan-api-Workspace~ac33a853-4123-47ef-a28e-6e4ce7043bd5/collection/26211928-cad79b08-d5d8-43f4-90a1-d67840212e73?action=share&creator=26211928&active-environment=26211928-3be2084b-1e3d-4ad2-9a97-935eb2fac948
