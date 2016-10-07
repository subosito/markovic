# marković

[![Build Status](https://travis-ci.org/subosito/markovic.svg?branch=master)](https://travis-ci.org/subosito/markovic)
[![GoDoc](https://godoc.org/github.com/subosito/markovic?status.svg)](https://godoc.org/github.com/subosito/markovic)

Go wrapper for [libcmark](https://github.com/jgm/cmark), a [CommonMark](http://commonmark.org/) parsing and rendering library.

## Installation

```
$ go get github.com/subosito/markovic
```

## Usage

```go
r := strings.NewReader("# Hello")
s := markovic.HTML(r)

fmt.Println(s) // Output: <h1>Hello</h1>\n
```

Supported formats:

- `markovic.HTML`: render markdown as HTML document
- `markovic.XML`: render markdown as XML document
- `markovic.Man`: render markdown as GNU roff (groff) man page
- `markovic.CommonMark`: render markdown as CommonMark document
- `markovic.LaTeX`: render markdown as LaTeX document

## Alternatives

- [go-commonmark](https://github.com/rhinoman/go-commonmark): While both are wrapper for libcmark, each of them use different paradigm. go-commonmark is full-featured wrapper, while marković tries to be very simple. 
- [blackfriday](https://github.com/russross/blackfriday): Pure Go markdown parser, very popular. Not fully CommonMark compatible yet.
- [mmark](https://github.com/miekg/mmark): Another pure Go markdown parser, originally forked from blackfriday. Adds additional features.

## Benchmarks

```
$ go test -bench=.                                                 
BenchmarkHTML_Markovic-4          500000              2868 ns/op                            
BenchmarkHTML_BlackFriday-4       500000              2562 ns/op                            
BenchmarkHTML_MMark-4             500000              2375 ns/op                            
PASS                                                                                        
ok      github.com/subosito/markovic    4.011s
```
