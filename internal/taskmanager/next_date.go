package taskmanager

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/alexvgor/go_final_project/internal/setup"
)

func NextDate(now time.Time, date_strint string, repeat_string string) (string, error) {
	if repeat_string == "" {
		return "", errors.New("правило для повтора не указано")
	}

	date, err := time.Parse(setup.ParseDateFormat, date_strint)
	if err != nil {
		return "", errors.New("исхоное время переданно в неверном формате")
	}

	switch {
	case strings.HasPrefix(repeat_string, "d "):
		days, err := strconv.Atoi(strings.TrimPrefix(repeat_string, "d "))
		if err != nil || days < 1 || days > 400 {
			return "", errors.New("недопустимое значение для переноса дней")
		}
		for {
			date = date.AddDate(0, 0, days)
			if date.After(now) {
				break
			}
		}
	case repeat_string == "y":
		for {
			date = date.AddDate(1, 0, 0)
			if date.After(now) {
				break
			}
		}

	default:
		return "", errors.New("неверный формат правила повтора")
	}

	return date.Format(setup.ParseDateFormat), nil
}
