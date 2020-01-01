package rand

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexChinese(t *testing.T) {
	// text := "⻋"
	// text := "车1231测试文案,欢迎，。.来【】()（）"
	text := "恭喜！你已通过⻋主审核。现在可以开始接单啦，看看谁和你顺路"
	m, _ := regexp.MatchString("[^\u4e00-\u9fa5\\w【】()（），。.!！]", text)
	fmt.Println(m)

	text2 := "恭喜！你已通过车主审核。现在可以开始接单啦，看看谁和你顺路"
	m, _ = regexp.MatchString("[^\u4e00-\u9fa5\\w【】()（），。.!！]", text2)
	fmt.Println(m)

	// text3 := "⻋"
	text3 := "车"
	m, _ = regexp.MatchString("\\w", text3)
	fmt.Println(m)
}

func TestRegexCase(t *testing.T) {
	text := "A"
	m, _ := regexp.MatchString("a", text)
	fmt.Println(m)

	text2 := "A"
	m2, _ := regexp.MatchString("(?i)a", text2)
	fmt.Println(m2)
}

func TestMatch(t *testing.T) {
	pattern := "([-\\d\\.]+)\\S*[元折券]"
	compileRegex := regexp.MustCompile(pattern)
	result := compileRegex.FindStringSubmatch("送您3.1元立减券")
	fmt.Println(result)
}