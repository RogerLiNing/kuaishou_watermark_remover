package kuaishou

func WatermarkRemover(url string)(string, error)  {

	videoLink, err := GetVideoLink(url)

	if err != nil {
		return "", err
	}

	html, err := GetVideoHtml(videoLink)

	if err != nil {
		return "", err
	}

	link := ExtractVideoLink(html)

	return link, nil

}

