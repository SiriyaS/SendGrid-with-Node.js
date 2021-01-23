package ping

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

func (ep *Endpoint) PingGetEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingGetEndpoint")

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) PingGetParamsEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingGetPramsEndpoint")

	//ดึงค่าจาก params ชื่อ name
	name := c.Params.ByName("name")
	log.Infof("Params name : %s", name)

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	//return success
	c.JSON(http.StatusOK, result)
	return
}

func (ep *Endpoint) PingPostParamsAndBodyEndpoint(c *gin.Context) {
	defer c.Request.Body.Close()
	log.Infof("Check Heartbeat : PingPostBodyEndpoint")

	//ดึงค่าจาก params ชื่อ name
	name := c.Params.ByName("name")
	log.Infof("Params name : %s", name)

	//ดึงค่าจาก body
	var request inputHeartbeat //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Infof("Body Name : %s", request.Name)
	log.Infof("Body Age : %d", request.Age)

	//เรียก logic
	result, err := checkHeartbeat()
	if err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	//return success
	c.JSON(http.StatusOK, result)
	return
}
