This project is a simple URL shortener written in Go.
[![Test](https://github.com/farpat/go-url-shortener/workflows/Test/badge.svg)](https://github.com/farpat/go-url-shortener/actions)

# API
## List all short URLs
```sh
GET /api/urls
```

## Get a short URL by slug
```sh
GET /api/urls/{slug}
```

## Create a new short URL
```sh
POST /api/urls
{
    "url": "https://example.com"
}
```

## Delete a short URL
```sh
DELETE /api/urls/{slug}
```

## Installation

```sh
git clone https://github.com/farpat/go-url-shortener.git
cd go-url-shortener
cp .env.example .env
make install
```

## Usage

```sh
make update-certificates
make run
```
and watch the instructions

If you want to debug the application with VSCode, you can use the following configuration:
```json
// .vscode/launch.json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Debug main.go (public)",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}/public/main.go",
      "envFile": "${workspaceFolder}/.env"
    }
  ]
}
```
