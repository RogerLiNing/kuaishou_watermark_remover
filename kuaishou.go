package kuaishou

func WatermarkRemover(url string)(string, error)  {

	result := GetVideoLink(url)

	if result.Err != nil {
		return "", result.Err
	}

	html, err := GetVideoHtml(result.Link, result.Cookies)

	if err != nil {
		return "", err
	}

	link := ExtractVideoLink(html)

	return link, nil

}

