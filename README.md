# document-database-from-scratch

- Writing a document database from scratch in Go: Lucene-like filters and indexes from https://notes.eatonphil.com/documentdb.html


## Example

Post document

```
$ curl -X POST -H 'Content-Type: application/json' -d '{"name": "Kevin", "age": "45"}' http://localhost:8080/docs
{"body":{"id":"9b7cecae-d88e-43ee-bcdb-db54d9828a84"},"status":"ok"}
```
