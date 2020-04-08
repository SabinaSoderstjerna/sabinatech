package aboutpage

import (
	"net/url"
	"time"
)

type AboutPage struct {
	Name                 FullName
	Age                  int
	Body                 string
	SocialMedia          SocialMedia
	Education            []Education
	VocationalExperience []Experience
	NonProfitExperience  []Experience
}

type About struct {
	Name            FullName
	Birthday        string
	SocialMediaKeys SocialMediaKeys
	Body            string
}

type SocialMedia struct {
	LinkedIn  url.URL
	GitHub    url.URL
	Facebook  url.URL
	Instagram url.URL
	Pinterest url.URL
}

type SocialMediaKeys struct {
	LinkedIn  string
	GitHub    string
	Facebook  string
	Instagram string
	Pinterest string
}

type Education struct {
	Type           string
	Name           string
	School         string
	Location       string
	GraduationYear int
	StartYear      int
	Projects       []Project
}

type Project struct {
	Title       string
	Name        string
	Supervisors []FullName
	Body        string
}

type FullName struct {
	FirstName string
	LastName  string
}

type Experience struct {
	Name       string
	SinceMonth time.Month
	SinceYear  int
	ToMonth    time.Month
	ToYear     int
	Position   string
	Body       string
	Projects   []Project
}
