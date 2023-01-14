package controller

import (
	"lecture/WBA-BC-Project-04/model"
	"net/http"
)

type Controller struct {
	md *model.Model
}

func NewCTL(rep *model.Model) (*Controller, error) {
	r := &Controller{
		md: rep,
	}

	return r, nil
}

func (p *Controller) GetOK(c *gin.Context, resp interface{}) {
	c.JSON(http.StatusOK, resp)
}
