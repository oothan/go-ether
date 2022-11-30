package main

import (
	"fmt"
	"strings"
)

func main() {
	strs := "resources.91douyin.club:tearesources01.com:www.xinghua03.com"
	str := "images.23ntt.com"
	testStr := "https://resources.91douyin.club/static/vod/20200812/38c02534bdc960418220f3c19a234089.png"

	var res1 string
	if strs != "" && str != "" {
		rSplits := strings.Split(strs, ":")
		if len(rSplits) > 0 {
			fmt.Println(len(rSplits))
			for _, s := range rSplits {
				if strings.Contains(testStr, s) {
					res1 = strings.ReplaceAll(testStr, s, str)
				}
			}
		}
	}

	fmt.Println(str)
	fmt.Println(res1)
}
