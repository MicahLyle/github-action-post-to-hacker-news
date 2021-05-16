package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/lukakerr/hkn"
)

func main() {
	release, githubRefFound := os.LookupEnv("GITHUB_REF")
	if githubRefFound == false {
		log.Fatal("GITHUB_REF environment variable not set.")
	}
	release = strings.TrimLeft(release, "vV")
	username, usernameFound := os.LookupEnv("HN_USERNAME")
	if usernameFound == false {
		log.Fatal("HN_USERNAME environment variable not set.")
	}
	password, passwordFound := os.LookupEnv("HN_PASSWORD")
	if passwordFound == false {
		log.Fatal("HN_PASSWORD environment variable not set.")
	}
	titleSpecifier, titleSpecifierFound := os.LookupEnv("HN_TITLE_FORMAT_SPECIFIER")
	if titleSpecifierFound == false {
		log.Fatal("HN_TITLE_FORMAT_SPECIFIER environment variable not set.")
	}
	urlSpecifier, urlSpecifierFound := os.LookupEnv("HN_URL_FORMAT_SPECIFIER")
	if urlSpecifierFound == false {
		log.Fatal("HN_URL_FORMAT_SPECIFIER environment variable not set.")
	}
	// If `HN_TEST_MODE` is present in the environment we'll just print the
	// result at the end instead of actually posting.
	testMode, testModeFound := os.LookupEnv("HN_TEST_MODE")
	inTestMode := testModeFound && (testMode == "true" || testMode == "True")

	title := fmt.Sprintf(titleSpecifier, release)
	url := fmt.Sprintf(urlSpecifier, release)

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
