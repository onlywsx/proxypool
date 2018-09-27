package getter

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-clog/clog"
	"github.com/henson/proxypool/pkg/models"
	"github.com/parnurzeal/gorequest"
)

// GBJ get ip from goubanjia.com
func GBJ() (result []*models.IP) {
	pollURL := "http://www.goubanjia.com/"

	resp, _, errs := gorequest.New().Get(pollURL).
		Set("User-Agent", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0`).
		End()
	if errs != nil {
		log.Println(errs)
		return
	}
	// fmt.Println(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println(errs)
		return
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	doc.Find("body div.container div.row div.container-fluid div.row-fluid div.span12 tbody tr").Each(func(_ int, s *goquery.Selection) {
		sf, _ := s.Find(".ip").Html()
		tee := regexp.MustCompile("<pstyle=\"display:none;\">.?.?</p>").ReplaceAllString(strings.Replace(sf, " ", "", -1), "")
		re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
		ip := models.NewIP()
		ip.Data = re.ReplaceAllString(tee, "")
		ip.Type1 = s.Find("td:nth-child(3) > a").Text()
		// fmt.Printf("ip.Data = %s , ip.Type = %s\n", ip.Data, ip.Type1)
		clog.Info("[GBJ] ip.Data: %s,ip.Type: %s", ip.Data, ip.Type1)
		result = append(result, ip)
	})

	clog.Info("GBJ done.")
	return
}
