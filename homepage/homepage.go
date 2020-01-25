package homepage

type HomePage struct {
	Banner Banner `json:"banner"`
	Timer Timer `json:"timer"`
	FeaturedSpeakers FeaturedSpeakers `json:"featuredSpeakers"`
	About About `json:"about"`
	TicketInfo []TicketInfo `json:"ticketInfo"`
}

type Banner struct {
	Background string `json:"background"`
	Date string `json:"date"`
	Location string `json:"location"`
}

type Timer struct {
	Starting string `json:"starting"`
}

type FeaturedSpeakers struct {
	Title string `json:"title"`
	Featured []Featured `json:"featured"`
}

type Featured struct {
	Src string `json:"src"`
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type About struct {
	Title string `json:"title"`
	Body []string `json:"body"`
	TeaserURL string `json:"teaserURL"`
}

type TicketInfo struct {
	Title string `json:"title"`
	Subtitle string `json:"subtitle"`
	Price int `json:"price"`
	Link string `json:"link"`
	Description []string `json:"description"`
}
