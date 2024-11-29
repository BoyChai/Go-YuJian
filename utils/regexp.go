package utils

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
)

// URL校验
func CheckURL(s string) bool {
	reg := regexp.MustCompile(`([hH][tT]{2}[pP]://|[hH][tT]{2}[pP][sS]://|[wW]{3}.|[wW][aA][pP].|[fF][tT][pP].|[fF][iI][lL][eE].)[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	return reg.MatchString(s)
}

// 纯数字校验
func CheckDigits(s string) bool {
	reg := regexp.MustCompile(`^\d+$`)
	return reg.MatchString(s)
}

// ExtractDomain 提取域名
func ExtractDomain(url string) string {
	splitURL := strings.Split(url, "://")
	if len(splitURL) < 2 {
		return ""
	}

	re := regexp.MustCompile(`^([a-zA-Z0-9.-]+)`)
	match := re.FindString(splitURL[1])

	return match
}

// 保存文件的名字
func GetSaveFileURLName(input string) string {
	parsedURL, err := url.Parse(input)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}

	scheme := parsedURL.Scheme
	host := parsedURL.Hostname()
	port := parsedURL.Port()

	if scheme == "http" {
		if port != "" {

			return fmt.Sprintf("http_%s_%s", host, port)
		}

		return fmt.Sprintf("http_%s", host)
	} else if scheme == "https" {
		if port != "" {

			return fmt.Sprintf("https_%s_%s", host, port)
		}

		return fmt.Sprintf("https_%s", host)
	}

	return ""
}
