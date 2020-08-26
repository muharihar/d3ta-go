package indexer

import (
	"fmt"
	"io"
	"strings"

	"github.com/muharihar/d3ta-go/system/indexer/adapter"
)

// IndexerType represent IndexerType
type IndexerType string

const (
	ElasticSearchIndexer IndexerType = "ElasticSearch"
)

func NewIndexer(indexerType IndexerType, indexerEngine adapter.IIndexerEngine) (*Indexer, error) {
	if indexerEngine == nil {
		return nil, fmt.Errorf("Invalid indexerEngine value")
	}

	idx := Indexer{
		_type:         indexerType,
		indexerEngine: indexerEngine,
	}

	// C4 prefix key
	idx.Context = "defaultContext"
	idx.Container = "defaultContainer"
	idx.Component = "defaultComponent"
	// idx.Code = ...

	return &idx, nil
}

// Indexer type
type Indexer struct {
	_type         IndexerType
	indexerEngine adapter.IIndexerEngine

	Context   string
	Container string
	Component string
}

func (i *Indexer) GetEngine() interface{} {
	return i.indexerEngine.GetEngine()
}

func (i *Indexer) CreateIndex(index string, mapping io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.CreateIndex(_index, mapping)
}

func (i *Indexer) DropIndex(indexs []string) error {
	var _indexs []string
	for _, v := range indexs {
		_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), v)
		_indexs = append(_indexs, _index)
	}
	return i.indexerEngine.DropIndex(_indexs)
}

func (i *Indexer) DocExist(index string, id string) (bool, error) {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.DocExist(_index, id)
}

func (i *Indexer) CreateDoc(index string, id string, body io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.CreateDoc(_index, id, body)
}

func (i *Indexer) GetDoc(index string, id string) ([]byte, error) {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.GetDoc(_index, id)
}

func (i *Indexer) UpdateDoc(index string, id string, body io.Reader) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.UpdateDoc(_index, id, body)
}

func (i *Indexer) DeleteDoc(index string, id string) error {
	_index := fmt.Sprintf("%s~%s~%s~%s", strings.ToLower(i.Context), strings.ToLower(i.Container), strings.ToLower(i.Component), index)
	return i.indexerEngine.DeleteDoc(_index, id)
}
