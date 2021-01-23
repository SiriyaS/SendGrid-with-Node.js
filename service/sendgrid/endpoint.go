package sendgrid

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	log "github.com/sirupsen/logrus"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

//รับ INPUT แปลงค่า

func (ep *Endpoint) SendGridPost(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Sendgrid : SendgridPostBodyEndpoint")

	//ดึงค่าจาก body
	var request inputMail //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("Body Sender : %s, %s", request.SenderName, request.SenderMail)
	log.Infof("Body Receiver : %s, %s", request.ReceiverName, request.ReceiverMail)
	log.Infof("Body Subject : %s", request.Subject)
	log.Infof("Body content : %s", request.Content)

	//เรียก logic
	result, err := sendMail(request)
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//return success
	c.JSON(http.StatusOK, result)
	return
}
