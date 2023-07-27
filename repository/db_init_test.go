package repository

import "testing"

func TestInit(t *testing.T) {
	options := Options{
		DBType: MySQL,
	}
	Init(options)
	// initTopicIndexMapFormMySQL(DB)
	initPostIndexMapFormMySQL(DB)
}

// func TestinitTopicIndexMapFormMySQL(t *testing.T) {

// }
