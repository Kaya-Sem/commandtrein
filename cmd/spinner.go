package cmd

import (
	"time"

	"github.com/briandowns/spinner"
)

type Spinner struct {
	spinner *spinner.Spinner
	prefix  string
	suffix  string
	sleep   time.Duration
}

func NewSpinner(prefix, suffix string, sleep time.Duration) *Spinner {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Prefix = prefix
	s.Suffix = suffix
	return &Spinner{
		spinner: s,
		prefix:  prefix,
		suffix:  suffix,
		sleep:   sleep,
	}
}

func (s *Spinner) Start() {
	s.spinner.Start()
	time.Sleep(s.sleep)
}

func (s *Spinner) Stop() {
	s.spinner.Stop()
}
