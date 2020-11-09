package goutils

import (
	"bytes"

	"github.com/penglongli/go-utils/json"
)

// DeepCopy used to copy object v1 to v2
func DeepCopy(v1, v2 interface{}) error {
	bs, err := json.Marshal(v1)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, v2)
}

// CompareObj used to compare between v1 and v2.
// If the bytes of them is same, return true.
func CompareObj(v1, v2 interface{}) (result bool, err error) {
	bs1, err := json.Marshal(v1)
	if err != nil {
		return result, err
	}
	bs2, err := json.Marshal(v2)
	if err != nil {
		return result, err
	}
	if bytes.Compare(bs1, bs2) != 0 {
		return false, nil
	}
	return true, nil
}
