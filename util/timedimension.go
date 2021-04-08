package util

import "fmt"

var (
	TimeDimensionSecond *TimeDimension = newTimeDimension("SECOND(%v)")
	TimeDimensionMinute *TimeDimension = newTimeDimension("MINUTE(%v)")
	TimeDimensionHour   *TimeDimension = newTimeDimension("HOUR(%v)")
)

type TimeDimension struct {
	groupByTemplate string
}

func newTimeDimension(groupByTemplate string) *TimeDimension {
	return &TimeDimension{groupByTemplate}
}

func GetTimeDimension(dimension string) (*TimeDimension, error) {
	if dimension == "second" {
		return TimeDimensionSecond, nil
	} else if dimension == "minute" {
		return TimeDimensionMinute, nil
	} else if dimension == "hour" {
		return TimeDimensionHour, nil
	}
	return nil, ErrInvalid
}

func (td *TimeDimension) GroupBy(columnName string) string {
	return fmt.Sprintf(td.groupByTemplate, columnName)
}
