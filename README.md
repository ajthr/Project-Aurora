# Project Aurora - Microservices API Architecture Implementation in Go

**Disclaimer: This project has been created as a personal project to explore and learn microservices architecture and implementation in-depth. It is important to note that, due to its educational nature, the project may or may not align with the industry standards expected in production systems.**

**Note: This project is currently under development and is highly unstable.**

## Overview

Project Aurora is an ongoing initiative crafted for the purpose of self-learning microservices architecture in-depth. The project comprises the following microservices:

1. **Identity Service:**
   - **Purposes:** 
        - Verification of required credentials (Bearer token) in each request to access secured endpoints.
        - User authentication using email OTP and Google OAuth2 methods, with the issuance of credentials (Bearer token) for accessing secured endpoints.
        - User account management

## Technology Stack

- **Go**
- **Docker**
- **Nginx**

### Build and Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ajthr/project-aurora.git
   cd project-aurora
   ```

2. **Build and run the Docker containers:**

   ```bash
   docker-compose up --build
   ```

3. **Run Tests**

    ```bash
    sh test.sh
    ```

## License

This project is licensed under the [LICENSE](MIT License).
