# gigasite

https://gigasite.tends.to

the answer to the question, "How long until we see webpages in the gigabytes?"

start server:

```
$ go run .
```

gigasite served at port 1024. visit on a web browser that supports gzip encoding or

```
$ curl -sH 'Accept-Encoding: gzip' localhost:1024 | gunzip -
```
