# go-storedcounter

A simple thread safe persisted counter

## Table of Contents

* [Description](./README.md#description)
* [Install](./README.md#install)
* [Usage](./README.md#usage)
* [License](./README.md#license)

## Description

The module provides a simple counter on top of a go-ipfs-datastore interface

## Install

```
go get github.com/filecoin-project/go-storedcounter
```

## Usage

Create/load a counter for a datastore and key:

```golang
import(
  "github.com/ipfs/go-datastore"
  "github.com/filecoin-project/go-storedcounter"
)

var ds datastore.Datastore
var name datastore.Key

storedCounter := storedcounter.New(ds, name)
```

Get the next value (will start at 0):

```golang
next := storedcCounter.Next()
```

## License

Dual-licensed under [MIT](https://github.com/filecoin-project/go-statemachine/blob/master/LICENSE-MIT) + [Apache 2.0](https://github.com/filecoin-project/go-statemachine/blob/master/LICENSE-APACHE)
