package gggitea

import "fmt"

func Mirror() {
	env, err := getEnvironment()
	if err != nil {
		fmt.Println("[!] FATAL")
	}
	fmt.Println("URL: " + env.baseUrl)
	fmt.Println("API Token: " + env.apiToken)
	// fmt.Println("Mirroring ...")
}
