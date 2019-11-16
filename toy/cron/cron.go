package cron

type Cron struct {
	entries []*Entry
}

type Entry struct {
	schedule *schedule
	job Job
}

func New() *Cron {
	return &Cron{
	}
}

func (c Cron) AddJob(schedule string, job Job) error {
	entry := &Entry{
		schedule: parse(schedule),
		job: job,
	}

	c.entries = append(c.entries, entry)
	return nil
}

func (c Cron) Start() {
}