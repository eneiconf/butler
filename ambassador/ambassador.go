package ambassador

type Ambassador struct {
	University   string `json:"university"`
	Elements  []Element `json:"elements"`
}

type Element struct {
	Name string `json:"name"`
	Img  string `json:"img"`
	Facebook string `json:"facebook"`
}

type Ambassadors []Ambassador
