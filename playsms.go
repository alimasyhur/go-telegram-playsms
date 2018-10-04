package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

//SetupPlaySMS function return url
func SetupPlaySMS(incomingMessage []string) (url string) {
	var message string
	req, err := http.NewRequest("GET", PlaySMSURL, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	message = getMessage(incomingMessage)

	q := req.URL.Query()
	q.Add("app", PlaySMSApp)
	q.Add("op", PlaySMSOp)
	q.Add("u", PlaySMSUser)
	q.Add("h", PlaySMSToken)
	q.Add("to", incomingMessage[1])
	q.Add("msg", message)
	req.URL.RawQuery = q.Encode()

	url = req.URL.String()
	return
}

func getMessage(incomingMessage []string) (message string) {
	if incomingMessage[0] == CommandLupapassword {
		message = "Kode anda : " + incomingMessage[2] + " . Gunakan kode tersebut pada kolom kode verifikasi di laman Profil UNS"
	} else if incomingMessage[0] == CommandSMS {
		message = strings.Join(incomingMessage[2:], " ")
	} else if incomingMessage[0] == CommandBedanomer {
		message = "mohon maaf. permintaan token hrs dikirim dr no yg sama dg yg terdaftar di siakad/simpeg. Update no. di siakad-old.uns.ac.id/registrasi atau ke operator simpeg"
	}
	return
}

//PlaySMSSend function
func PlaySMSSend(url string) (message string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}
	var jsonData ResponseError
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}

	if jsonData.Status == "ERR" {
		message = "Status: " + jsonData.Status + " - " + jsonData.ErrorString
	} else {
		var jsonData ResponseSuccess
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			fmt.Println(err.Error())
			message = err.Error()
			return
		}
		fmt.Println(jsonData.Data[0].SmsLog)
		message = "Status: " + jsonData.Data[0].Status
	}
	return
}

//PlaySMSGetRequest function
func PlaySMSGetRequest(url string) (message string) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}
	var jsonData ResponseError
	err = json.Unmarshal(body, &jsonData)
	if err != nil {
		fmt.Println(err.Error())
		message = err.Error()
		return
	}

	if jsonData.Status == StatusError {
		message = "Status: " + jsonData.Status + " - " + jsonData.ErrorString
	} else {
		var jsonData ResponseSuccessProfil
		err = json.Unmarshal(body, &jsonData)
		if err != nil {
			fmt.Println(err.Error())
			message = err.Error()
			return
		}
		arrMessage := strings.Split(jsonData.Data[0].Message, " ")

		if len(arrMessage) < 4 {
			message = "Nomor " + jsonData.Data[0].Destination + " belum pernah request kode."
		} else {
			kode := arrMessage[3]
			if strings.ToLower(kode) == strings.ToLower("uns") {
				message = "sms " + jsonData.Data[0].Destination + " " + jsonData.Data[0].Message
			} else {
				message = "lupapassword " + jsonData.Data[0].Destination + " " + kode
			}
		}
		fmt.Println(jsonData.Data[0].Message)
	}
	return
}
