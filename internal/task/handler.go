package task

import (
	l "GOtestprogect/internal/logger"
	"encoding/json"
	"net/http"
	"strconv"
)

type Handler struct {
	Logchan chan l.Logger
	DataService
}

func (h *Handler) HandlerAddTaks(w http.ResponseWriter, r *http.Request) {
	task := Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		log := l.NewLog("Invalid data", err.Error())
		h.Logchan <- log
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	_, err = strconv.Atoi(task.Id)
	if err != nil {
		log := l.NewLog("Invalid id task", err.Error())
		h.Logchan <- log
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	err = h.DataService.ServiceAdd(task)
	if err != nil {
		log := l.NewLog("failed add new task", err.Error())
		h.Logchan <- log
		http.Error(w, "already exist", http.StatusInternalServerError)
		return
	}
	log := l.NewLog("Task added", "")
	h.Logchan <- log
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Task added"))
}

func (h *Handler) HandlerTakeTasks(w http.ResponseWriter, r *http.Request) {
	res := h.DataService.ServiceTakeAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log := l.NewLog(somewrong, err.Error())
		h.Logchan <- log
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandlerTaskByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query()
	id := query.Get("id")
	if id == "" {
		log := l.NewLog("Invalid Query Params", "")
		h.Logchan <- log
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	task, err := h.DataService.ServiceById(id)
	if err != nil {
		log := l.NewLog("Invalid ID", err.Error())
		h.Logchan <- log
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		log := l.NewLog(somewrong, err.Error())
		h.Logchan <- log
		http.Error(w, "Server Error", http.StatusInternalServerError)
		return
	}
}
