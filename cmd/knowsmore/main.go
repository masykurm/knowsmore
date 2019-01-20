package main

import (
	"log"
	"time"

	"gopkg.in/olivere/elastic.v5"

	"github.com/labstack/echo"
	"github.com/robertotambunan/knowsmore/member/api"
	"github.com/robertotambunan/knowsmore/member/repo"
	"github.com/robertotambunan/knowsmore/member/usecase"
)

func main() {
	// still on progress creating service
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()

	mElasticRepo := repo.NewElasticMemberRepository(client, "member", "_doc")
	mUsecase := usecase.NewMemberUsecase(mElasticRepo, 100*time.Millisecond)

	api.NewMemberHTTPtpHandler(e, mUsecase)
	e.Start(":9002")
}
