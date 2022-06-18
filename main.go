package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
)


type Rss struct {

	XMLName 	xml.Name		`xml:"rss"`
	Channel		*Channel		`xml:"channel"`
}


type Channel struct {
	Title 		string			`xml:"title"`
	ItemList 	[]Item			`xml:"item"`
}

type Item struct {
	Title			string			`xml:"title"`
	Link 			string			`xml:"link"`
	Traffic		string			`xml:"approx_traffic"`
	NewsItems	[]News			`xml:"news_item"`
}
 
type News struct {
	Headline string   			`xml:"news_item_title"`
	HeadlineLink string 		`xml:"news_item_url"`
}

func main(){
	var r Rss
	data := readGoogleTrends()
	err := xml.Unmarshal(data, &r)
	if err != nil {
		fmt.Println("error" , err)
		
	}

  fmt.Println("\n below are all the google trends for today !")
	fmt.Println("----------------------------")

	for i := range r.Channel.ItemList {
		
		rank := i + 1
		fmt.Println("#", rank)
		fmt.Println("SearchTerm", r.Channel.ItemList[i].Title) 
		fmt.Println("Link to the new Trend", r.Channel.ItemList[i].Link)

	}
}

func getGoogleTrends() *http.Response{
  resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")

	if err!= nil {
		fmt.Println(err)
     os.Exit(1)
	}

	return resp;

}

func readGoogleTrends()[]byte{
   resp := getGoogleTrends();
	// defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	
		return body		
}
