package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// URL của chương truyện
	urls := []string{
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4460-3934303534352D36332D343436302D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4461-3934303534352D36332D343436312D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4462-3934303534352D36332D343436322D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4463-3934303534352D36332D343436332D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4464-3934303534352D36332D343436342D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4465-3934303534352D36332D343436352D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4466-3934303534352D36332D343436362D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4467-3934303534352D36332D343436372D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4468-3934303534352D36332D343436382D3133333033362D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-4469-3934303534352D36332D343436392D3133333033362D30",
	}

	for _, v := range urls {
		time.Sleep(time.Millisecond * 600)
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
	doc.Find("#borderchapter p.padtext").Each(func(i int, s *goquery.Selection) {
		content := strings.TrimSpace(s.Text())
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
