package main

import "strconv"

func InterfaceToInt64(x interface{}) int64 {
	if x == nil {
		return 0
	}

	switch x.(type) {
	case string:
		res, _ := strconv.ParseInt(x.(string), 10, 64)
		return res
	case float64:
		res := int64(x.(float64))
		return res
	case int:
		res := int64(x.(int))
		return res
	}
	return 0
}
