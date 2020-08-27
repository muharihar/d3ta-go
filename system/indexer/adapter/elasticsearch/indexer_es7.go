package elasticsearch

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	es7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/muharihar/d3ta-go/system/indexer/adapter"
)

// NewIndexerES7 new Elastic Search 7 Indexer
func NewIndexerES7(cfg es7.Config) (adapter.IIndexerEngine, error) {
	var err error

	idx := &IndexerES7{
		esVersion: ESVersion7,
	}
	idx.ctx = context.Background()
	if idx.engine, err = es7.NewClient(cfg); err != nil {
		return nil, err
	}

	return idx, nil
}

// IndexerES7 type
type IndexerES7 struct {
	ctx       context.Context
	esVersion ESVersion
	engine    *es7.Client
}

// GetEngine get Indexer Engine
func (i *IndexerES7) GetEngine() interface{} {
	return i.engine
}

func (i *IndexerES7) Search(query io.Reader, prettify bool) ([]byte, error) {
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

func (i *IndexerES7) SearchIndexDoc(index string, query io.Reader, size int, prettify bool) ([]byte, error) {
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

func (i *IndexerES7) IndexExist(indexs []string) (bool, error) {
	res, err := i.engine.Indices.Exists(indexs)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

func (i *IndexerES7) CreateIndex(index string, mapping io.Reader) error {
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

func (i *IndexerES7) DropIndex(indexs []string) error {
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

func (i *IndexerES7) DocExist(index string, id string) (bool, error) {

	res, err := i.engine.Exists(index, id)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	return res.StatusCode == 200, nil
}

func (i *IndexerES7) CreateDoc(index string, id string, body io.Reader) error {

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

func (i *IndexerES7) GetDoc(index string, id string) ([]byte, error) {

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

func (i *IndexerES7) UpdateDoc(index string, id string, body io.Reader) error {

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

func (i *IndexerES7) DeleteDoc(index string, id string) error {

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
