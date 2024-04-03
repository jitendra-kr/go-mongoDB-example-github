# Go + MongoDB API with Echo (Simple & Efficient)

This is a simple go example/sample application demonstrating how to perform CRUD (Create, Read, Update, Delete) operations using GoLang and MongoDB, with the [Echo](https://echo.labstack.com) framework.

## Prerequisites

Before running this application, ensure you have the following installed:

- Go (Golang): [Installation Guide](https://golang.org/doc/install)
- MongoDB: [Installation Guide](https://docs.mongodb.com/manual/installation/)
- Sample Data [Random String Generator](https://www.fireboxtools.com/text-tools/random-string-generator) 

## Setup

1. Clone this repository:

    ```bash
    git clone <repository_url>
    ```

2. Install Go dependencies:

    ```bash
    go mod tidy
    ```

3. Start MongoDB server.

4. Update `.env` file with your MongoDB connection details.

## Usage

### Running the Application

To run the application, execute:

    go run main.go
    or
    make run

# Example REST API Documentation

This is a simple REST API documentation demonstrating various endpoints for CRUD (Create, Read, Update, Delete) operations.

## Endpoints

### Create a New Entry

- **URL:** `/user`
- **Method:** POST
- **Description:** Creates a new entry with provided data.
- **Request Body:** JSON
    ```json
    {
        "name": "Sample name",
        "email": "Sample email"
    }
    ```
- **Success Response:**
    - **Code:** 201 Created
    - **Content:** 
        ```json
        {
            "id": "660cf7c159659cdb62b1b182",
        }
        ```

### Read Entry

- **URL:** `/user/:id`
- **Method:** GET
- **Description:** Retrieve entry by user ID.
- **Success Response:**
    - **Code:** 200 OK
    - **Content:**
        ```json
                {
                    "ID": "660cfb20e7bc4042db0dba96",
                    "name": "Sample name",
                    "email": "Sample email"
                }
        ```

### Delete an Entry

- **URL:** `/user/:id`
- **Method:** DELETE
- **Description:** Deletes an entry by ID.
- **URL Params:**
    - `id` (string): The ID of the entry to delete.
- **Success Response:**
    - **Code:** 204 No Content
    - **Description:** Indicates that the entry was successfully deleted with no content returned.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

    
