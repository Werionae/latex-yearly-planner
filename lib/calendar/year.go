package calendar

import "time"

type Year struct {
	Quarters [4]Quarter
	year     int
}

func NewYear(year int, wd time.Weekday) Year {
	return Year{
		year: year,
		Quarters: [4]Quarter{
			NewQuarter(year, FirstQuarter, wd),
			NewQuarter(year, SecondQuarter, wd),
			NewQuarter(year, ThirdQuarter, wd),
			NewQuarter(year, FourthQuarter, wd),
		},
	}
}

func (y Year) InWeeks() Weeks {
	weeks := make(Weeks, 0, 53)

	week := y.Quarters[0].Months[0].Weeks[0].backfill()
	weeks = append(weeks, week)

	for {
		week = week.Next()
		if week.TailYear() != y.year {
			break
		}

		weeks = append(weeks, week)
	}

	if week.HeadYear() == y.year {
		weeks = append(weeks, week)
	}

	weeks[0] = weeks[0].SetFirst()
	weeks[len(weeks)-1] = weeks[len(weeks)-1].SetLast()

	return weeks
}

func (y Year) Days() Days {
	days := make(Days, 0, 366)

	for _, quarter := range y.Quarters {
		for _, month := range quarter.Months {
			for _, week := range month.Weeks {
				for _, day := range week.Days {
					if day.IsZero() {
						continue
					}

					days = append(days, day)
				}
			}
		}
	}

	return days
}

func (y Year) Year() int {
	return y.year
}

func (y Year) Months() []Month {
	out := make([]Month, 0, 12)

	for _, quarter := range y.Quarters {
		for _, month := range quarter.Months {
			out = append(out, month)
		}
	}

	return out
}
