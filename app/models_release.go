package app

var releaseTypes = []string{"tv", "anime", "movies"}

func (c *Connector) ReleaseGet(id string) (*Release, error) {
	m, err := c.Release.Get(id, &Release{})
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) ReleaseList(page, limit int) ([]*Release, int64, error) {
	q := c.Release.Query().Limit(limit).Skip((page - 1) * limit).Desc("created_at")

	total, err := q.Count()
	if err != nil {
		return nil, 0, err
	}

	list, err := c.Release.Query().Limit(limit).Skip((page - 1) * limit).Run()
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

func (c *Connector) ReleasesAll() ([]*Release, error) {
	return c.Release.Query().Limit(-1).Run()
}

func (c *Connector) ReleaseSetting(id, setting string, value bool) error {
	release := &Release{}
	err := c.Release.Find(id, release)
	if err != nil {
		return err
	}

	switch setting {
	case "verified":
		release.Verified = value
	}

	return c.Release.Update(release)
}
