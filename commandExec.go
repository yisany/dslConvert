package main

import (
	"errors"
	"fmt"
	"os"
	c "yisany.top/milu/dslConvert/convert"
)

func newInit() (string, bool, error) {
	if len(os.Args) < 2 {
		return "", false, errors.New("convert error, param is missing")
	}
	sql := os.Args[1]
	isPretty := false
	if len(os.Args) > 2 && "-p" == os.Args[2] {
		isPretty = true
	}
	fmt.Printf("Init sql:\n%s\n", sql)
	return sql, isPretty, nil
}

func command() {
	sql, isPretty, e := newInit()
	if e != nil {
		fmt.Printf("init error: %s\n", e.Error())
		os.Exit(1)
	}
	var dsl string
	var err error
	if isPretty {
		dsl, _, err = c.ConvertPretty(sql)
	} else {
		dsl, _, err = c.Convert(sql)
	}
	if err != nil {
		fmt.Printf("sql convert error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Convert dsl:\n%s\n", dsl)
}
