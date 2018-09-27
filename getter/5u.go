package getter

import (
	"log"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-clog/clog"
	"github.com/henson/proxypool/pkg/models"
	"github.com/parnurzeal/gorequest"
)

//Data5u is not work now
// Data5u get ip from data5u.com
func Data5u() (result []*models.IP) {
	pollURL := "http://www.data5u.com/free/index.shtml"
	resp, _, errs := gorequest.New().Get(pollURL).
		Set("User-Agent", `Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0`).
		End()
	if errs != nil {
		log.Println(errs)
		return
	}
	if resp.StatusCode != 200 {
		log.Println(errs)
		return
	}
	// fmt.Println(resp.Body)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err.Error())
		return
	}

	doc.Find("body > div.wlist > ul > li:nth-child(2) > ul").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		node := strconv.Itoa(i + 1)
		ss := s.Find("ul:nth-child(" + node + ") > span:nth-child(1) > li").Text()
		sss := s.Find("ul:nth-child(" + node + ") > span:nth-child(2) > li").Text()
		ssss := s.Find("ul:nth-child(" + node + ") > span:nth-child(4) > li").Text()
		ip := models.NewIP()
		ip.Data = ss + ":" + sss
		ip.Type1 = ssss
		// fmt.Printf("ip.Data = %s, ip.Type = %s", ip.Data, ip.Type1)
		clog.Info("[Data5u] ip.Data: %s,ip.Type: %s", ip.Data, ip.Type1)
		result = append(result, ip)
	})
	clog.Info("Data5u done.")
	return
}
