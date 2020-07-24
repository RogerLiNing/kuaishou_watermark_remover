package kuaishou

import (
	"strings"
	"testing"
)

func TestAvailableVideo(t *testing.T) {

	url := "https://v.kuaishou.com/7ZtNvs"
	t.Log("测试有效视频短链接：" + url)

	u, err := WatermarkRemover(url)
	if !strings.Contains(u, ".mp4") {
		t.Fail()
	}

	if err != nil {
		t.Fail()
		t.Log(err)
	}

}

func TestUnAvailableVideo(t *testing.T) {

	url := "https://v.douyin.com/JNhu000"
	t.Log("测试无效视频短链接："+url)
	u, err := WatermarkRemover(url)
	t.Log(u)
	if err != nil {
		t.Fail()
		t.Log(err)
	}

	if len(u) != 0 {

		t.Fail()
	}

}