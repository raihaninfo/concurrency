package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	timeoutcontext, cancel := context.WithTimeout(context.Background(), time.Microsecond*100)
	defer cancel()
	req, err := http.NewRequestWithContext(timeoutcontext, http.MethodGet, "https://place-hold.it/300", nil)
	if err != nil {
		log.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	imageData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("size %d\n", len(imageData))
}
