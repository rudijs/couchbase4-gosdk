# Couchbase 4 Go SDK

- Start Couchbase 4
- Start Go program and enter in some JSON documents

```
❯ go run main.go
==> Enter a JSON document:
{"id":101}
Inserted document CAS is `140ed05555b70000`
==> Enter a JSON document:
{"id":102}
Inserted document CAS is `140ed0572b680000`
```

- Stop Couchbase 4
- Enter in more documents from the **same** running Go program

```
==> Enter a JSON document:
{"id":103}
ERROR: The operation has timed out.
==> Enter a JSON document:
{"id":104}
ERROR: The operation has timed out.
==> Enter a JSON document:
```

- Restart Couchbase 4
- Enter in more documents from the same running Go program

```
{"id":105}
ERROR: The operation has timed out.
==> Enter a JSON document:
```

So we can see the Go SDK client is not reconnecting.

- Exit and the start again the Go program, which will then reconnect to Couchbase 4


```
==> Enter a JSON document:
^Cexit status 2
/m/r/c/p/g/s/g/r/couchbase4-gosdk ❯❯❯ go run main.go                                                                                           ⏎
==> Enter a JSON document:
{"id":106}
Inserted document CAS is `140ed07055fc0000`
==> Enter a JSON document:
```

So now at the end of this process Couchbase 4 has documents 101, 102 and 106

Documents 103 and 104 were not upserted as Couchbase 4 was off.

Document 105 was not upserted even though Couchbase was running again.

What would be some idiomatic Go lang approaches to re-try and/or reconnect?
