package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"testserver/fileController"
	"time"

	"golang.org/x/oauth2"
)

func startGdrive(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("Token")
	if err != nil {
		fmt.Println(err.Error())
	}
	//BMToken := []byte(cookie.Value)
	//fmt.Println(cookie.Value)
	var Token oauth2.Token
	strtoToken(cookie.Value, &Token)
	//err = json.Unmarshal(BMToken, &Token)
	//if err != nil {
	//	fmt.Println("error in Unmarshal")
	//	fmt.Println(err.Error())
	//}

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
	for i := 0; i <= tot; i++ {
		fileController.UploadFile(token, googleOauthConfig, filename, i)

	}

	fmt.Println("Workdone check it")
	fmt.Fprintf(w, "WorkDone Check it on Google Drive")

}

func strtoToken(strToken string, token *oauth2.Token) {
	strToken2 := strToken[1 : len(strToken)-1]
	lis := strings.Split(strToken2, ",")
	token.AccessToken = strings.Split(lis[0], ":")[1]
	token.TokenType = strings.Split(lis[1], ":")[1]

	r1, err := time.Parse("RFC3339", strings.Split(lis[2], ":")[1])
	if err != nil {
		fmt.Println("error processing time")
		fmt.Println(err.Error())
		return
	}
	token.Expiry = r1
}
