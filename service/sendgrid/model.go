package sendgrid

import "time"

type outputModel struct {
	Message  string    `json:"message"`
	DateTime time.Time `json:"date_time"`
}

type inputMail struct {
	SenderName   string `json:"sender_name"`
	SenderMail   string `json:"sender_mail"`
	ReceiverName string `json:"receiver_name"`
	ReceiverMail string `json:"receiver_mail"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
}
