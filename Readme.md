# Readme

## Project Overview
This repository contains the code and resources for the SharingVision Test project. The goal of this project is to create and maintain a Article platform.

## Features
- Article post creation and management

## Project Structure
- **Backend**: The `backend` folder contains the API built with Go (Golang).
- **Frontend**: The `frontend` folder contains the user interface built with React.

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/erwan690/sharing-vision.git
    ```
2. Navigate to the project directory:
    ```bash
    cd SharingVision-Blog
    ```

### Backend Setup
1. Navigate to the `backend` folder:
    ```bash
    cd backend
    ```
2. Build the Go application:
    ```bash
    go build
    ```
3. Run the API server:
    ```bash
    ./backend
    ```

### Frontend Setup
1. Navigate to the `frontend` folder:
    ```bash
    cd frontend
    ```
2. Install dependencies:
    ```bash
    npm install
    ```
3. Start the development server:
    ```bash
    npm start
    ```

## Documentation

The `docs` folder in the `backend` directory contains detailed documentation for the API. It includes the following:

1. **API Endpoints**: A comprehensive list of all available API endpoints, their methods, and expected inputs/outputs.
2. **Error Codes**: A guide to the error codes returned by the API and their meanings.
3. **Examples**: Sample requests and responses for common use cases.

To access the documentation, navigate to the `docs` folder in the `backend` directory:
```bash
cd backend/docs
```

Alternatively, you can view the API documentation in your browser by accessing the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

## Usage
- Access the API server at `http://localhost:8080` (default port for the backend).
- Access the frontend application at `http://localhost:3000`.



## Contact
For questions or feedback, please contact [erwanakse@gmail.com].