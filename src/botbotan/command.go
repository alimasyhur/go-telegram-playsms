package main

//SetMessageReply funcion
func SetMessageReply(url string, incomingMessage []string) (message string) {
	url = SetupPlaySMS(incomingMessage)
	message = PlaySMSSend(url)
	return
}

//CheckCommand function
//Check Availeble Command
func CheckCommand(incomingMessage []string) bool {
	allowedCommand := []string{CommandGetSMS, CommandBedanomer, CommandLupapassword, CommandSMS, CommandGet}
	command := incomingMessage[0]
	res, _ := InArray(command, allowedCommand)
	return res
}

func getSendMessage(arrStr []string) (message string) {
	if arrStr[0] == CommandGetSMS {
		if len(arrStr) <= 1 {
			message = "Pastikan perintah Anda benar: `getsms <no_hp>`"
			return
		}
	} else if arrStr[0] == CommandBedanomer {
		if len(arrStr) <= 1 {
			message = "Pastikan perintah Anda benar: `bedanomer <no_hp>`"
			return
		}
	} else if arrStr[0] == CommandSMS {
		if len(arrStr) <= 1 {
			message = "Pastikan perintah Anda benar: `sms <no_hp>`"
			return
		}
	} else if arrStr[0] == CommandGet {
		if len(arrStr) <= 1 {
			message = "Pastikan perintah Anda benar: `get <no_hp>`"
			return
		}
	} else if arrStr[0] == CommandLupapassword {
		if len(arrStr) <= 2 {
			message = "Pastikan perintah Anda benar: `lupapassword <no_hp> <kode lupa password>`"
			return
		}
	}

	if CheckCommand(arrStr) {
		if arrStr[0] == CommandGetSMS {
			urlProfil := SetPlaySMSProfil(arrStr[1])
			message = PlaySMSGetRequest(urlProfil)
		} else {
			url := SetupPlaySMS(arrStr)
			message = SetMessageReply(url, arrStr)
		}
	}
	return
}
