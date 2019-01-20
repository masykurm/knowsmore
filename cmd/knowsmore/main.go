package main

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"

	"github.com/robertotambunan/knowsmore/member/repo"
)

func main() {
	// still on progress creating service
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}

	eRepo := repo.NewElasticMemberRepository(client, "member", "_doc")
	eRepo.GetByAutocomplete(context.Background(), "r")
}
