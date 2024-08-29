# Train Ticketing System - gRPC Service

This project implements a basic train ticketing system using Golang and gRPC. The system allows users to purchase tickets, view their receipts, and manage seat allocations.

## Project Structure

- **`main.go`**: The main file containing the gRPC server implementation, including methods for purchasing tickets, viewing receipts, managing users, and modifying seat allocations.
- **`ticketing.proto`**: The Protocol Buffers file that defines the gRPC services and message types.

## Prerequisites

Before running the project, ensure you have the following installed:

- Go (version 1.16 or higher)
- Protocol Buffers Compiler (`protoc`)
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/SidharthSasikumar/train-ticket-grpc.git
   cd train-ticket-grpc

2. **Install Dependencies**

    Make sure you have the necessary Go dependencies by running:
    ```bash 
    go mod tidy

## Running the Project
To run the gRPC server:

1. Start the gRPC Server
   Run the server with the following command:
   ```bash
   go run main.go 

  The server will start listening on port 50051 by default.
2. gRPC Endpoints
   The server provides the following gRPC methods:

    * **PurchaseTicket**: Purchase a train ticket.
    * **GetReceipt**: Retrieve a user's receipt.
    * **ViewUsers**: View users and their seat allocations.
    * **RemoveUser**: Remove a user and free their seat.
    * **ModifySeat**: Modify a user's seat allocation.
   
## Testing the Project
Test cases for various functionalities are included in the project.
1. Run Tests

    Execute the tests using the following command:
   ```bash
   go test -v
   
This command will run all the tests and display detailed output in your terminal.



