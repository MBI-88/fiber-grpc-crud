# grpc-crud

This project, `grpc-crud`, is a demonstration of a CRUD (Create, Read, Update, Delete) operation using gRPC in Go. It showcases how to implement a simple user management system where users can be created, updated, deleted, and fetched using gRPC calls.

## Features
- **Create User**: Allows the creation of a new user with details such as name, password, website, DNI (Document Number Identification), address, and phone.
- **Update User**: Permits updating an existing user's information.
- **Delete User**: Enables the deletion of a user from the system.
- **Get Users**: Fetches the list of users from the system.

## Requirements
- Go 1.15 or higher
- gRPC
- Protocol Buffers

## Installation
To set up the project locally, follow these steps:
1. Clone the repository to your local machine.
2. Navigate to the cloned directory.
3. Install the required dependencies by running `go mod tidy`.

## Usage
To use the application, you need to start the server and then run the client with the desired operation.

### Starting the Server
Navigate to the server directory and run:

