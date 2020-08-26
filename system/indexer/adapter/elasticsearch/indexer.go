package elasticsearch

import (
	"fmt"
	"io"

	"github.com/muharihar/d3ta-go/system/indexer/adapter"

	es6 "github.com/elastic/go-elasticsearch/v6"
	es7 "github.com/elastic/go-elasticsearch/v7"
	es8 "github.com/elastic/go-elasticsearch/v8"
)

// ESVersion represent Elastic Version Type
type ESVersion string

const (
	ESVersion6 ESVersion = "6"
	ESVersion7 ESVersion = "7"
	ESVersion8 ESVersion = "8"
)

// NewIndexer new Elastic Search Indexer
func NewIndexer(version ESVersion, cfg interface{}) (adapter.IIndexerEngine, error) {
	var err error

	idx := &Indexer{
		esVersion: version,
	}

	cfgType := fmt.Sprintf("%T", cfg)
	if cfgType != "elasticsearch.Config" {
		return nil, fmt.Errorf("Invalid Configuration Type (should be: `elasticsearch.Config`)")
	}

	switch version {
	case ESVersion6:
		if idx.engine, err = NewIndexerES6(cfg.(es6.Config)); err != nil {
			return nil, err
		}
	case ESVersion7:
		if idx.engine, err = NewIndexerES7(cfg.(es7.Config)); err != nil {
			return nil, err
		}
	case ESVersion8:
		if idx.engine, err = NewIndexerES8(cfg.(es8.Config)); err != nil {
			return nil, err
		}
	}

	return idx, nil
}

// Indexer type
type Indexer struct {
	esVersion ESVersion
	engine    adapter.IIndexerEngine
}

// GetEngine get Indexer Engine
func (i *Indexer) GetEngine() interface{} {
	return i.engine
}

func (i *Indexer) IndexExist(indexs []string) (bool, error) {
	return i.engine.IndexExist(indexs)
}

func (i *Indexer) CreateIndex(index string, mapping io.Reader) error {
	return i.engine.CreateIndex(index, mapping)
}

func (i *Indexer) DropIndex(indexs []string) error {
	return i.engine.DropIndex(indexs)
}

func (i *Indexer) DocExist(index string, id string) (bool, error) {
	return i.engine.DocExist(index, id)
}

func (i *Indexer) CreateDoc(index string, id string, body io.Reader) error {

	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("Index Document already exist: %s[id=%s]", index, id)
	}
	return i.engine.CreateDoc(index, id, body)
}

func (i *Indexer) GetDoc(index string, id string) ([]byte, error) {
	return i.engine.GetDoc(index, id)
}

func (i *Indexer) UpdateDoc(index string, id string, body io.Reader) error {

	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Index Document does not exist: %s[id=%s]", index, id)
	}
	return i.engine.UpdateDoc(index, id, body)
}

func (i *Indexer) DeleteDoc(index string, id string) error {
	exist, err := i.DocExist(index, id)
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("Index Document does not exist: %s[id=%s]", index, id)
	}
	return i.engine.DeleteDoc(index, id)
}
