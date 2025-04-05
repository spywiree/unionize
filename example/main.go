package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

func GetInt(v Union) int64 {
	switch v.Type() {
	case UnionUint:
		return int64(v.Uint()) //#nosec G115
	case UnionInt:
		return v.Int()
	case UnionString:
		d, _ := strconv.Atoi(v.GetString())
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
	u.SetInt(math.MinInt32)
	PrintJson(u.GetString())
	PrintJson(u.Int())
	PrintJson(u.Uint())
	PrintJson(u.Type(), u.Type().String())
}
