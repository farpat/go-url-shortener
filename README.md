# Go URL Shortener
A simple URL shortener written in Go.
[![Test](https://github.com/farpat/go-url-shortener/workflows/Test/badge.svg)](https://github.com/farpat/go-url-shortener/actions)

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
make install
```

## Usage

```sh
make run
```
