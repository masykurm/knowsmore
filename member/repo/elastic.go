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
}

// NewElasticMemberRepository will create an implementation of member.Repository
func NewElasticMemberRepository(client *elastic.Client) member.Repository {
	return &elasticMemberRepo{client}
}

func (e *elasticMemberRepo) GetByID(ctx context.Context, id string) (member m.Member, err error) {
	log.Panic("GetByID method has not implemented yet")
	return
}

func (e *elasticMemberRepo) GetByAutocomplete(ctx context.Context, keyword string) (members []m.Member, err error) {
	tagSuggester := elastic.NewCompletionSuggester("data").Text("r").Field("suggest").Size(10)
	searchSource := elastic.NewSearchSource().Suggester(tagSuggester)
	searchResult, err := e.Client.Search().Index("member").Type("_doc").SearchSource(searchSource).Do(ctx)
	for _, ops := range searchResult.Suggest["data"] {
		for _, op := range ops.Options {
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
