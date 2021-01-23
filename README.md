# SendGrid-with-Go

**Sending email via Twilio SendGrid with Go**

## SendGrid
※ Go to app.sendgrid.com
#####
※ Create an account and sign in to your account (You may need to **Create a sender identity** please do so)
#####
※ Go to **Settings > API Keys > Create API Key** and create a new `API Key` and don't forget to copy it

## Implementation
※ Open your project directory
#### Create an environment variable
※ Because `API KEY` is confidential so we need to create `API_KEY` as an environment variable, run the following in your terminal
```sh
$ echo "export SENDGRID_API_KEY=<YOUR_API_KEY>" > sendgrid.env
$ echo "sendgrid.env" >> .gitignore
$ source ./sendgrid.env
```
#### Installation
※ Install the **sendgrid-go** package
```sh
go get github.com/sendgrid/sendgrid-go 
```
#### Time to send your mail
※ Create a file and define detail about your mail
#####
**sender's name, sender's mail, receiver's name receiver's mail, mail's subject, and mail's content(plainText and html)**
```sh
import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {
	from := mail.NewEmail("Example User", "test@example.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "test@example.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
```

※ run this file and check your receiver's mail inbox or go to **Activity** tab to see information about the mail you sent

> You can also watch the steps from this video https://youtu.be/bZwJPbp_JMY
>
> or follow instructions at **Setup Guide > Integrate using our Web API or SMTP relay > Web API > Go**
