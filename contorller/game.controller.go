package controller

import (
	"lecture/WBA-BC-Project-04/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Controller) GetTest(c *gin.Context) {
	c.JSON(http.StatusOK, "")
}

func (p *Controller) GamePostMatchController(c *gin.Context) {

	match := model.CreateMatch{}
	if err := c.ShouldBindJSON(&match); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "매개변수 오류!",
			"error":   err.Error(),
		})
		return
	}
	err := p.md.CreateMatchModel(match)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "매치 생성 오류!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (p *Controller) GameEndMatchController(c *gin.Context) {

	match := model.EndMatch{}
	if err := c.ShouldBindJSON(&match); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "매개변수 오류!",
			"error":   err.Error(),
		})
		return
	}
	err := p.md.EndMatchModel(match)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "매치 종료 오류!",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, "ok")
}
