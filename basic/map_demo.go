package basic

import "fmt"

func RangePrintMap() {
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}
}
