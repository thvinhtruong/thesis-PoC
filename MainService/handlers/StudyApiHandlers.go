package handlers

import (
	"encoding/json"
	"net/http"
	GrpcStudyService "server/MainService/GrpcClients/StudyService"
	"server/MainService/config"
	reverseproxy "server/MainService/reverse_proxy"
	_struct "server/MainService/struct"
	"strconv"

	"github.com/gorilla/mux"
)

type StudyHandler struct {
	Repo GrpcStudyService.StudyServiceRepository
	C    config.Config
}

func NewStudyHandler(c config.Config, repo GrpcStudyService.StudyServiceRepository) StudyHandler {
	return StudyHandler{
		Repo: repo,
		C:    c,
	}
}

func (s *StudyHandler) CreateUserRecord(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(r.FormValue("UserId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	moduleId, err := strconv.Atoi(r.FormValue("ModuleId"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	weight, err := strconv.Atoi(r.FormValue("Weight"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	score, err := strconv.Atoi(r.FormValue("Score"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := GrpcStudyService.CreateUserRecordRequest{
		UserId:   int32(userId),
		ModuleId: int32(moduleId),
		Weight:   int32(weight),
		Score:    int32(score),
	}

	response := s.Repo.CreateUserRecord(&request)

	out, _ := json.Marshal(response)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (s *StudyHandler) GetUserRecord(w http.ResponseWriter, r *http.Request) {
	isEnableCache := true
	if reverseproxy.HttpCacheWriter(w, r, nil, isEnableCache) {
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	request := GrpcStudyService.GetUserRecordRequest{
		UserId: int32(userId),
	}

	response := s.Repo.GetUserRecord(&request)

	if response == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	message := _struct.ApiMessage{
		ErrorCode: 1,
		Message:   "success",
		Data:      _struct.GetUserRecordResponseData(userId, response),
	}

	out, _ := json.Marshal(message)

	if !reverseproxy.HttpCacheWriter(w, r, out, isEnableCache) {
		w.Write(out)
	}
}
