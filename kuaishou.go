package kuaishou

import (
	"strings"
)

type Data struct {
	Success       bool     `json:"success"`
	Type          int      `json:"type"`
	VideoLink     string   `json:"video_link"`
	ImageLinkList []string `json:"image_link_list"`
}

func WatermarkRemover(url string) (Data, error) {
	data := Data{Success: false}

	result := GetVideoLink(url)

	if result.Err != nil {
		return data, result.Err
	}

	html, err := GetVideoHtml(result.Link, result.Cookies)

	if err != nil {
		return data, err
	}

	if strings.Contains(html, "图集分享") {
		imageLinks := ExtractImageLink(html)
		if len(imageLinks) > 0 {
			data.ImageLinkList = imageLinks
			data.Success = true
		}
	} else {
		videoLink := ExtractVideoLink(html)
		if len(videoLink) > 0 {
			data.VideoLink = videoLink
			data.Success = true
		}
	}

	return data, nil

}
