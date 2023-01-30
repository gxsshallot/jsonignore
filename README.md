# jsonignore

[![status](https://github.com/gaoxiaosong/jsonignore/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gaoxiaosong/jsonignore/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gaoxiaosong/jsonignore/branch/master/graph/badge.svg?token=AOXNUDXAS7)](https://codecov.io/gh/gaoxiaosong/jsonignore)
[![gover](https://img.shields.io/badge/Go-v1.2+-blue)](https://go.dev/)
[![godoc](https://pkg.go.dev/badge/github.com/gaoxiaosong/jsonignore?status.svg)](https://pkg.go.dev/github.com/gaoxiaosong/jsonignore)
[![apache](https://img.shields.io/badge/License-Apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[中文文档](README_cn.md)

This package is used to ignore some fields in object json string.

Develop for logging module. When logging, some long value or secret information should be hidden.

Paramaters:

* Ignore mode: a string to replace for ignore fields. If it is `IgnoreModeDelete`, delete the ignore fields.
* Ignore field format：
  * "k1": means data["k1"]
  * "k1.k2": means data["k1"]["k2"]
  * "k1~.k2": "k1" is a json string. delete "k2" in it after unmarshal "k1"
  * "k1*.k2": "k1" is an array. delete "k2" for each object in array.

You can see [Test Cases](process_test.go) to know how to call the method.

Examples:

```go
import "github.com/gaoxiaosong/jsonignore"

// input = {"a":1}
// result = {"a":"-"}
result := jsonignore.ProcessObject(input, "-", []string{"a"})

// input = {"a":{"a1": 1}}
// result = {"a":{}}
result := jsonignore.ProcessObject(input, IgnoreModeDelete, []string{"a.a1"})

// input = {"a":"{\"a1\":1}"}
// result = {"a":"{\"a1\":\"-\"}"}
result := jsonignore.ProcessObject(input, "-", []string{"a~.a1"})

// input = {"a":[{"a1":1},{"a2":2}]}
// result = {"a":[{},{"a2":2}]}
result := jsonignore.ProcessObject(input, IgnoreModeDelete, []string{"a*.a1"})
```
