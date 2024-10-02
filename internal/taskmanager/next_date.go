package taskmanager

import (
	"errors"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/alexvgor/go_final_project/internal/setup"
)

func getWeekDay(now time.Time) int {
	weekDay := int(now.Weekday())
	if weekDay == 0 {
		weekDay = 7
	}
	return weekDay
}

func getClosestWeekDay(repeatWeekDays []string, weekDay int) int64 {
	var closestRepeatWeekDay int64
	for _, day := range repeatWeekDays {
		if repeatWeekDay, err := strconv.ParseInt(day, 10, 64); err == nil {
			if repeatWeekDay < 1 || repeatWeekDay > 7 {
				return 0
			}
			if int(repeatWeekDay) <= weekDay {
				repeatWeekDay += 7
			}
			repeatWeekDay -= int64(weekDay)
			if (closestRepeatWeekDay == 0) || (repeatWeekDay < closestRepeatWeekDay) {
				closestRepeatWeekDay = repeatWeekDay
			}
		}
	}
	return closestRepeatWeekDay
}

func NextDate(now time.Time, dateString string, repeatString string) (string, error) {
	if repeatString == "" {
		return "", errors.New("правило для повтора не указано")
	}

	date, err := time.Parse(setup.ParseDateFormat, dateString)

	var parsedDate string

	if err != nil {
		return "", errors.New("исхоное время переданно в неверном формате")
	}

	switch {
	case strings.HasPrefix(repeatString, "d "):
		days, err := strconv.Atoi(strings.TrimPrefix(repeatString, "d "))
		if err != nil || days < 1 || days > 400 {
			return "", errors.New("недопустимое значение для переноса дней")
		}
		for {
			if date.After(now) {
				parsedDate = date.Format(setup.ParseDateFormat)
				if dateString < parsedDate {
					break
				}
			}
			date = date.AddDate(0, 0, days)
		}
	case repeatString == "y":
		for {
			if date.After(now) {
				parsedDate = date.Format(setup.ParseDateFormat)
				if dateString < parsedDate {
					break
				}
			}
			date = date.AddDate(1, 0, 0)
		}
	case strings.HasPrefix(repeatString, "w "):
		repeatWeekDays := strings.Split(strings.TrimPrefix(repeatString, "w "), ",")

		var closestWeekDay int64

		for {
			if date.After(now) {
				parsedDate = date.Format(setup.ParseDateFormat)
				if dateString < parsedDate {
					break
				}
			}
			closestWeekDay = getClosestWeekDay(repeatWeekDays, getWeekDay(date))
			if closestWeekDay == 0 {
				return "", errors.New("недопустимое значение для переноса дней недели")
			}
			date = date.AddDate(0, 0, int(closestWeekDay))
		}

	case strings.HasPrefix(repeatString, "m "):

		daysMonthData := strings.Split(strings.TrimPrefix(repeatString, "m "), " ")

		var days []int
		for _, dayString := range strings.Split(daysMonthData[0], ",") {
			if day, err := strconv.ParseInt(dayString, 10, 64); err == nil {
				if day < -2 || day > 31 {
					return "", errors.New("недопустимое значение дня для переноса дней месяца")
				}
				days = append(days, int(day))
			}
		}

		var monthes []int
		if len(daysMonthData) < 2 {
			monthes = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
		} else {
			monthesAsStrings := strings.Split(daysMonthData[1], ",")
			for _, monthAsString := range monthesAsStrings {
				if month, err := strconv.ParseInt(monthAsString, 10, 64); err == nil {
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
				currentMonth := int(date.Month())
				if currentMonth > month {
					continue
				} else if currentMonth < month {
					if date.Day() > 1 {
						date = date.AddDate(0, 0, 1-date.Day())
					}
					date = date.AddDate(0, month-int(currentMonth), 0)
				}

				daysInMonthFrom := int(date.Day())
				daysInMonth := time.Date(date.Year(), date.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
				daysForThisMonth := make([]int, len(days))
				copy(daysForThisMonth, days)
				for dayIndex := range daysForThisMonth {
					if daysForThisMonth[dayIndex] == -1 {
						daysForThisMonth[dayIndex] = daysInMonth
					} else if daysForThisMonth[dayIndex] == -2 {
						daysForThisMonth[dayIndex] = daysInMonth - 1
					}
				}
				slices.Sort(daysForThisMonth)

				for _, day := range daysForThisMonth {
					for dayInMonth := daysInMonthFrom; dayInMonth < daysInMonth+1; dayInMonth++ {
						if dayInMonth == day {
							date = date.AddDate(0, 0, day-date.Day())
							if date.After(now) {
								parsedDate = date.Format(setup.ParseDateFormat)
								if dateString < parsedDate {
									return parsedDate, nil
								}
							}
						} else if dayInMonth > day {
							break
						}
					}
				}
			}
			if date.Day() > 1 {
				date = date.AddDate(0, 0, 1-date.Day())
			}
			date = date.AddDate(0, (12-int(date.Month()))+monthes[0], 0)
		}

	default:
		return "", errors.New("неверный формат правила повтора")
	}

	return parsedDate, nil
}
