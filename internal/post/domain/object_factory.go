package domain

func CreateImages(image []string) Images {
	return Images(image)
}

func CreateDescription(description string) Desciption {
	return Desciption(description)
}

func CreateTitle(title string) Title {
	return Title(title)
}

func CreateTags(tags []string) Tags {
	return Tags(tags)
}
