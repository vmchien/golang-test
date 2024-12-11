package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL của chương truyện
	urls := []string{
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1570-15.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1571-14.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1572-15.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1573-15.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1574-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1575-15.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1576-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1577-15.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1578-17.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1579-16.html",

		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1580-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1581-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1582-17.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1583-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1584-17.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1585-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1586-16.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1587-18.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1588-17.html",
		"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1589-16.html",

		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1560-15.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1561-16.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1562-15.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1563-16.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1564-17.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1565-16.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1566-15.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1567-16.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1568-16.html",
		//"https://metruyenfull.org/chien-than-bat-bai-tieu-chinh-van/chuong-1569-16.html",
	}

	for _, v := range urls {
		read(v)
	}

}

func read(url string) {
	// Gửi yêu cầu GET
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching URL:", err)
		return
	}
	defer resp.Body.Close()

	// Kiểm tra mã phản hồi
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Error: %d\n", resp.StatusCode)
		return
	}

	// Tải nội dung HTML
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Parse HTML bằng goquery
	doc, err := goquery.NewDocumentFromReader(stringReader(string(body)))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		return
	}

	// Tìm thẻ chứa nội dung truyện (thay đổi theo cấu trúc trang web)
	doc.Find(".chapter-c").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		lines := strings.Split(content, "\n")
		for _, line := range lines {
			if strings.Contains(line, "Click Theo Dõi ->") {
				continue
			}
			if strings.Contains(line, "Các bạn đang đọc truyện tại metruyenfull.org") {
				break
			}

			fmt.Println(line)
		}
	})
}

// stringReader chuyển chuỗi thành io.Reader
func stringReader(s string) io.Reader {
	return io.NopCloser(io.Reader(&stringReaderStruct{s}))
}

// stringReaderStruct là một tiện ích nhỏ để thực hiện stringReader
type stringReaderStruct struct {
	data string
}

func (sr *stringReaderStruct) Read(p []byte) (n int, err error) {
	n = copy(p, sr.data)
	if n < len(sr.data) {
		sr.data = sr.data[n:]
		return n, nil
	}
	sr.data = ""
	return n, io.EOF
}
