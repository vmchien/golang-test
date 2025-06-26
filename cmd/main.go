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
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3521-3934303534352D36332D333532312D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3522-3934303534352D36332D333532322D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3523-3934303534352D36332D333532332D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3524-3934303534352D36332D333532342D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3525-3934303534352D36332D333532352D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3526-3934303534352D36332D333532362D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3527-3934303534352D36332D333532372D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3528-3934303534352D36332D333532382D3135373338322D30",
		"https://hotruyen1.com/chuong/chien-than-o-re-duong-thanh-tan-thanh-tam-bat-bai-chien-than-duong-than-chuong-3529-3934303534352D36332D333532392D3135373338322D30",
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
