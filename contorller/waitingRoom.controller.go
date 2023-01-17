package controller

import (
	// "lecture/WBA-BC-Project-04/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Controller) GetWaitingRoomTest(c *gin.Context) {
	c.JSON(http.StatusOK, "waitingroom")
}
