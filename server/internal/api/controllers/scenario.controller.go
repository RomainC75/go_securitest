package controllers

import (
	"encoding/json"
	"net/http"
	dto_req "server/internal/api/dtos/requests"
	"server/internal/queue"
)

type AnalyseCtrl struct {
	Queue *queue.SQueue
}

func NewAnalyseCtrl() *AnalyseCtrl {
	return &AnalyseCtrl{
		Queue: queue.GetQueue(),
	}
}

func (c *AnalyseCtrl) HandleAnalyse(w http.ResponseWriter, r *http.Request) {
	var u dto_req.UserCredsDto

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
