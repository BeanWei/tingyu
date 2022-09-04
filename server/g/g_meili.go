package g

import (
	"sync"

	"github.com/meilisearch/meilisearch-go"
)

var (
	meili     *meilisearch.Client
	meiliOnce sync.Once
)

func Meili() *meilisearch.Client {
	meiliOnce.Do(func() {
		meili = meilisearch.NewClient(meilisearch.ClientConfig{
			Host:   Cfg().Meilisearch.Host,
			APIKey: Cfg().Meilisearch.APIKey,
		})
		meili.CreateIndex(&meilisearch.IndexConfig{
			Uid: "posts",
		})
	})
	return meili
}
