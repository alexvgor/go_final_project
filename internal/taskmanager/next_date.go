package taskmanager

import (
	"errors"
	"slices"
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

	case strings.HasPrefix(repeat_string, "w "):
		week_day := int(now.Weekday())
		if week_day == 0 {
			week_day = 7
		}
		repeat_week_days := strings.Split(strings.TrimPrefix(repeat_string, "w "), ",")
		var closest_repeat_week_day int64
		for _, day := range repeat_week_days {
			if repeat_week_day, err := strconv.ParseInt(day, 10, 64); err == nil {
				if repeat_week_day < 1 || repeat_week_day > 7 {
					return "", errors.New("недопустимое значение для переноса дней недели")
				}
				if int(repeat_week_day) <= week_day {
					repeat_week_day += 7
				}
				if closest_repeat_week_day == 0 || repeat_week_day < closest_repeat_week_day {
					closest_repeat_week_day = repeat_week_day
				}
			}
		}
		date = now.AddDate(0, 0, int(closest_repeat_week_day)-week_day)

	case strings.HasPrefix(repeat_string, "m "):
		days_month_data := strings.Split(strings.TrimPrefix(repeat_string, "m "), " ")

		var days []int
		for _, day_string := range strings.Split(days_month_data[0], ",") {
			if day, err := strconv.ParseInt(day_string, 10, 64); err == nil {
				if day < -2 || day > 31 {
					return "", errors.New("недопустимое значение дня для переноса дней месяца")
				}
				days = append(days, int(day))
			}
		}

		var monthes []int
		if len(days_month_data) < 2 {
			for i := range 12 {
				monthes = append(monthes, i+1)
			}
		} else {
			monthes_as_strings := strings.Split(days_month_data[1], ",")
			for _, month := range monthes_as_strings {
				if month, err := strconv.ParseInt(month, 10, 64); err == nil {
					if month < 1 || month > 12 {
						return "", errors.New("недопустимое значение месяца для переноса дней месяца")
					}
					monthes = append(monthes, int(month))
				}
			}
		}
		slices.Sort(monthes)

		for {
			for _, month := range monthes {

				current_month := int(date.Month())
				if current_month > month {
					continue
				} else if current_month < month {
					date = date.AddDate(0, month-int(current_month), 0)
					if date.Day() > 1 {
						date = date.AddDate(0, 0, 1-date.Day())
					}
				}

				days_in_month_from := int(date.Day())
				days_in_month := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
				days_for_this_month := days
				for day_index := range days_for_this_month {
					if days_for_this_month[day_index] == -1 {
						days_for_this_month[day_index] = days_in_month
					} else if days_for_this_month[day_index] == -2 {
						days_for_this_month[day_index] = days_in_month - 1
					}
				}
				slices.Sort(days_for_this_month)

				for _, day := range days_for_this_month {
					for day_in_month := days_in_month_from; day_in_month < days_in_month+1; day_in_month++ {
						if day_in_month == day {
							date = date.AddDate(0, 0, day-date.Day())
							if date.After(now) {
								return date.Format(setup.ParseDateFormat), nil
							}
						} else if day_in_month > day {
							break
						}
					}
				}
			}
			date = date.AddDate(0, (12-int(date.Month()))+monthes[0], 0)
			if date.Day() > 1 {
				date = date.AddDate(0, 0, 1-date.Day())
			}
		}

	default:
		return "", errors.New("неверный формат правила повтора")
	}

	return date.Format(setup.ParseDateFormat), nil
}
