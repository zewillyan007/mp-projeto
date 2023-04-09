package helper

import (
	"time"
)

type DateHelper struct{}

func NewDateHelper() *DateHelper {
	return &DateHelper{}
}

// func (*DateHelper) translate(layout string) string {
// 	layout = strings.ReplaceAll(layout, "t", "-0700")
// 	layout = strings.ReplaceAll(layout, "u", "000")
// 	layout = strings.ReplaceAll(layout, "Y", "2006")
// 	layout = strings.ReplaceAll(layout, "d", "02")
// 	layout = strings.ReplaceAll(layout, "m", "01")
// 	layout = strings.ReplaceAll(layout, "y", "06")
// 	layout = strings.ReplaceAll(layout, "h", "03")
// 	layout = strings.ReplaceAll(layout, "H", "15")
// 	layout = strings.ReplaceAll(layout, "i", "04")
// 	layout = strings.ReplaceAll(layout, "s", "05")
// 	return layout
// }

func (o *DateHelper) Parse(layout string, date string) (*time.Time, error) {

	//	parsed, err := time.ParseInLocation(o.translate(layout), date, time.Local)
	parsed, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

func (o *DateHelper) Format(layout string, date time.Time) string {

	//	return date.In(time.Local).Format(o.translate(layout))
	return date.In(time.Local).Format(layout)
}

func (o *DateHelper) AddSeconds(seconds int, date time.Time) time.Time {

	return date.Add(time.Second * time.Duration(seconds))
}

func (o *DateHelper) AddMinutes(minutes int, date time.Time) time.Time {

	return date.Add(time.Minute * time.Duration(minutes))
}

func (o *DateHelper) AddHours(hours int, date time.Time) time.Time {

	return date.Add(time.Hour * time.Duration(hours))
}

func (o *DateHelper) AddDays(days int, date time.Time) time.Time {

	return date.Add(time.Hour * 24 * time.Duration(days))
}
