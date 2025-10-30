package process

import (
	"goquery-extract-html-data/common"
	"goquery-extract-html-data/parse_html_code"
)

func Run() {
	file := `/app/don_jones_2025_10_30.html`
	data := parse_html_code.Run(file)
	saved_file := `/app/json/data.json`
	common.GenJsonFile(data, saved_file)
}
