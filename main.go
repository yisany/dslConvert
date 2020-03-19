package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	c "yisany.top/milu/dslConvert/convert"
)

func main() {
	//gin.SetMode(gin.ReleaseMode)
	//router := routers.InitRouter()
	//endless.ListenAndServe(":3232", router)

	if len(os.Args) < 2 {
		log.Fatalf("Error: convert error, param is missing")
		os.Exit(1)
	}

	var sql string
	if strings.HasPrefix(os.Args[1], "\"") {
		sql = os.Args[1]
	} else {
		sql = strings.Join(os.Args[1:], " ")
	}
	fmt.Printf("Init sql:\n%s\n", sql)

	dsl, _, err := c.Convert(sql)
	if err != nil {
		log.Fatalf("sql convert error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Convert dsl:\n%s\n", dsl)

}
