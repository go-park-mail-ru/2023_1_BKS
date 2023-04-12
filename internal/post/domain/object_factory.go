package domain

import "time"

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

func CreatePrice(price string) Price {
	return Price(price)
}

func CreateClose(close bool) Close {
	return Close(close)
}

func CreateTimeStamp(times time.Time) TimeStamp {
	return TimeStamp(times)
}
