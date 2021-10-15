package storage

func GetModels() []interface{} {
	return []interface{}{
		Question{},
		Option{},
		Answer{},
	}
}
