package kuaishou

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetVideoHtml(url string)(string, error)  {

	client := &http.Client{}
	request, err := http.NewRequest("GET", url, nil)

	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("Host", "c.kuaishou.com")
	request.Header.Add("Cookie","did=web_d6c2effa98354bcbaa518816cfdd3d45; didv=1594704296000; clientid=3; client_key=65890b29; Hm_lvt_86a27b7db2c5c0ae37fee4a8a35033ee=1594703274,1594728587,1595351859,1595352041")


	if err != nil {
		log.Fatal(err)
		return "", err
	}

	response, _ := client.Do(request)

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
		return "", nil
	}

	defer response.Body.Close()


	return string(body), nil
}

