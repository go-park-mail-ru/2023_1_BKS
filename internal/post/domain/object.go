package domain

type Images []string

func (l Images) String() []string {
	var newString []string
	for i := 0; i < len(l); i++ {
		newString = append(newString, l[i])
	}
	return newString
}

type Desciption string

func (l Desciption) String() string {
	return string(l)
}

type Title string

func (l Title) String() string {
	return string(l)
}

type Tags []string

func (l Tags) String() []string {
	var newString []string
	for i := 0; i < len(l); i++ {
		newString = append(newString, l[i])
	}
	return newString
}
