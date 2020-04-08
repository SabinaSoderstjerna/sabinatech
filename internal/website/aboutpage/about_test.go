package aboutpage

import (
	"net/url"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewAboutPage(t *testing.T) {
	// Given
	pathToAboutFile := filepath.Join("testfiles", "testfile_about.json")
	l, _ := url.Parse("https://www.linkedin.com/in/linkedin")
	g, _ := url.Parse("https://www.github.com/github")
	f, _ := url.Parse("https://www.facebook.com/profile.php?id=facebook")
	i, _ := url.Parse("https://www.instagram.com/instagram")
	p, _ := url.Parse("https://www.pinterest.se/pinterest")
	expected := AboutPage{
		Name: FullName{
			FirstName: "first",
			LastName:  "last",
		},
		Age:  1,
		Body: "body",
		SocialMedia: SocialMedia{
			LinkedIn:  *l,
			GitHub:    *g,
			Facebook:  *f,
			Instagram: *i,
			Pinterest: *p,
		},
		Education:            nil,
		VocationalExperience: nil,
		NonProfitExperience:  nil,
	}
	// When
	actual := NewAboutPage(pathToAboutFile)
	// Then
	require.Equal(t, expected, actual)
}

func TestAboutPage_GetEducations(t *testing.T) {
	// Given
	a := AboutPage{}
	pathToEducationsFolder := filepath.Join("testfiles", "education")
	expectedEducation := []Education{
		{
			Type:           "Education Type",
			Name:           "Name of education",
			School:         "School",
			Location:       "Location",
			GraduationYear: 2015,
			StartYear:      2012,
			Projects: []Project{
				{
					Title: "Project 1",
					Name:  "Name of project 1",
					Supervisors: []FullName{
						{
							FirstName: "First name 1",
							LastName:  "Last name 1",
						},
						{
							FirstName: "First name 2",
							LastName:  "Last name 2",
						},
					},
				},
			},
		},
	}
	// When
	actual := a.GetEducations(pathToEducationsFolder)
	// Then
	require.Equal(t, expectedEducation, actual)
}

func TestAboutPage_GetExperience(t *testing.T) {
	// Given
	a := AboutPage{}
	pathToEmployeeFolder := filepath.Join("testfiles", "employee")
	expectedExperience := []Experience{
		{
			Name:       "Company",
			SinceMonth: 1,
			SinceYear:  2018,
			ToMonth:    3,
			ToYear:     2020,
			Position:   "Position",
			Body:       "Body",
			Projects: []Project{
				{
					Title:       "",
					Name:        "Project 1",
					Supervisors: nil,
					Body:        "Body to project 1",
				},
				{
					Title:       "",
					Name:        "Project 2",
					Supervisors: nil,
					Body:        "Body to project 2",
				},
			},
		},
	}
	// When
	actual := a.GetExperience(pathToEmployeeFolder)
	// Then
	require.Equal(t, expectedExperience, actual)
}
