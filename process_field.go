package jsonignore

import "encoding/json"

// Process json object fields recursively
func processField(
	data map[string]interface{},
	ignoreMode string,
	ignoreSplittedField []string, // ignore field format split by '.'
) {
	// no field
	if len(ignoreSplittedField) == 0 {
		return
	}
	// get a field to process
	cur := ignoreSplittedField[0]
	if len(cur) == 0 {
		return
	}
	isString := cur[len(cur)-1] == '~'
	isArray := cur[len(cur)-1] == '*'
	if isString || isArray {
		cur = cur[:len(cur)-1]
	}
	v, ok := data[cur]
	if !ok {
		return
	}
	// if only one level remains, delete or replace directly
	if len(ignoreSplittedField) == 1 {
		if ignoreMode == IgnoreModeDelete {
			delete(data, cur)
		} else {
			data[cur] = ignoreMode
		}
		return
	}
	// process field value recursively
	switch v := v.(type) {
	case map[string]interface{}:
		processField(v, ignoreMode, ignoreSplittedField[1:])
	case []interface{}:
		if isArray {
			for _, vItem := range v {
				switch vItem := vItem.(type) {
				case map[string]interface{}:
					processField(vItem, ignoreMode, ignoreSplittedField[1:])
				default:
					continue
				}
			}
		}
	case string:
		if isString {
			var innerData map[string]interface{}
			if e := json.Unmarshal([]byte(v), &innerData); e != nil {
				return
			}
			processField(innerData, ignoreMode, ignoreSplittedField[1:])
			if newStr, e := json.Marshal(innerData); e == nil {
				data[cur] = string(newStr)
			}
		}
	}
}
