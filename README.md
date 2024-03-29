<h1 align="center">mingo</h1>

<p align="center">
<a href="https://github.com/koki-develop/mingo/releases/latest"><img src="https://img.shields.io/github/v/release/koki-develop/mingo" alt="GitHub release (latest by date)"></a>
<a href="https://github.com/koki-develop/mingo/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/koki-develop/mingo/ci.yml?logo=github" alt="GitHub Workflow Status"></a>
<a href="https://goreportcard.com/report/github.com/koki-develop/mingo"><img src="https://goreportcard.com/badge/github.com/koki-develop/mingo" alt="Go Report Card"></a>
<a href="./LICENSE"><img src="https://img.shields.io/github/license/koki-develop/mingo" alt="LICENSE"></a>
</p>

<p align="center">
<b><i>Go also wants to be minified.</i></b>
</p>

## Contents

- [Contents](#contents)
- [Installation](#installation)
  - [Homebrew Tap](#homebrew-tap)
  - [`go install`](#go-install)
  - [Releases](#releases)
- [Usage](#usage)
  - [Example](#example)
- [LICENSE](#license)

## Installation

### Homebrew Tap

```console
$ brew install koki-develop/tap/mingo
```

### `go install`

```console
$ go install github.com/koki-develop/mingo@latest
```

### Releases

Download the binary from the [releases page](https://github.com/koki-develop/mingo/releases/latest).

## Usage

```console
$ mingo --help
Go language also wants to be minified.

Usage:
  mingo [flags] [files]...

Flags:
  -h, --help      help for mingo
  -v, --version   version for mingo
  -w, --write     write result to (source) file instead of stdout
```

### Example

```go
// main.go
package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	n := 10
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}
```

```console
$ mingo main.go
```

```go
package main;import "fmt";func fib(n int)int{if n<=1{return n};return fib(n-1)+fib(n-2)};func main(){n:=10;for i:=0;i<n;i++{fmt.Println(fib(i))}};
```

---

An example of minifying [cLive](https://github.com/koki-develop/clive):

- [Minified #128 - koki-develop/clive](https://github.com/koki-develop/clive/pull/128)

## LICENSE

[MIT](./LICENSE)
