package common

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
)

func GetChallenge(year string, day string) []byte {
	cdir, cerr := os.Getwd()
	if cerr != nil {
		log.Fatalf("Got error on local dir: %s", cerr)
	}
	dir := os.DirFS(fmt.Sprintf("%s/%s/day%s", cdir, year, day))
	file := fmt.Sprintf("%s/input", dir)
	input, err := os.ReadFile(file)
	if err == nil {
		// return binary.BigEndian.Uint64(input)
		return input
	} else {
		cookie := &http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC_COOKIE"),
		}
		jar, cookieerr := cookiejar.New(nil)
		if cookieerr != nil {
			log.Fatalf("Got error while creating cookie jar %s", cookieerr.Error())
		}
		client := http.Client{
			Jar: jar,
		}
		inputURL := fmt.Sprintf("https://%s/%s/day/%s/input", os.Getenv("AOC_URL"), year, day)
		req, reqerr := http.NewRequest("GET", inputURL, nil)
		if reqerr != nil {
			log.Fatalf("Got error %s", reqerr.Error())
		}
		req.AddCookie(cookie)

		resp, resperr := client.Do(req)
		if resperr != nil {
			log.Fatalf("Error occured. Error is: %s", resperr.Error())
		}
		defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading body: %s", err)
		}
		os.WriteFile(file, bodyBytes, 0666)
		// return binary.BigEndian.Uint64(bodyBytes)
		return bodyBytes
	}
}
