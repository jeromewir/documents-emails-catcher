package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type mailHandler struct {
	dropboxClient *DropboxClient
}

type MailAttachment struct {
	Filename  string `json:"filename"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	ContentID string `json:"content-id"`
}

// Returns a new instance of mailHandler
func NewMailHandler(dc *DropboxClient) *mailHandler {
	return &mailHandler{
		dropboxClient: dc,
	}
}

func (m *mailHandler) HandleIncomingEmail(c *gin.Context) {
	rawAttachments := c.PostForm("attachment-info")

	var attachments map[string]MailAttachment

	err := json.Unmarshal([]byte(rawAttachments), &attachments)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	erroredFiles := []string{}

	for key, attachment := range attachments {
		fd, _, err := c.Request.FormFile(key)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		err = m.dropboxClient.UploadDatedFile(fd, attachment.Filename)

		if err != nil {
			erroredFiles = append(erroredFiles, attachment.Filename)
			fmt.Println(err)
		}
	}

	if len(erroredFiles) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Some files have not been processed",
			"filesError": erroredFiles,
		})
		return
	}

	c.Status(204)
}
