# document-database-from-scratch

- Writing a document database from scratch in Go: Lucene-like filters and indexes from https://notes.eatonphil.com/documentdb.html


## Example

Post document

```
$ curl --get -H 'Content-Type: application/json' -d '{"name": "Kevin", "age": "45"}' http://localhost:8080/docs
{"body":{"id":"9b7cecae-d88e-43ee-bcdb-db54d9828a84"},"status":"ok"}
```

Get document by id

```
$ curl http://localhost:8080/docs/9b7cecae-d88e-43ee-bcdb-db54d9828a84
{"body":{"document":{"age":"45","name":"Kevin"}},"status":"ok"}
```

Search documents

```
$ curl --get http://localhost:8080/docs
{"body":{"count":1,"documents":[{"body":{"age":"45","name":"Kevin"},"id":"9b7cecae-d88e-43ee-bcdb-db54d9828a84"}]},"status":"ok"}

$ curl --get http://localhost:8080/docs --data-urlencode 'q=name:Kevin'
{"body":{"count":1,"documents":[{"body":{"age":"45","name":"Kevin"},"id":"9b7cecae-d88e-43ee-bcdb-db54d9828a84"}]},"status":"ok"}

$ curl --get http://localhost:8080/docs --data-urlencode 'q=age:<30'
{"body":{"count":0,"documents":null},"status":"ok"}
```
