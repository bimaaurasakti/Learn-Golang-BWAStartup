package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("failed to get campaigns", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := campaign.FormatCampaigns(campaigns)
	response := helper.APIResponse("get campaigns success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.CampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse("failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignData, err := h.campaignService.GetCampaign(input)
	if err != nil {
		response := helper.APIResponse("failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := campaign.FormatCampaignDetail(campaignData)
	response := helper.APIResponse("get campaign detail success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)
	if err != nil {
		response := helper.APIResponse("failed to create campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := campaign.FormatCampaign(newCampaign)
	response := helper.APIResponse("create campaign success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) EditCampaign(c *gin.Context) {
	var inputID campaign.CampaignDetailInput
	var inputData campaign.CreateCampaignInput

	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.APIResponse("failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		response := helper.APIResponse("failed to update campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	campaign, err := h.campaignService.EditCampaign(inputID, inputData)
	if err != nil {
		errorMessages := gin.H{"errors": err.Error()}

		response := helper.APIResponse("failed to update campaign", http.StatusBadRequest, "error", errorMessages)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("update campaign success", http.StatusOK, "error", campaign)
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadCampaignImage(c *gin.Context) {
	var input campaign.UploadCampaignImageInput

	err := c.ShouldBind(&input)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Get user from Middleware
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	input.User = currentUser

	// Create file name
	splitedFileName := strings.Split(file.Filename, ".")
	fileFormat := splitedFileName[len(splitedFileName) - 1]
	path := fmt.Sprint("images/campaign/", userID, "_", time.Now().Format("010206150405"), ".", fileFormat)

	// Save image to directory
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.campaignService.UploadCampaignImage(input, path)
	if err != nil {
		data := gin.H{
			"is_uploaded": false,
		}
		response := helper.APIResponse("failed to upload campaign image", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{
		"is_uploaded": true,
	}
	response := helper.APIResponse("upload campaign image success", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}