package flag

import (
	"flag"
	"strconv"
)

func StrVar(ptr *string, name, value, usage string, config map[string]string) {
	flag.StringVar(ptr, name, value, usage)
	setStr(ptr, name, config)
}

func BoolVar(ptr *bool, name string, value bool, usage string, config map[string]string) {
	flag.BoolVar(ptr, name, value, usage)
	setBool(ptr, name, config)
}

func setStr(flag *string, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	*flag = val
}

func setBool(flag *bool, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	b, parseErr := strconv.ParseBool(val)

	*flag = b

	if parseErr != nil {
		panic(parseErr)
	}
}
