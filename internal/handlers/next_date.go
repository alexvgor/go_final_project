package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexvgor/go_final_project/internal/setup"
	"github.com/alexvgor/go_final_project/internal/taskmanager"
)

type NextDateHandler struct {
}

func NewNextDateHandler() *NextDateHandler {
	return &NextDateHandler{}
}

func (h *NextDateHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		now, err := time.Parse(setup.ParseDateFormat, r.FormValue("now"))
		if err != nil {
			http.Error(w, "Неверный формат времени от которого ищется ближайшая дата", http.StatusBadRequest)
			return
		}

		nextDate, err := taskmanager.NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, nextDate)
	}
}
