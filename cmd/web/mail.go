package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"time"
)

type Response struct {
	Location struct {
		Long            float64     `json:"long"`
		CountryOrRegion string      `json:"countryOrRegion"`
		ProvinceOrState interface{} `json:"provinceOrState"`
		County          interface{} `json:"county"`
		IsoCode         string      `json:"isoCode"`
		Lat             float64     `json:"lat"`
	} `json:"location"`
	UpdatedDateTime time.Time `json:"updatedDateTime"`
	News            []struct {
		Path              string      `json:"path"`
		Title             string      `json:"title"`
		Excerpt           string      `json:"excerpt"`
		Heat              int         `json:"heat"`
		Tags              []string    `json:"tags"`
		Type              string      `json:"type"`
		WebURL            string      `json:"webUrl"`
		AmpWebURL         string      `json:"ampWebUrl"`
		CdnAmpWebURL      string      `json:"cdnAmpWebUrl"`
		PublishedDateTime string      `json:"publishedDateTime"`
		UpdatedDateTime   interface{} `json:"updatedDateTime"`
		Provider          struct {
			Name       string      `json:"name"`
			Domain     string      `json:"domain"`
			Images     interface{} `json:"images"`
			Publishers interface{} `json:"publishers"`
			Authors    interface{} `json:"authors"`
		} `json:"provider"`
		Images []struct {
			URL         string      `json:"url"`
			Width       int         `json:"width"`
			Height      int         `json:"height"`
			Title       string      `json:"title"`
			Attribution interface{} `json:"attribution"`
		} `json:"images"`
		Locale     string   `json:"locale"`
		Categories []string `json:"categories"`
		Topics     []string `json:"topics"`
	} `json:"news"`
}

func mail(address string) {
	// API 연동
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.smartable.ai/coronavirus/news/US", nil)
	req.Header.Set("Subscription-Key", "731335f75df5479ebc467601772365f0")
	res, _ := client.Do(req)

	responseData, responseError := ioutil.ReadAll(res.Body)
	if responseError != nil {
		log.Fatal(responseError)
	}

	response := Response{}
	jsonErr := json.Unmarshal(responseData, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// 뉴스 제목
	var newsTitle, newsURL, newsPublishedDateAndTime, newsImage string

	for _, v := range response.News {
		newsTitle = v.Title
		newsURL = v.WebURL
		newsPublishedDateAndTime = v.PublishedDateTime
		for _, k := range v.Images {
			newsImage = k.URL
		}
	}

	// 메일서버 로그인 정보 설정
	auth := smtp.PlainAuth("", "your gmail user id", "your gmail user password", "smtp.gmail.com")

	from := "legitcode267@gmail.com"
	to := []string{address} // 복수 수신자 가능

	// 메시지 작성
	headerSubject := "Subject: Corona Today\r\n"
	headerBlank := "\r\n"
	body := newsTitle + "URL:" + newsURL + "\n" + newsPublishedDateAndTime
	msg := []byte(headerSubject + headerBlank + body)

	// 메일 보내기
	err := smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		panic(err)
	}
}
