package service

import (
	"main/database"
	"main/models"
	"time"

	"github.com/google/uuid"
)

func CreateStatusReport(session *models.Session, request *models.CreateStatusReportRequest) (*models.StatusReport, error) {
	model := &models.StatusReport{
		ID:         uuid.NewString(),
		LocationID: request.LocationID,
		UserID:     session.UserID,
		Busyness:   request.Busyness,
		Loudness:   request.Loudness,
		CreatedAt:  time.Now(),
	}

	database.SaveStatusReport(model)

	return model, nil
}

func GetStatusForLocation(session *models.Session, request *models.StatusRequest) (*models.StatusResponse, error) {

	var reportList []*models.StatusReport
	var err error

	if request.Recency != nil {
		reportList, err = database.GetStatusReportsForLocationWithRecency(request.LocationID, *request.Recency)
	} else {
		reportList, err = database.GetStatusReportsForLocation(request.LocationID)
	}
	if err != nil {
		return nil, err
	}

	busynessTotal := 0.0
	loudnessTtotal := 0.0

	for _, report := range reportList {
		busynessTotal += report.Busyness
		loudnessTtotal += report.Loudness
	}

	resp := &models.StatusResponse{
		LocationID:  request.LocationID,
		AvgBusyness: busynessTotal / float64(len(reportList)),
		AvgLoudness: loudnessTtotal / float64(len(reportList)),
	}

	return resp, nil
}
