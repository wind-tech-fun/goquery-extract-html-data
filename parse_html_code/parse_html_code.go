package parse_html_code

import (
	"log"
	"os"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func Run(html_file string) ([]map[string]string) {
	f, err := os.Open(html_file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	var s01 []map[string]string

	classTitleMapping := _classTitleMapping()
	indexMapping 	  := _getIndexMapping()

	html_tag := `table#tablepress-42 > tbody > tr`

	/**
	* GoQuery business logic example
	*/
	doc.Find(html_tag).Each(func(i int, sel * goquery.Selection) {
		_temp := make(map[string]string)
		sel.Find("td").Each(func(i01 int, sel01 * goquery.Selection) {
			_attr, _ := sel01.Attr("class")
			for _, _subStr := range indexMapping {
				if strings.Contains(_attr , _subStr)  {
					_temp[classTitleMapping[_subStr]] = strings.TrimSpace(sel01.Text())
				}
			}
		})
		s01 = append(s01, _temp)
	})

	return s01
}

func _getIndexMapping() ([]string) {
	return []string {
		"column-1",
		"column-2",
		"column-3",
		"column-4",
		"column-5",
		"column-6",
		"column-7",
		"column-8",
	}
}

func _classTitleMapping() (map[string]string) {
	return map[string]string {
		"column-1": "Symbol",
		"column-2": "Company",
		"column-3": "Price",
		"column-4": "Yield",
		"column-5": "Market Cap",
		"column-6": "1d Chg",
		"column-7": "1m Chg",
		"column-8": "12m Chg",
	}
}