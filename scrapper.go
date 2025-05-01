package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
	"gorm.io/gorm"
)

func RunScraper(db *gorm.DB) {
	startURLs := map[string]string{
		"http://www.ugc-universities.gov.bd/public-universities":        "Public",
		"http://www.ugc-universities.gov.bd/private-universities":       "Private",
		"http://www.ugc-universities.gov.bd/international-universities": "International",
		"http://www.ugc-universities.gov.bd/cbhe":                       "CBHE",
	}

	listCollector := colly.NewCollector(
		colly.AllowedDomains("www.ugc-universities.gov.bd"),
	)

	detailCollector := listCollector.Clone()

	detailCollector.OnHTML("div.col-md-12", func(e *colly.HTMLElement) {
		detail := make(map[string]string)

		e.ForEach("table.university-detail tbody tr", func(_ int, row *colly.HTMLElement) {
			tds := row.DOM.Find("td")
			if tds.Length() < 2 {
				return
			}
			key := strings.TrimSpace(tds.Eq(0).Text())
			val := strings.TrimSpace(tds.Eq(1).Text())
			if key != "" {
				detail[key] = val
			}
		})

		univName := strings.TrimSpace(e.DOM.Find("h1.page-title").Text())
		if univName == "" {
			fmt.Println("⚠️ Skipping page with no university name:", e.Request.URL.String())
			return
		}

		category := e.Request.Ctx.Get("category")
		detailURL := e.Request.URL.String()

		uni := University{
			Name:              univName,
			Category:          category,
			Website:           detail["Website"],
			DetailURL:         detailURL,
			YearEstablished:   detail["Year of establishment"],
			PermanentCampus:   detail["Permanent campus"],
			ViceChancellor:    detail["Vice Chancellor"],
			ProViceChancellor: detail["Pro Vice Chancellor"],
			Treasurer:         detail["Treasurer"],
			Registrar:         detail["Registrar"],
			Contact:           detail["Contact"],
			Email:             detail["Email"],
			Telephone:         detail["Telephone / Mobile"],
			Fax:               detail["Fax"],
		}

		var existing University
		db.Where("name = ? AND category = ?", uni.Name, uni.Category).First(&existing)
		if existing.ID == 0 {
			db.Create(&uni)
		}
	})

	listCollector.OnHTML("table.table tbody tr", func(e *colly.HTMLElement) {
		detailURL := e.ChildAttr("td:nth-child(2) a", "href")
		category := e.Request.Ctx.Get("category")
		if detailURL != "" {
			ctx := colly.NewContext()
			ctx.Put("category", category)
			detailCollector.Request("GET", detailURL, nil, ctx, nil)
		}
	})

	for url, category := range startURLs {
		ctx := colly.NewContext()
		ctx.Put("category", category)
		fmt.Println("Scraping:", category)
		_ = listCollector.Request("GET", url, nil, ctx, nil)
	}

	listCollector.Wait()
	detailCollector.Wait()
	fmt.Println("✅ Scraping complete and stored in DB")
}
