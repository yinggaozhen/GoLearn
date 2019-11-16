package cron

type Job interface {
	Run()
}