# devto-go
[![Go Report Card](https://goreportcard.com/badge/github.com/ShiraazMoollatjie/devtogo?style=flat-square)](https://goreportcard.com/report/github.com/ShiraazMoollatjie/devtogo)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/ShiraazMoollatjie/devtogo)

devto-go is a REST API Wrapper for the dev.to api written in go.

# Usage

Import the package into your go file:

```go
import "github.com/ShiraazMoollatjie/devtogo"
```

Thereafter, create a client and specify your API token:
```go
cl := devtogo.NewClient(devtogo.WithApiKey("MY_API_KEY"))
```

It is also possible to not use an API key for anonymous operations.

## Retrieving articles
To retrieve a list of articles, simply use the `GetArticles` function:
```go
articles, err := cl.GetArticles(devtogo.Defaults())
```
It is also possible for us to add query parameters. For example, it's useful to retrieve articles for a specific `tag`.
The way to do this would be:
```go
al, err := cl.GetArticles(devtogo.Arguments{
		"tag": "go",
	})
```

To retrieve a single article, you need to specify the `article id`:
```go
article, err := client.GetArticle("167919")
```

## Retrieving your own articles

It is possible to retrieve your own articles using this API wrapper. There are four endpoints for this:

`GetMyArticles` returns all the articles created by you. `GetAllMyArticles` does the same thing.
```go
al, err := cl.GetMyArticles(devtogo.Defaults())
	if err != nil {
		panic(err)
	}
```

`GetMyPublishedArticles` returns all your published articles: 
```go
al, err := cl.GetMyPublishedArticles(devtogo.Defaults())
	if err != nil {
		panic(err)
	}
```

`GetMyUnpublishedArticles` returns all your draft articles.
```go
al, err := cl.GetMyUnpublishedArticles(devtogo.Defaults())
	if err != nil {
		panic(err)
	}
```

## Create a post
To create a post, use the `CreateArticle`:
```go
np, err := cl.CreateArticle(devtogo.CreateArticle{
  Title:        "My new dev.to post",
  Tags:         []string{"go"},
  BodyMarkdown: "my long markdown article that is preferably read from a file",
})
```
This will create a post with a title, tags and some content. We can also use the `Published` field to automatically
publish articles to dev.to.

## Update an article
Articles can be updated using the  `UpdateArticle` function:
```go
ua, err := cl.UpdateArticle(np.ID, devtogo.CreateArticle{
		Title:        "My updates dev.to post using the API",
		BodyMarkdown: "my new updated content",
		Published:    true,
	})
```

# Reference
https://docs.dev.to/api/