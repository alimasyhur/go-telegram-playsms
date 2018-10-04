package main

import (
	"log"
	"net/http"
	"os"
)

//SetPlaySMSProfil function
func SetPlaySMSProfil(number string) (url string) {
	req, err := http.NewRequest("GET", PlaySMSURL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("app", PlaySMSApp)
	q.Add("op", PlaySMSOpProfil)
	q.Add("u", PlaySMSUserProfil)
	q.Add("h", PlaySMSTokenProfil)
	q.Add("dst", number)
	req.URL.RawQuery = q.Encode()

	url = req.URL.String()
	return
}
