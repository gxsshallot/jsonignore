# jsonformat

[![status](https://github.com/gaoxiaosong/jsonformat/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/gaoxiaosong/jsonformat/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/gaoxiaosong/jsonformat/branch/master/graph/badge.svg?token=AOXNUDXAS7)](https://codecov.io/gh/gaoxiaosong/jsonformat)
[![gover](https://img.shields.io/badge/Go-v1.2+-blue)](https://go.dev/)
[![godoc](https://pkg.go.dev/badge/github.com/gaoxiaosong/jsonformat?status.svg)](https://pkg.go.dev/github.com/gaoxiaosong/jsonformat)
[![apache](https://img.shields.io/badge/License-Apache%202-blue.svg)](https://opensource.org/licenses/Apache-2.0)

这个包是用来忽略Json序列化结果的一些字段的。

主要用于日志模块，因为日志记录时，某些Json结果中包含非常长的字符串、需要保密的隐私信息等，这些需要被忽略输出。

参数：

* 忽略模式：如果是`IgnoreModeDelete`，则删除忽略字段，否则替换忽略字段。
* 忽略字段格式：
  * "k1": 表示 data["k1"]
  * "k1.k2": 表示 data["k1"]["k2"]
  * "k1~.k2": "k1"对应一个Json字符串，解码后删除"k2"字段
  * "k1*.k2": "k1"是一个数组，在数组的每个元素对象中删除"k2"字段

可以参照[Test Cases](process_test.go)了解如何使用这些方法。

样例：

```go
import "github.com/gaoxiaosong/jsonformat"

// input = {"a":1}
// result = {"a":"-"}
result := jsonformat.ProcessObject(input, "-", []string{"a"})

// input = {"a":{"a1": 1}}
// result = {"a":{}}
result := jsonformat.ProcessObject(input, IgnoreModeDelete, []string{"a.a1"})

// input = {"a":"{\"a1\":1}"}
// result = {"a":"{\"a1\":\"-\"}"}
result := jsonformat.ProcessObject(input, "-", []string{"a~.a1"})

// input = {"a":[{"a1":1},{"a2":2}]}
// result = {"a":[{},{"a2":2}]}
result := jsonformat.ProcessObject(input, IgnoreModeDelete, []string{"a*.a1"})
```
