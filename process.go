package jsonformat

import (
	"encoding/json"
	"strings"
)

// Ignore general object.
// obj: must be json serializable.
func ProcessObject(
	obj interface{},
	ignoreMode string,
	ignoreFields []string,
) string {
	str, _ := json.Marshal(obj)
	return string(ProcessString(str, ignoreMode, ignoreFields))
}

// Ignore json string.
func ProcessString(
	jsonStr []byte,
	ignoreMode string,
	ignoreFields []string,
) []byte {
	// no fields set
	if len(ignoreFields) == 0 {
		return jsonStr
	}
	// do nothing when unmarshal failed
	var data map[string]interface{}
	if err := json.Unmarshal(jsonStr, &data); err != nil {
		return jsonStr
	}
	// process fields
	for _, ignoreField := range ignoreFields {
		fields := strings.Split(ignoreField, ".")
		processField(data, ignoreMode, fields)
	}
	// marshal to return
	if str, err := json.Marshal(data); err != nil {
		return jsonStr
	} else {
		return str
	}
}
