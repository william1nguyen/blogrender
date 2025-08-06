# blogrender

## Overview

**blogrender** is a simple web application for rendering Markdown blog posts as HTML via API endpoints.

## Installation

Clone the repository and install dependencies:

```sh
git clone https://github.com/yourusername/blogrender.git
cd blogrender
go mod tidy
```

## Build & Run

To build the project:

```sh
make build
```

Or run directly:

```sh
make run
```

The application will start at `http://localhost:3000/`.

## API Usage

### Get List of Posts

```sh
curl -X GET http://localhost:3000/
```

Returns a list of available Markdown post filenames.

### Get HTML Content of a Post

```sh
curl -X GET "http://localhost:3000/html?name=welcome.md"
```

Returns the HTML content of the specified Markdown post.

## Testing

Run all tests:

```sh
make test
```

Run benchmark:

```sh
make benchmark
```

## Directory Structure

- `md/`: Contains Markdown post files.
- `cmd/blogrender/`: Entry point for the application.
- `internal/app/`: Core logic for post handling and HTML rendering.
