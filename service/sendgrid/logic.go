package sendgrid

import (
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func sendMail(body inputMail) (result outputModel, err error) {

	from := mail.NewEmail(body.SenderName, body.SenderMail)
	subject := body.Subject
	to := mail.NewEmail(body.ReceiverName, body.ReceiverMail)
	plainTextContent := body.Content
	htmlContent := "<strong>" + body.Content + "</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return
	} else {
		log.Infof("StatusCode : %d", response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
		if response.StatusCode != 202 {
			return
		}
		result.Message = "Send mail successfully"
		result.DateTime = time.Now()
	}

	return
}
