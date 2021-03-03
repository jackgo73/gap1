package metric

import "time"

type ValueType int

const (
	_ ValueType = iota
	Counter
	Gauge
	Untyped
	Summary
	Histogram
)

type Tag struct {
	Key   string
	Value string
}

type Field struct {
	Key   string
	Value interface{}
}

type Metric interface {
	Name() string
	Tags() map[string]string
	TagList() []*Tag
	Fields() map[string]interface{}
	FieldList() []*Field
	Time() time.Time
	Type() ValueType
	HashID() uint64
	Copy() Metric
	Accept()
	Reject()
	Drop()

	SetName(name string)
	SetTime(t time.Time)
	AddPrefix(prefix string)
	AddSuffix(suffix string)
	GetTag(key string) (string, bool)
	HasTag(key string) bool
	AddTag(key, value string)
	RemoveTag(key string)
	GetField(key string) (interface{}, bool)
	HasField(key string) bool
	AddField(key string, value interface{})
	RemoveField(key string)
	SetAggregate(bool)
	IsAggregate() bool
}
