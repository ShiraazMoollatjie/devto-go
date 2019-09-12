# devto-go
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/devtogo?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/devtogo)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/devtogo)

A REST API Wrapper for the dev.to api written in go

# Usage

To use the API, create a client:
```go
cl := NewClient()
```

## Retrieving articles

To retrieve a list of articles:

```go
articles, err := client.GetArticles()
```

To retrieve a single article, you need to specify the `article id`:

```go
article, err := client.GetArticle("167919")
```

# Reference
https://docs.dev.to/api/