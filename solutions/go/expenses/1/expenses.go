package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filteredRecord []Record
	for _, record := range in {
		if predicate(record) {
			filteredRecord = append(filteredRecord, record)
		}
	}
	return filteredRecord
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		if r.Day >= p.From && r.Day <= p.To {
			return true
		}
		return false
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		if r.Category == c {
			return true
		}
		return false
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	filteredRecord := Filter(in, ByDaysPeriod(p))
	totalAmount := 0.0
	for _, record := range filteredRecord {
		totalAmount += record.Amount
	}
	return totalAmount
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	filteredRecord := Filter(in, ByCategory(c))
	if len(filteredRecord) == 0 {
		return 0, errors.New("unknow category entertainment")
	}
	filteredRecord = Filter(filteredRecord, ByDaysPeriod(p))
	if len(filteredRecord) == 0 {
		return 0, nil
	}
	return TotalByPeriod(filteredRecord, p), nil
}
