package util

import (
	"github.com/samayamnag/boilerplate/config"
	"time"
)

func FormatMongoDate(t time.Time) string {
	return t.Format(config.Get().MagicalDate)
}
