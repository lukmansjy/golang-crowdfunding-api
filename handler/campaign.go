package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-crowdfunding-api/campaign"
	"golang-crowdfunding-api/helper"
	"golang-crowdfunding-api/user"
	"net/http"
	"strconv"
	"time"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (handler campaignHandler) GetCampaigns(c *gin.Context) {
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

func (handler campaignHandler) GetCampaign(c *gin.Context) {
	campaignID, err := strconv.Atoi(c.Param("campaign_id"))
	if err != nil {
		response := helper.APIResponse("Error parse campaign_id to int", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaignDetail, err := handler.service.GetCampaignsById(campaignID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := campaign.FormatCampaignDetail(campaignDetail)
	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

func (handler *campaignHandler) CreateCampaign(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.InputCampaign
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	input.User = currentUser

	createCampaign, err := handler.service.CreateCampaign(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to create campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to create campaign", http.StatusOK, "success", campaign.FormatCampaign(createCampaign))
	c.JSON(http.StatusOK, response)
}

func (handler *campaignHandler) UpdateCampaign(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	campaignID, err := strconv.Atoi(c.Param("campaign_id"))
	if err != nil {
		response := helper.APIResponse("Error parse campaign_id to int", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.InputCampaign
	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	input.User = currentUser

	updateCampaign, err := handler.service.UpdateCampaign(campaignID, input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Failed to update campaign", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Success to update campaign", http.StatusOK, "success", campaign.FormatCampaign(updateCampaign))
	c.JSON(http.StatusOK, response)
}

func (handler *campaignHandler) UploadImage(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	var input campaign.InputImage
	err := c.ShouldBind(&input)
	if err != nil {
		//errorMsg := helper.FormatValidationError(err)
		errorMessage := gin.H{"is_uploaded": false, "errors": err.Error()}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	path := fmt.Sprintf("images/campagin-%d-%s", time.Now().UnixMilli(), file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = handler.service.InsertImage(input, path, currentUser)

	if err != nil {
		data := gin.H{"is_uploaded": false, "errors": err.Error()}
		response := helper.APIResponse("Failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Campaign image success upload", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}
