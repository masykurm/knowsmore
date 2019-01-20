package repo

import (
	"context"
	"encoding/json"
	"log"

	"gopkg.in/olivere/elastic.v5"

	"github.com/robertotambunan/knowsmore/member"
	m "github.com/robertotambunan/knowsmore/member/model"
)

type elasticMemberRepo struct {
	Client *elastic.Client
	Index  string
	Type   string
}

// NewElasticMemberRepository will create an implementation of member.Repository
func NewElasticMemberRepository(client *elastic.Client, index, types string) member.Repository {
	return &elasticMemberRepo{
		Client: client,
		Index:  index,
		Type:   types,
	}
}

func (e *elasticMemberRepo) GetByAutocomplete(ctx context.Context, keyword string) (members []m.Member, err error) {
	members = []m.Member{}
	searchSuggester := elastic.NewCompletionSuggester("data").Text(keyword).Field("suggest").Size(10)
	searchSource := elastic.NewSearchSource().Suggester(searchSuggester)
	searchResult, err := e.Client.Search().Index(e.Index).Type(e.Type).SearchSource(searchSource).Do(ctx)
	for _, ops := range searchResult.Suggest["data"] {
		for _, op := range ops.Options {
			if op.Source == nil {
				continue
			}
			var member m.Member
			err := json.Unmarshal(*op.Source, &member)
			if err != nil {
				log.Println(err)
				continue
			}
			members = append(members, member)
		}
	}
	return
}
