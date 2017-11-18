Gandi LiveDNS Go library
========================

[![GoDoc](https://godoc.org/github.com/tiramiseb/go-gandi-livedns?status.svg)](https://godoc.org/github.com/tiramiseb/go-gandi-livedns)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/tiramiseb/go-gandi-livedns/master/LICENSE)

This library interacts with [Gandi's LiveDNS API](http://doc.livedns.gandi.net/), to manage domains hosted on Gandi. This library returns some data as HTTP headers, please note those are ignored by this library.

**Gandi is currently (as of Nov. 2017) migrating on a new platform, this library is for the NEW platform.**

A simple CLI is also shipped with this library. It returns responses to the requests in JSON.

Compiling the CLI
-----------------

```
cd cmd
go build -o gandi
```
