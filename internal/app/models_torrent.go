package app

func (c *Connector) TorrentGet(id string) (*Torrent, error) {
	m := &Torrent{}
	err := c.Torrent.Find(id, m)
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) TorrentList(page, limit int) ([]*Torrent, error) {
	skip := (page - 1) * limit
	list, err := c.Torrent.Query().Limit(limit).Skip(skip).Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}
