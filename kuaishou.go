package kuaishou

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

	// 直接运行判断，非此即彼

	imageLinks := ExtractImageLink(html)
	videoLink := ExtractVideoLink(html)

	// 0 是视频，1是图集
	if len(imageLinks) > 0 {
		data.ImageLinkList = imageLinks
		data.Success = true
		data.Type = 1

	} else if len(videoLink) > 0 {
		data.VideoLink = videoLink
		data.Success = true
		data.Type = 0
	}

	return data, nil

}
