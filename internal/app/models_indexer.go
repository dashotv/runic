package app

import "github.com/dashotv/fae"

func (c *Connector) IndexerGet(id string) (*Indexer, error) {
	m, err := c.Indexer.Get(id, &Indexer{})
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) IndexerByName(name string) (*Indexer, error) {
	list, err := c.Indexer.Query().Where("name", name).Run()
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, fae.Errorf("indexer %s not found", name)
	}
	if len(list) > 1 {
		return nil, fae.Errorf("multiple indexers with name %s", name)
	}

	return list[0], nil
}

func (c *Connector) IndexerList(page, limit int) ([]*Indexer, int64, error) {
	q := c.Indexer.Query().Limit(limit).Skip((page - 1) * limit).Desc("active").Asc("name")

	count, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	indexers, err := q.Run()
	if err != nil {
		return nil, 0, err
	}

	indexers = append([]*Indexer{{Name: "rift", URL: app.Config.RiftURL, Active: true, Categories: []int{5070}}}, indexers...)
	return indexers, count, nil
}

func (c *Connector) IndexerActive() ([]*Indexer, error) {
	indexers, err := c.Indexer.Query().Where("active", true).Desc("created_at").Limit(-1).Run()
	if err != nil {
		return nil, err
	}

	return indexers, nil
}
