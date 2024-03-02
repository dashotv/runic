package app

func (c *Connector) MinionGet(id string) (*Minion, error) {
	m, err := c.Minion.Get(id, &Minion{})
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) MinionList(limit, skip int) ([]*Minion, error) {
	if limit == 0 {
		limit = 25
	}
	list, err := c.Minion.Query().Limit(limit).Skip(skip).Desc("created_at").Run()
	if err != nil {
		return nil, err
	}

	return list, nil
}
