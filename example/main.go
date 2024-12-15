package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func GetInt(v Union) int64 {
	switch v.Type() {
	case UnionType_Uint64:
		return int64(v.Uint64())
	case UnionType_Int64:
		return v.Int64()
	case UnionType_String:
		d, _ := strconv.Atoi(v.String())
		return int64(d)
	default:
		panic("unreachable")
	}
}

func PrintJson(v ...any) {
	var data []byte
	var err error
	switch len(v) {
	case 0:
		return
	case 1:
		data, err = json.MarshalIndent(v[0], "", "    ")
	default:
		data, err = json.MarshalIndent(v, "", "    ")
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func main() {
	var u Union
	u.Int64Put(math.MinInt32)
	PrintJson(u.String())
	PrintJson(u.Int64())
	PrintJson(u.Uint64())
	PrintJson(u.Type(), u.Type().String())
}
