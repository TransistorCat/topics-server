package repository

func Init(options Options) error {

	switch options.DBType {
	case 1:
		InitLocalFile(&DefaultLocalFile)
	case 2:
		initMySQL()
	}
	return nil
}
