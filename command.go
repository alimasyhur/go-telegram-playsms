package main

//SetMessageReply funcion
func SetMessageReply(url string, incomingMessage []string) string {
	url = SetupPlaySMS(incomingMessage)
	return PlaySMSSend(url)
}

//CheckCommand function
//Check Availeble Command
func CheckCommand(incomingMessage []string) bool {
	allowedCommand := []string{CommandHelp, CommandGetSMS, CommandBedanomer, CommandLupapassword, CommandSMS, CommandGet}
	command := incomingMessage[0]
	res, _ := InArray(command, allowedCommand)
	return res
}

func getSendMessage(arrStr []string) (message string) {
	if len(arrStr) <= 1 {
		switch arrStr[0] {
		case CommandGetSMS:
			message = "Pastikan perintah Anda benar: `getsms <no_hp>`"
		case CommandBedanomer:
			message = "Pastikan perintah Anda benar: `bedanomer <no_hp>`"
		case CommandGet:
			message = "Pastikan perintah Anda benar: `get <no_hp>`"
		case CommandLupapassword:
			message = "Pastikan perintah Anda benar: `lupapassword <no_hp> <kode lupa password>`"
		case CommandHelp:
			message = "Perintah yang tersedia: \n - Request Kode Lupa Password: getsms <no_hp> \n - Mengirim Pesan Nomor di Email terdaftar berbeda dengan nomor tersebut: berbedanomer <no_hp> \n - Mengirim SMS: sms <no_hp> \n - Mengirim Kode Lupapassword: lupapassword <no_hp> <kode lupa password>"
		}
		return
	}

	if len(arrStr) <= 2 {
		if arrStr[0] == CommandSMS {
			message = "Pastikan perintah Anda benar: `sms <no_hp> <pesan>`"
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
