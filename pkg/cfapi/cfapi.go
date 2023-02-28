package cfapi

import (
	"github.com/qu-bit1/project_new/pkg/models"
)

type CodeforcesAPI interface {
	RecentActions(maxCount int) ([]models.RecentAction, error)
}
