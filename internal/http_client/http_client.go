package httpclient

import (
	"log"
	"net/http"
)

var (
	cybersportUrl string = "https://www.cybersport.ru"
	httpClient           = http.Client{}
)

func DoGET() (resp *http.Response) {
	request, errRequest := http.NewRequest("GET", cybersportUrl, nil)

	if errRequest != nil {
		log.Fatal(errRequest)
	}

	resp, errDo := httpClient.Do(request)

	if errDo != nil {
		log.Fatal(errDo)
	}

	if resp.StatusCode != 200 {
		log.Fatal("StatusCode is:", resp.StatusCode)
	}

	return
}
