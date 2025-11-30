package http

import (
	"encoding/json"
	"geoserv/internal/domain/models"
	"geoserv/internal/usecase"
	"net/http"
	"strconv"
)

type DriverHandler struct {
	dpu *usecase.DriverPositionUsecase
}

func NewDriverHandler(dpu *usecase.DriverPositionUsecase) *DriverHandler {
	return &DriverHandler{dpu}
}

func (dh *DriverHandler) AddDriver(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var position models.DriverPosition
	if err := json.NewDecoder(r.Body).Decode(&position); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := dh.dpu.AddDriverPosition(ctx, position); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (dh *DriverHandler) GetNerdyDrivers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query := r.URL.Query()
	lat, _ := strconv.ParseFloat(query.Get("lat"), 64)
	lon, _ := strconv.ParseFloat(query.Get("lon"), 64)
	radius, _ := strconv.ParseFloat(query.Get("radius"), 64)
	res, err := dh.dpu.GetDriversByRadius(ctx, lon, lat, radius)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
