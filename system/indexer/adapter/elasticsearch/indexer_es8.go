package elasticsearch

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	es8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"github.com/muharihar/d3ta-go/system/indexer/adapter"
)

// NewIndexerES8 new Elastic Search 8 Indexer
func NewIndexerES8(cfg es8.Config) (adapter.IIndexerEngine, error) {
	var err error

	idx := &IndexerES8{
		esVersion: ESVersion8,
	}
	idx.ctx = context.Background()
	if idx.engine, err = es8.NewClient(cfg); err != nil {
		return nil, err
	}

	return idx, nil
}

// IndexerES8 type
type IndexerES8 struct {
	ctx       context.Context
	esVersion ESVersion
	engine    *es8.Client
}

// GetEngine get Indexer Engine
func (i *IndexerES8) GetEngine() interface{} {
	return i.engine
}

func (i *IndexerES8) Search(query io.Reader, prettify bool) ([]byte, error) {
	var res *esapi.Response
	var err error

	if prettify {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithPretty(),
		)
	} else {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithBody(query),
		)
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (i *IndexerES8) SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error) {
	var res *esapi.Response
	var err error

	if prettify {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithIndex(index),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithSize(size),
			i.engine.Search.WithPretty(),
		)

	} else {
		res, err = i.engine.Search(
			i.engine.Search.WithContext(i.ctx),
			i.engine.Search.WithIndex(index),
			i.engine.Search.WithBody(query),
			i.engine.Search.WithSize(size),
		)
	}
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (i *IndexerES8) IndexExist(indexs []string) (bool, error) {
	res, err := i.engine.Indices.Exists(indexs)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

func (i *IndexerES8) CreateIndex(index string, mapping io.Reader) error {
	res, err := i.engine.Indices.Create(index, i.engine.Indices.Create.WithBody(mapping))
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (i *IndexerES8) DropIndex(indexs []string) error {
	res, err := i.engine.Indices.Delete(indexs)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (i *IndexerES8) DocExist(index string, id string) (bool, error) {

	res, err := i.engine.Exists(index, id)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

func (i *IndexerES8) CreateDoc(index string, id string, body io.Reader) error {

	res, err := i.engine.Create(index, id, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (i *IndexerES8) GetDoc(index string, id string) ([]byte, error) {

	res, err := i.engine.Get(index, id)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (i *IndexerES8) UpdateDoc(index string, id string, body io.Reader) error {

	res, err := i.engine.Update(index, id, body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}

func (i *IndexerES8) DeleteDoc(index string, id string) error {

	res, err := i.engine.Delete(index, id)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", string(body))
	}

	return nil
}
