package main

import (
	"fmt"
	"loan_apps/config/setting"
	"loan_apps/internal/server/rest"
)

func main() {
	// TODO: add logic
	fmt.Println(setting.APIPathPrefix)
	rest.Main()
}
