package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Localimizde bir yere istek atarken get kullanamıyor muyuz? listen gibi bir şeyler mi
	// kullanmalıyız?
	// go.mod'daki github'ı sormayı unutma!
	resp, err := http.Get("http://host.docker.internal:5000")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}
