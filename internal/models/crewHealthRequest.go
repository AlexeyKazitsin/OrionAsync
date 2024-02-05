package models

type CrewHealthRequest struct {
	AccessKey     int64 `json:"access_key"`
	IsCrewHealthy int   `json:"is_crew_healthy"`
}
