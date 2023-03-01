package cfapi

import (
	"encoding/json"
	"fmt"
	"github.com/qu-bit1/project_new/pkg/models"
	"go.uber.org/zap"
	"io"
	"net/http"
)

type CodeforcesAPI interface {
	RecentActions(maxCount int) ([]models.RecentAction, error)
}

type CodeforcesClient struct {
	client http.Client
}

func (cfClient *CodeforcesClient) RecentActions(maxCount int) ([]models.RecentAction, error) {
	resp, err := cfClient.client.Get("https://codeforces.com/api/recentActions?maxCount=30")
	if err != nil {
		zap.S().Errorf("error occurred while calling cfapi ; %v", err) //zap.S logs in a nice and pretty manner
		return nil, err
	}
	defer resp.Body.Close()
	data, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("error occurred while reading body")
		return nil, err1
	}
	//zap.S().Info(string(data))

	// TODO- unmarshall the data into models.recentActions ( will make a wrapper struct for that)
	wrapper := struct {
		Status string
		Result []models.RecentAction
	}{}

	json.Unmarshal(data, &wrapper)
	return wrapper.Result, err
}

// NewCodeforcesClient return an empty interface on which recent actions method can be applied
func NewCodeforcesClient() CodeforcesAPI {
	obj := new(CodeforcesClient)
	return obj
}
