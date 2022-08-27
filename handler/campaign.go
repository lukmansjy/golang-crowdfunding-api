package handler

import (
	"github.com/gin-gonic/gin"
	"golang-crowdfunding-api/campaign"
	"golang-crowdfunding-api/helper"
	"net/http"
	"strconv"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (handler campaignHandler) GetCampaign(c *gin.Context) {
	userIdQuery := c.Query("user_id")
	userId := 0
	if userIdQuery != "" {
		userIdCov, err := strconv.Atoi(userIdQuery)
		if err != nil {
			response := helper.APIResponse("Error parse user_id to int", http.StatusBadRequest, "error", err.Error())
			c.JSON(http.StatusBadRequest, response)
			return
		}
		userId = userIdCov
	}

	campaigns, err := handler.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}
