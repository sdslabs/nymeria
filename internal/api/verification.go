package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/sdslabs/nymeria/internal/smtp"
	"github.com/sdslabs/nymeria/internal/utils"
)

func HandleGetVerificationFlow(c *gin.Context) {
	csrfToken, err := utils.GenerateCSRFToken(c.GetHeader("X-User-ID")) // TODO: Get user ID from session
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to generate CSRF token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     "success",
		"message":    "Email verification flow initiated",
		"csrf_token": csrfToken,
	})
}

func HandlePostVerificationCodeFlow(c *gin.Context) {
	var req VerifyEmailRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	if req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email is required",
		})
		return
	}

	otp, err := smtp.SendOTPHandler(req.Email, req.IsIITRCheck)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email verification code sent",
		"data": gin.H{
			"otp": otp,
		},
	})
}

func HandlePostVerifyEmailFlow(c *gin.Context) {
	var req VerifyEmailCodeRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request body",
		})
		return
	}

	if req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email is required",
		})
		return
	}

	if req.Code == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Code is required",
		})
		return
	}

	err = smtp.VerifyOTPHandler(req.Email, req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Email verified",
	})
}
