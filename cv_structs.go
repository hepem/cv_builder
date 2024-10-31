package main

type CV struct {
	Keywords   Keywords
	Name       string
	Birthday   string
	Place      string
	Resume     string
	Email      string
	Phone      string
	Social     SocialLinks
	Education  []Education
	Experience []Experience
	Projects   []Projects
}

type Keywords struct {
	Experience string
	Education  string
	Projects   string
	Tech       string
}

type SocialLinks struct {
	LinkedIn string
	GitHub   string
}

type Education struct {
	Place string
	Date  string
	Title string
}

type Experience struct {
	Date     string
	Place    string
	Company  string
	Role     string
	Comments []string
}

type Projects struct {
	Link     string
	Title    string
	Comments []string
}

// CVData mirrors the JSON structure for flexible parsing
type CVData struct {
	Keywords   Keywords     `json:"keywords"`
	Name       string       `json:"name"`
	Birthday   string       `json:"birthday"`
	Place      string       `json:"place"`
	Resume     string       `json:"resume"`
	Email      string       `json:"email"`
	Phone      string       `json:"phone"`
	Social     SocialLinks  `json:"social"`
	Education  []Education  `json:"education"`
	Experience []Experience `json:"experience"`
	Projects   []Projects   `json:"projects"`
}
