package calendar

import (
	"strconv"
	"time"
)

type Quarter struct {
	Months [3]Month
	number int
}

func NewQuarter(year int, qrtr int, wd time.Weekday) Quarter {
	if qrtr < FirstQuarter || qrtr > FourthQuarter {
		return Quarter{} //nolint:exhaustivestruct
	}

	mo := time.Month((qrtr+1)*3 - 2) //nolint:gomnd

	return Quarter{
		number: qrtr + 1,

		Months: [3]Month{
			NewMonth(year, mo, wd),
			NewMonth(year, mo+1, wd),
			NewMonth(year, mo+2, wd), //nolint:gomnd
		},
	}
}

func (r Quarter) Name() string {
	return "Q" + strconv.Itoa(r.number)
}
