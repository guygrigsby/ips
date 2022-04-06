## Request IP gathering
```
goos: darwin
goarch: amd64
pkg: github.com/guygrigsby/ips
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkManyRequestHandled-16    	1000000000	         0.3234 ns/op
BenchmarkOneRequestHandled-16     	1000000000	         0.0000010 ns/op
PASS
ok  	github.com/guygrigsby/ips	8.356s
```

Storing ips is accomplished by maintaining a map of ip adresses to the count of each IP.

Retrieving the top 100 ips is done by maintaining a binary tree of pointers to the values in the above mentioned map. Then a depth first traversal is done to gather the top 100 ips with the most requests.

> :warning: While I believe the implementation hypothesis to be sound, currently there is a bug that causes the top100 to timeout. 

In retrospect, I think it would be better to use iteration rather than recursion for inserting a new value into the binary tree.
