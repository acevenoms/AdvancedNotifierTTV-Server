package model

import (
	"fmt"
	"regexp"
	"time"
)

type IFilter interface {
	Pass(INotification) bool
	String() string
}

type FilterSet struct {
	Name string
	Any  []IFilter
	All  []IFilter
}

func (fs FilterSet) Pass(notif INotification) bool {
	pass := true
	any := len(fs.Any) < 1
	for _, f := range fs.Any {
		any = any || f.Pass(notif)
	}

	all := true
	for _, f := range fs.All {
		all = all && f.Pass(notif)
	}

	pass = (any && all)
	return pass
}

func (fs FilterSet) String() string {
	return fmt.Sprintf("Filter Set %s: Any of %d filters, All of %d filters", fs.Name, len(fs.Any), len(fs.All))
}

type StringFilter struct {
	Key    string
	Regex  string
	Invert bool
}

func (sf StringFilter) Pass(notif INotification) bool {
	match, err := regexp.MatchString(sf.Regex, notif.Value(sf.Key))
	return (err != nil) && (match != sf.Invert)
}

func (sf StringFilter) String() string {
	invertString := ""
	if sf.Invert {
		invertString = "NOT "
	}
	return fmt.Sprintf("String filter: When %s does %smatch %s", sf.Key, invertString, sf.Regex)
}

type TimeFilter struct {
	Start time.Time
	End   time.Time
}

func (tf TimeFilter) Pass(notif INotification) bool {
	now := time.Now()
	return now.After(tf.Start) && now.Before(tf.End)
}

func (tf TimeFilter) String() string {
	return fmt.Sprintf("Time Filter: Between %s and %s", tf.Start, tf.End)
}
