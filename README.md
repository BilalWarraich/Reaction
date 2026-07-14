# Reaction
A reaction module built as a Go microservice with a GraphQL API.

## What this is
A small backend service showing how I structure Go microservices: clean separation between service logic and the API layer, GraphQL instead of plain REST. It basicalyy is a reaction module for a bigger application i built which had a feature of placing reaction on a post like we do on social media platforms like facebook

## Stack
Go, GraphQL

## Why it matters
Same architecture style I later used in production on the Ministry of Tourism (KSA) digital certificate platform, which has issued 100,000+ certificates and was load tested to 100,000+ concurrent requests.

## Running it


## Build

```
$ docker-compose up -d --build
```

Open <http://localhost:8050/graphiql> in your browser.
