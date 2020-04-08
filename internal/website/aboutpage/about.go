package aboutpage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"time"
)

func NewAboutPage(path string) AboutPage {
	about := getAbout(path)
	return AboutPage{
		Name:        about.Name,
		Age:         getAge(about.Birthday),
		Body:        about.Body,
		SocialMedia: getSocialMedia(about.SocialMediaKeys),
	}
}

func getAbout(path string) About {
	var info About
	body, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	if err = json.Unmarshal(body, &info); err != nil {
		fmt.Printf("error unmarshalling: %v\n", err)
	}
	return info
}

func getAge(birthday string) int {
	todayYear, _, _ := time.Now().Date()
	time, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	birthYear, _, _ := time.Date()
	return todayYear - birthYear
}

const scheme = "https"

func getSocialMedia(keys SocialMediaKeys) SocialMedia {
	linkedin := url.URL{
		Scheme: scheme,
		Host:   "www.linkedin.com",
		Path:   "/in/" + keys.LinkedIn,
	}
	github := url.URL{
		Scheme: scheme,
		Host:   "www.github.com",
		Path:   "/" + keys.GitHub,
	}
	facebook := url.URL{
		Scheme:   scheme,
		Host:     "www.facebook.com",
		Path:     "/profile.php",
		RawQuery: "id=" + keys.Facebook,
	}
	instagram := url.URL{
		Scheme: scheme,
		Host:   "www.instagram.com",
		Path:   "/" + keys.Instagram,
	}
	pinterest := url.URL{
		Scheme: scheme,
		Host:   "www.pinterest.se",
		Path:   "/" + keys.Pinterest,
	}
	return SocialMedia{
		LinkedIn:  linkedin,
		GitHub:    github,
		Facebook:  facebook,
		Instagram: instagram,
		Pinterest: pinterest,
	}
}

func (a *AboutPage) GetEducations(path string) []Education {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	info := make([]Education, len(files))
	for i, file := range files {
		body, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("error read file: %v\n", err)
		}
		if err = json.Unmarshal(body, &info[len(files)-1-i]); err != nil {
			fmt.Printf("error json unmarshal: %v\n", err)
		}
	}
	return info
}

func (a *AboutPage) GetExperience(path string) []Experience {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	info := make([]Experience, len(files))
	for i, file := range files {
		body, err := ioutil.ReadFile(filepath.Join(path, file.Name()))
		if err != nil {
			fmt.Printf("error read file: %v\n", err)
		}
		if err = json.Unmarshal(body, &info[len(files)-1-i]); err != nil {
			fmt.Printf("error json unmarshal: %v\n", err)
		}
	}
	return info
}
