# Levenshtein Distance in Golang

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/AgFlow/levenshtein)
[![Build Status](https://travis-ci.org/AgFlow/levenshtein.png?branch=master)](https://travis-ci.org/AgFlow/levenshtein)

> Calculate levenshtein distance in Golang.

## Install

By go tool: `go get github.com/AgFlow/levenshtein`

## Usage

This uses default calculator which has cost of 1 for additions, deletions and substitutions.

```go
import github.com/AgFlow/levenshtein

levenshtein.Dist("aaa", "ab") // 2
```

You can specify different weights to increment/deletion and substitutions.

```go
levenshtein.New(1, 1).Dist("aaa", "ab") // 2
levenshtein.New(1, 2).Dist("aaa", "ab") // 3
levenshtein.New(1, 3).Dist("aaa", "ab") // 3
levenshtein.New(1, 4).Dist("aaa", "ab") // 3
levenshtein.New(2, 2).Dist("aaa", "ab") // 4
levenshtein.New(3, 2).Dist("aaa", "ab") // 5
```

If you don't care difference more than some predefined value and strings are encoded in ascii,
and also you want to compare same string to multiple different strings then there is one more performant interface you can utilize:

```go
d := levenshtein.FromBytes([]byte("Mustafa"), 2)
d.Dist([]byte("Kemal")) // 2
d.Dist([]byte("Mustfa")) // 1

d = levenshtein.FromBytes("Mustafa", 6)
d.Dist([]byte("Kemal")) // 6
d.Dist([]byte("Mustfa")) // 1
```

## LICENSE

MIT © [AgFlow](https://github.com/AgFlow)
