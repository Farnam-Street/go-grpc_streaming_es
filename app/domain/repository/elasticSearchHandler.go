package repository

import (
	"context"
	"encoding/json"
	"farnam-street-api-go/models"
	"fmt"
)

func GETESClient()(*elastic.Client, error){
	client, err :=  elastic.NewClient(elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return client, err

}

func countESByMatchQuery(fieldKey string, fieldValue string) int{
	ctx := context.Background()
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(fieldKey, fieldValue))

	client, err := GETESClient()
	searchService := client.Search().Index("asset").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return 0
	}
	return searchResult.TotalHits()
}

func rangeAggregationByValue(){

}

func countByRangeQuery(fieldKey string, fromValue int, toValue int) int{
	ctx := context.Background()
	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewRangeQuery(fieldKey).From(fromValue).To(toValue))

	client, err := GETESClient()
	searchService := client.Search().Index("asset").SearchSource(searchSource)

	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("[ProductsES][GetPIds]Error=", err)
		return 0
	}
	return searchResult.TotalHits()
	
}


