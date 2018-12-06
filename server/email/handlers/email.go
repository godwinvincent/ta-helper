package handlers

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func (ctx *Context) EmailSendHandler(w http.ResponseWriter, r *http.Request, user *User) {
	if r.Method == "GET" {
		randNum, _ := rand.Int(rand.Reader, big.NewInt(100000000000))
		ctx.UserStore.SetVerifCode(user.UserName, randNum.String())
		code := randNum.String()
		from := mail.NewEmail("TA Helper", "TAHelper@godwinv.com")
		subject := "TA Helper Verification"
		to := mail.NewEmail(user.FirstName+" "+user.LastName, user.Email)
		plainTextContent := "Hello " + user.FirstName + ",<br> Thanks you for registering with TA pal! Verification Code: " + code + "<br>" + "Thanks,<br>The TA Pal Team"
		htmlContent := "Hello " + user.FirstName + ",<br> Thanks you for registering with TA pal! Verification Code: " + code + "<br>" + "Thanks,<br>The TA Pal Team"
		message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
		response, err := client.Send(message)
		if err != nil {
			log.Println(err)
		} else {
			fmt.Println(response.StatusCode)
			fmt.Println(response.Body)
			fmt.Println(response.Headers)
			w.Write([]byte("sent email"))
		}
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}

func (ctx *Context) EmailVerifyHandler(w http.ResponseWriter, r *http.Request, user *User) {
	if r.Method == "GET" {
		received := r.URL.Query().Get("c")
		userCode, err := ctx.UserStore.GetVerifCode(user.UserName)
		if err != nil {
			//do something
		}
		if userCode == received {
			w.Write([]byte("verified"))
			if err := ctx.UserStore.SetUserVerified(user.UserName); err != nil {
				//do something
			}
		} else {
			w.Write([]byte("wrongCode"))
			//increment fail counter
		}
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}
