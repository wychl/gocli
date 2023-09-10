package cron

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

type Cron struct{}

func New() *Cron {
	return &Cron{}
}

// Name command name
func (c *Cron) Name() string {
	return "cron"
}

// Run 运行command
func (c *Cron) Run(config interface{}) error {
	cfg, ok := config.(*Config)
	if !ok {
		return errors.New("config invalid")
	}
	var parser cron.Parser
	if cfg.WithSecond {
		parser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	} else {
		parser = cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	}
	schedule, err := parser.Parse(string(cfg.Exp))
	if err != nil {
		log.Fatal(err)
	}
	zone := time.Local
	if cfg.Zone != "" {
		var err error
		zone, err = time.LoadLocation(cfg.Zone) //nolint
		if err != nil {
			log.Fatal(err)
		}
	}
	startDate := time.Now()
	if cfg.Start != "" {
		startDate, err = time.Parse(time.DateOnly, cfg.Start)
		if err != nil {
			log.Fatal(err)
		}
	}
	builder := strings.Builder{}
	cur := startDate.In(zone)
	for i := 0; i < 10; i++ {
		cur = schedule.Next(cur).In(zone)
		builder.WriteString(fmt.Sprintf("%s\n", cur))
	}
	fmt.Println(builder.String())

	return nil
}
