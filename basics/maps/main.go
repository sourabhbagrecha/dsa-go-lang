package main

import "fmt"

func getEmailByDetails(detail map[string]string) string {
	return detail["username"] + "@" + detail["domain"]
}

func main() {
	detail := map[string]string{"username": "sourabhbagrecha", "domain": "gmail.com"}
	fmt.Println(getEmailByDetails(detail))
}
