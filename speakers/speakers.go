package speakers

type Speaker struct {
	Name string `json:"name"`
	Img string `json:"img"`
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
	Github string `json:"github"`
	Linkedin string `json:"linkedin"`
	Website string `json:"website"`
	Shortbio string `json:"shortbio"`
	Title string `json:"title"`
	Company string `json:"company"`
}

type Speakers []Speaker
