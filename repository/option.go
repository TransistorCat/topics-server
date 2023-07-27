package repository

type Options struct {
	DBType   DBType
	FilePath string
}

type DBType int8

const (
	LocalData DBType = iota + 1
	MySQL
)

var DefaultOptions = Options{
	DBType:   2,
	FilePath: "./data/",
}
