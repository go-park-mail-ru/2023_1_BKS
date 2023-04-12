package config

func CreateConfig() Config {
	return Config{
		Http:  CreateHttpConfig(),
		Grcp:  CreateGrcpConfig(),
		Db:    CreateDataBaseConfig(),
		Valid: CreateValidConfig(),
	}
}

func CreateHttpConfig() HttpConfig {
	return HttpConfig{8080}
}

func CreateGrcpConfig() GrcpConfig {
	return GrcpConfig{8081}
}

func CreateDataBaseConfig() DataBaseConfig {
	return DataBaseConfig{
		User:         "appuniq",
		DataBaseName: "user",
		Password:     "uniq123",
		Host:         "127.0.0.1",
		Port:         5432,
		Sslmode:      "disable",
	}
}

func CreateValidConfig() ValidConfig {
	return ValidConfig{
		TitleValidate:       CreateTitleConfig(),
		DescriptionValidate: CreateDescriptionConfig(),
		ImagesValidate:      CreateImagesConfig(),
	}
}
func CreateTitleConfig() TitleConfig {
	return TitleConfig{5, 40}
}

func CreateDescriptionConfig() DescriptionConfig {
	return DescriptionConfig{20, 150}
}

func CreateImagesConfig() ImagesConfig {
	return ImagesConfig{1, 5, 20971520}
}

func CreateTagsConfig() TagsConfig {
	return TagsConfig{1, 5}
}
