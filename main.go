package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lukakerr/hkn"
)

func main() {
	celery_release, celeryReleaseFound := os.LookupEnv("CELERY_RELEASE")
	if celeryReleaseFound == false {
		log.Fatal("CELERY_RELEASE environment variable not set.")
	}
	celery_release = strings.TrimLeft(celery_release, "vV")
	username, usernameFound := os.LookupEnv("HN_USERNAME")
	if usernameFound == false {
		log.Fatal("HN_USERNAME environment variable not set.")
	}
	password, passwordFound := os.LookupEnv("HN_PASSWORD")
	if passwordFound == false {
		log.Fatal("HN_PASSWORD environment variable not set.")
	}
	// If `HN_TEST_MODE` is present in the environment we'll just print the
	// result at the end instead of actually posting.
	testMode, testModeFound := os.LookupEnv("HN_TEST_MODE")
	inTestMode := testModeFound && (testMode == "true" || testMode == "True")

	title := fmt.Sprintf("Celery v%s Released!", celery_release)
	url := fmt.Sprintf("https://docs.celeryproject.org/en/v%s/changelog.html", celery_release)

	if inTestMode {
		fmt.Println("Prepared Data For Test Mode.")
		fmt.Println("Title:", title)
		fmt.Println("URL:", url)
		fmt.Println("Test Mode finished for posting story to Hacker News!")
	} else {

		client := hkn.NewClient()

		cookie, loginErr := client.Login(username, password)
		if loginErr != nil {
			log.Fatal("Failed to log in with given credentials.")
		}

		_, createStoryErr := client.CreateStoryWithURL(title, url, cookie)
		if createStoryErr != nil {
			log.Fatal(createStoryErr)
		}

		fmt.Println("Successfully posted story to Hacker News!")
	}
}
