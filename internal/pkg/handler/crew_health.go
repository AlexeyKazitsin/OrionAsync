package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lab8/internal/models"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) issueCrewHealth(c *gin.Context) {
	var input models.Request
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("handler.issueCrewHealth:", input)

	c.Status(http.StatusOK)

	go func() {
		time.Sleep(10 * time.Second)
		sendClinicalTrialRequest(input)
	}()
}

func sendClinicalTrialRequest(request models.Request) {

	var crewHealth = 0
	crewHealth = rand.Intn(2)

	answer := models.CrewHealthRequest{
		AccessKey:     123,
		IsCrewHealthy: crewHealth,
	}

	client := &http.Client{}

	jsonAnswer, _ := json.Marshal(answer)
	bodyReader := bytes.NewReader(jsonAnswer)

	requestURL := fmt.Sprintf("http://127.0.0.1:8000/api/flights/%d/update_crew_health/", request.FlightId)

	req, _ := http.NewRequest(http.MethodPut, requestURL, bodyReader)

	req.Header.Set("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		return
	}

	defer response.Body.Close()

	fmt.Println("PUT Request Status:", response.Status)
}
