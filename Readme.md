# GO Concurrency Project

This project demonstrates the use of concurrency in Go, focusing on generating and posting user data to a PostgreSQL database.

## Building the Docker Image

1. Ensure you have Docker installed on your machine.
2. Clone this repository to your local machine.
3. Navigate to the project directory in your terminal.
4. Go to `build` directory.

To build the Docker image, use the following command:

```bash
docker build -t go-postgres .
```

This command will build the Docker image using the Dockerfile provided in the repository. The `-t` flag is used to tag the image with the name `go-postgres`, but you can replace it with any other name you prefer.

## Running the Docker Container

To run the Docker container using the image you just built, use the following command:

```bash
docker run -d -p 5432:5432 go-postgres
```

This command will start a new Docker container in detached mode (`-d` flag) and map port 5432 on your local machine to port 5432 in the container, allowing you to access the PostgreSQL database. The container will be created from the `go-postgres` image.

## Project Structure

- `internal/database`: Contains the database package, which handles database initialization and data clearing.
- `internal/errors`: Contains the errors package, which provides a function for checking and handling errors.
- `internal`: Contains the main functionality of the project, including generating and posting user data.

## How It Works

1. The `Generate` function in `internal` generates user data and sends it to a channel.
2. The `Worker2` function in `internal` receives data from the channel and either stores it in memory or posts it to the database periodically.
3. The `database` package initializes the database connection and provides a function for clearing the database.
