package jsonignore

import (
	"encoding/json"
	"fmt"
	"testing"
)

type KV map[string]interface{}
type FD []string

func TestJsonOutput(t *testing.T) {
	D, R := IgnoreModeDelete, "-"
	testList := []struct {
		Origin       KV
		IgnoreMode   string
		IgnoreFields FD
		Result       KV
	}{
		{KV{"a": 0}, D, FD{}, KV{"a": 0}},
		{KV{"b": "1"}, D, FD{"b"}, KV{}},
		{KV{"b": 1}, R, FD{"b"}, KV{"b": R}},
		{KV{"b": 1}, D, FD{"c"}, KV{"b": 1}},
		{KV{"c": KV{"c1": "21", "c2": 22}}, D, FD{"c.c1"}, KV{"c": KV{"c2": 22}}},
		{KV{"c": KV{"c1": 21, "c2": 22}}, R, FD{"c.c1"}, KV{"c": KV{"c1": R, "c2": 22}}},
		{KV{"c": KV{"c1": 21, "c2": 22}}, D, FD{"c.c3"}, KV{"c": KV{"c1": 21, "c2": 22}}},
		{KV{"d": "{\"d1\":\"31\",\"d2\":32}"}, D, FD{"d~.d1"}, KV{"d": "{\"d2\":32}"}},
		{KV{"d": "{\"d1\":31,\"d2\":32}"}, R, FD{"d~.d1"}, KV{"d": "{\"d1\":\"-\",\"d2\":32}"}},
		{KV{"d": "{\"d1\":31,\"d2\":32}"}, D, FD{"d~.d3"}, KV{"d": "{\"d1\":31,\"d2\":32}"}},
		{KV{"e": []KV{{"e1": "41"}, {"e2": 43}}}, D, FD{"e*.e1"}, KV{"e": []KV{{}, {"e2": 43}}}},
		{KV{"e": []KV{{"e1": 41}, {"e2": 43}}}, R, FD{"e*.e1"}, KV{"e": []KV{{"e1": R}, {"e2": 43}}}},
		{KV{"e": []KV{{"e1": 41}, {"e2": 43}}}, D, FD{"e*.e3"}, KV{"e": []KV{{"e1": 41}, {"e2": 43}}}},
	}
	for _, item := range testList {
		output := ProcessObject(item.Origin, item.IgnoreMode, item.IgnoreFields)
		resultStr, _ := json.Marshal(item.Result)
		if output != string(resultStr) {
			t.Error(fmt.Sprintf("invalid result\nresult: %s\noutput: %s", string(resultStr), output))
		}
	}
}
