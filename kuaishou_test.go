package kuaishou

import (
	"testing"
)

func TestAvailableVideo(t *testing.T) {

	url := "https://v.kuaishou.com/5u1OpM" //  https://v.kuaishouapp.com/s/4znp78gd https://v.kuaishou.com/5u1OpM https://v.kuaishouapp.com/s/o0rrbRHb  https://v.kuaishou.com/79cv50
	t.Log("测试有效视频短链接：" + url)

	u, _ := WatermarkRemover(url)
	t.Log(u)
	if !u.Success {
		t.Fail()
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

	if len(u.VideoLink) != 0 {

		t.Fail()
	}

}