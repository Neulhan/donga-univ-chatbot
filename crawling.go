package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type haksikType struct {
	name string
	list []string
}

func trimInner(s string) string {
	returnString := ""
	for i, s := range strings.Split(s, "\n") {
		if s != "" {
			returnString += s
			if i != len(strings.Split(s, "\n")) {
				returnString += "\n"
			}
		}
	}
	return returnString
}

func (h *haksikType) toText() (text string) {
	text = ""
	구분 := [3]string{"정식", "일품", "양분식"}
	text += "\n◼ " + h.name + "\n"
	for i, v := range h.list {
		//구분[i]
		text += "\n▪ " + 구분[i] + "\n" + trimInner(v)
	}
	return
}

func crawlingHaksik(campus string) (text string) {
	fmt.Println(campus)
	text = ""
	res, err := http.Get(haksikEndpoint)
	if err != nil {
		fmt.Println("crawling(haksik): ", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("query(haksik): ", err)
	}
	doc.Find(".gzTable").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		if i == 1 && campus == "승학" {
			리스트 := [5]haksikType{
				haksikType{"분류", []string{}},
				haksikType{"교수회관", []string{}},
				haksikType{"학생회관", []string{}},
				haksikType{"공과대학", []string{}},
				haksikType{"도서관", []string{}},
			}
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				리스트[i%5].list = append(리스트[i%5].list, s.Text())
			})
			text += "🏫 승학캠퍼스\n"
			for i, v := range 리스트 {
				if i != 0 {
					text += v.toText()
				}
			}
		}

		if i == 2 && campus == "구덕-부민" {
			리스트 := [5]haksikType{
				haksikType{"분류", []string{}},
				haksikType{"학생회관", []string{}},
				haksikType{"국제회관 기숙사", []string{}},
				haksikType{"부민(교직원)", []string{}},
				haksikType{"강의동 학생회관", []string{}},
			}

			s.Find("td").Each(func(i int, s *goquery.Selection) {
				리스트[i%5].list = append(리스트[i%5].list, s.Text())
			})
			text += "🏫 구덕/부민 캠퍼스\n"
			for i, v := range 리스트 {
				if i != 0 {
					text += v.toText()
				}
			}
		}
	})

	return
}

func crwalingLibrary() string {

	text := ""
	res, err := http.Get(libraryEndpoint)
	if err != nil {
		fmt.Println("crawling(library): ", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("query(library): ", err)
	}

	li := []string{
		"0", "1", "2",
		"한림 그룹스터디실(4층)", "한림 열람실A(5층)", "한림 열람실B(5층)",
		"한림 열람실C(5층)", "한림 열람실D(5층)", "한림 열람실E(노트북 열람석)",
		"한림 열람실F(5층)", "부민 열람실1-A(5층)", "부민 열람실1-B(5층)",
		"부민 열람실2-A(6층)", "부민 열람실2-B(6층)", "부민 전자자료실(7층)",
		"부민 연속간행물실(8층)", "부민 인문과학실(9층)", "부민 사회과학실(10층)",
	}
	doc.Find("tr").Each(func(i int, s *goquery.Selection) {
		if i > 2 {
			text += li[i] + "\n"
			data := [3]string{}
			s.Find("td:nth-child(n+2):nth-child(-n+4) font").Each(func(i int, s *goquery.Selection) {
				data[i] = s.Text()
			})
			fmt.Println(data)
			text += "[" + data[1] + " / " + data[0] + " ] - 잔여좌석(" + data[2] + " )\n\n"
		}
	})

	return text
}

func crawlingInformation() (text string) {
	text = ""
	res, err := http.Get(informationEndpoint)
	if err != nil {
		fmt.Println("crawling(library): ", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("query(library): ", err)
	}

	doc.Find(".gzTable tr td:nth-child(-n+2)").Each(func(i int, s *goquery.Selection) {
		text += "잔여: " + s.Text() + strings.Repeat("\n", i%2+1)
	})
	text = text[:len(text)-2]

	return
}

func getWeatherData(nx string, ny string) (string, string) {
	var weatherImage string
	var skyText string
	data := make(map[string]string)

	t := time.Now()
	fmt.Println(t.Format(time.RFC3339)[:10])
	fmt.Println(strconv.Itoa(t.Hour()))
	url := weatherEndpoint + "?" + "serviceKey=" + weatherKey + "&base_time=" + timeMapping[strconv.Itoa(t.Hour())] + "00&base_date=" + strings.ReplaceAll(t.Format(time.RFC3339)[:10], "-", "") + "&nx=" + nx + "&ny=" + ny
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("crawling(weather): ", err)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("query(weather): ", err)
	}

	fmt.Println(doc.Text())
	doc.Find("item").Each(func(index int, s *goquery.Selection) {
		category := s.Find("category").Text()
		value := s.Find("fcstValue").Text()
		fmt.Println(category, categoryMatching[category])
		data[categoryMatching[category]] = value
	})

	fmt.Println(data)
	fmt.Println(data["하늘상태"])

	if data["하늘상태"] == "1" {
		weatherImage = sunnyURL
		skyText = "\n날씨: 맑음"
	} else if data["하늘상태"] == "3" || data["하늘상태"] == "4" {
		weatherImage = cloudyURL
		skyText = "\n날씨: 흐림"
	} else if data["강수형태"] == "1" || data["강수형태"] == "4" {
		weatherImage = rainyURL
		skyText = "\n날씨: 비"
	} else if data["강수형태"] == "2" || data["강수형태"] == "3" {
		weatherImage = snowyURL
		skyText = "\n날씨: 눈"
	} else {
		weatherImage = cloudyURL
		skyText = "\n날씨 : 비(" + data["강수확률"] + "%)"
	}

	text := fmt.Sprintf("%s\n기온: %s", skyText, data["3시간 기온"])
	if data["강수확률"] != "" {
		text += "\n강수확률: " + data["강수확률"] + "%"
	}

	if data["습도"] != "" {
		text += "\n습도: " + data["습도"] + "%"
	}
	// fmt.Println("NODATA: ", data["ㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋㅋ"] == "")
	return weatherImage, text
}
