package handlers

import (
	"errors"
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
			RespondErrorUnableToFindNextDate(w, errors.New("неверный формат времени от которого ищется ближайшая дата"))
			return
		}

		nextDate, err := taskmanager.NextDate(now, r.FormValue("date"), r.FormValue("repeat"))
		if err != nil {
			RespondErrorUnableToFindNextDate(w, err)
			return
		}

		fmt.Fprint(w, nextDate)
	}
}
