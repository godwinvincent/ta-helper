package handlers

import (
	"crypto/rand"
	"encoding/base64"
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
		randCode := randNum.Int64()
		code := base64.URLEncoding.EncodeToString([]byte(string(randCode)))
		from := mail.NewEmail("TA Helper", "TAHelper@godwinv.com")
		subject := "TA Helper Verification"
		to := mail.NewEmail(user.FirstName+" "+user.LastName, user.Email)
		plainTextContent := "Hello " + user.FirstName + ",<br> Thanks you for registering with TA pal! Please click the following link to verify your email address: " + "http://localhost:8080/v1/verifyEmail?c=" + code + "<br>" + "Thanks,<br>The TA Pal Team"
		htmlContent := "Hello " + user.FirstName + ",<br> Thanks you for registering with TA pal! Please click the following link to verify your email address: " + "http://localhost:8080/v1/verifyEmail?c=" + code + "<br>" + "Thanks,<br>The TA Pal Team"
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
		// userCode := user.verificationCode
		userCode := string(user.VerificationCode)
		received := r.URL.Query().Get("c")
		code, _ := base64.URLEncoding.DecodeString(received)
		if userCode == string(code) {
			w.Write([]byte("verified"))
			//write verified to DB
		} else {
			w.Write([]byte("wrongCode"))
			//increment fail counter
		}
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}
