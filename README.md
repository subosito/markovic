# marković

A [CommonMark](http://commonmark.org/) parser.

## Benchmarks

```
$ go test -bench=.                                                 
BenchmarkHTML_Markovic-4          500000              2868 ns/op                            
BenchmarkHTML_BlackFriday-4       500000              2562 ns/op                            
BenchmarkHTML_MMark-4             500000              2375 ns/op                            
PASS                                                                                        
ok      github.com/subosito/markovic    4.011s
```
