package models

type Task struct {
	Type     string
	Params   map[string]string
	CharCode string
	DateFrom string
	DateTo   string
	Uuid     string
}
