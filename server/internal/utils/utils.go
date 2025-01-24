package utils

import (
	"bytes"
	"fmt"
	"runtime"
	"runtime/debug"
	"strconv"
)

func WrapError(err error) error {
	stack := string(debug.Stack())
	return fmt.Errorf("%w|%s", err, stack)
}

func GoroutineID() uint64 {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := bytes.Fields(buf[:n])[1]
	id, _ := strconv.ParseUint(string(idField), 10, 64)
	return id
}

type AnyJson map[string]interface{}

func (aj AnyJson) GetString(key string) string {
	if aj == nil {
		return ""
	}
	s, _ := aj[key].(string)
	return s
}

func (aj AnyJson) GetStringList(key string) []string {
	if aj == nil {
		return nil
	}
	list, _ := aj[key].([]interface{})
	result := make([]string, len(list))
	for i, v := range list {
		result[i] = v.(string)
	}
	return result
}

func (aj AnyJson) GetInt64List(key string) []int64 {
	if aj == nil {
		return nil
	}
	list, _ := aj[key].([]interface{})
	result := make([]int64, len(list))
	for i, v := range list {
		result[i] = int64(v.(float64))
	}
	return result
}

func (aj AnyJson) GetNumber(key string) float64 {
	if aj == nil {
		return 0
	}
	num, _ := aj[key].(float64)
	return num
}

func (aj AnyJson) GetInt64(key string) int64 {
	return int64(aj.GetNumber(key))
}
