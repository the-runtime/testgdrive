package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testserver/fileController"

	"golang.org/x/oauth2"
)

func startGdrive(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		fmt.Println(err.Error())
	}
	BMToken := []byte(cookie.Value)
	var Token oauth2.Token
	err = json.Unmarshal([]byte(BMToken), &Token)
	if err != nil {
		fmt.Println(err.Error())
	}

	token := &Token
	downUrl := r.FormValue("url")

	if downUrl == "" {
		fmt.Println("Empty url")
		return
	}
	filename, tot := fileController.StratDu(downUrl)
	if filename == "" {
		return
	}
	for i := 0; i < tot; i++ {
		fileController.UploadFile(token, googleOauthConfig, filename, i)

	}

	fmt.Println("Workdone check it")
	fmt.Fprintf(w, "WorkDone Check it on Google Drive")

}
