# jsonformat

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
