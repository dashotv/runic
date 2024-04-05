package app

func (c *Connector) IndexerGet(id string) (*Indexer, error) {
	m, err := c.Indexer.Get(id, &Indexer{})
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) IndexerList(page, limit int) ([]*Indexer, int64, error) {
	q := c.Indexer.Query().Limit(limit).Skip((page - 1) * limit).Desc("processed_at").Asc("name")

	count, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	indexers, err := q.Run()
	if err != nil {
		return nil, 0, err
	}

	indexers = append(indexers, &Indexer{Name: "rift", Url: "https://rift.com", Active: true, Categories: []int{5070}})
	return indexers, count, nil
}

func (c *Connector) IndexerActive() ([]*Indexer, error) {
	indexers, err := c.Indexer.Query().Where("active", true).Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return indexers, nil
}
