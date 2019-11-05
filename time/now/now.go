package now

import (
	"time"
)

type Now struct {
	time.Time
}

func (n Now) beginningOfMinute() Now {
	return Now{
		time.Now().Truncate(time.Minute),
	}
}

func (n Now) BeginningOfHour() Now {
	return Now{
		time.Now().Truncate(time.Hour),
	}
}

func (n Now) BeginningOfDay() Now {
	y, m, d := time.Now().Date()

	return Now{
		time.Date(y, m, d, 0, 0, 0, 0, n.Location()),
	}
}

func (n Now) BeginningOfWeek() Now {
	day := n.BeginningOfDay()

	for day.Weekday() != time.Monday {
		day = Now{
			day.Add(time.Hour * -24),
		}
	}

	return day
}

func (n Now) BeginningOfMonth() Now {
	y, m, _ := time.Now().Date()

	return Now{
		time.Date(y, m, 1, 0, 0, 0, 0, n.Location()),
	}
}

func (n Now) String() string {
	return n.Format("2006-01-02 15:04:05")
}
