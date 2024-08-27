# Loan-Tracking-API

This repository contains the source code for a Loan Tracker API developed in Go (Golang). The API is designed to manage and track loans for users, providing essential features like loan creation, tracking outstanding debt, calculating interest, and managing repayment schedules.

## Key Features:

- User Management: Register, authenticate, and manage users, ensuring secure access to loan tracking features.
- Loan Tracking: Create and manage loans, including details such as loan amount, interest rate, term, and repayment status.
- Debt Monitoring: Track outstanding debt and total loan amounts for each user, providing up-to-date financial information.
- Automated Calculations: Automatically calculate interest and outstanding balances, keeping loan details accurate over time.
- Robust Validation: Ensure data integrity with comprehensive validation rules on all user and loan data.
- Secure Data Handling: Passwords and sensitive data are securely stored, following best practices for encryption and data protection.

## Tech Stack:

- Golang: The core language used for developing the API, leveraging its performance and concurrency features.
- Gin Framework: A lightweight and fast HTTP web framework used to build the API endpoints.
- MongoDB: The database used to store user and loan information, providing flexibility and scalability.
- Validator: A package used for validating input data, ensuring that all API requests meet the required criteria.

## Project Structure:

- `cmd/`: Contains the main entry point of the application.
- `api/`: Handles the API routes, controllers, and middleware.
- `domain/`: Contains the domain models, representing the core business entities such as User and Loan.
- `repository/`: Manages data persistence and retrieval, providing an abstraction layer over MongoDB interactions.
- `usecase/`: Implements the business logic and core functionalities of the application.
- `bootstrap/`: Initializes the application, including setting up the database connection and environment configurations.

## Getting Started:

1. Clone the Repository: `git clone git@github.com:Nahom-Derese/Loan-Tracking-API.git`
2. Install Dependencies: Run `go mod tidy` to install the required Go packages.
3. Setup Environment Variables: Configure your `.env` file using the provided `.env.example`
4. Run the Application: Start the API by running `go run cmd/main.go`.
5. Access the API: The API will be accessible at `http://localhost:8081`.

## Contributing:

We welcome contributions! Please read the `CONTRIBUTING.md` file for guidelines on how to contribute to the project.

- Postman Doc [https://documenter.getpostman.com/view/23769577/2sAXjGduck]
