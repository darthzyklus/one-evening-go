package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {
	lifetime, err := NewDateRange(
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(lifetime.Hours())

	}

	travelInTime, err := NewDateRange(
		time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(travelInTime.Hours())
	}
}

func NewDateRange(start time.Time, end time.Time) (DateRange, error) {
	if start.IsZero() || end.IsZero() {
		return DateRange{}, errors.New("One of the dates is empty")
	}

	if end.Before(start) {
		return DateRange{}, errors.New("End date is before than start date")
	}

	var dateRange = DateRange{
		Start: start,
		End:   end,
	}

	return dateRange, nil
}
