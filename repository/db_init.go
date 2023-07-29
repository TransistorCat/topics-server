package repository

import (
	"github.com/TransistorCat/topics-server/repository/localfile"
	"github.com/TransistorCat/topics-server/repository/mysql"
)

func Init(options Options) error {

	switch options.DBType {
	case 1:
		localfile.Init(&localfile.Default)
	case 2:
		mysql.Init()
	}
	return nil
}
