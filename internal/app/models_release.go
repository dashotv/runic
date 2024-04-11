package app

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/dashotv/fae"
)

var releaseTypes = []string{"tv", "anime", "movies"}

func (r *Release) Created(ctx context.Context) error {
	return app.Events.Send("runic.releases", r)
}
func (r *Release) Updated(ctx context.Context, result *mongo.UpdateResult) error {
	return app.Events.Send("runic.releases", r)
}

func (c *Connector) ReleaseGet(id string) (*Release, error) {
	m, err := c.Release.Get(id, &Release{})
	if err != nil {
		return nil, err
	}

	// post process here

	return m, nil
}

func (c *Connector) ReleaseList(page, limit int) ([]*Release, int64, error) {
	total, err := c.Release.Query().Count()
	if err != nil {
		return nil, 0, err
	}

	list, err := c.Release.Query().Limit(limit).Skip((page - 1) * limit).Desc("published_at").Desc("created_at").Run()
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

var popularIntervals = map[string]int{
	"daily":   1,
	"weekly":  7,
	"monthly": 30,
}

func (c *Connector) ReleasesPopular(interval string) (map[string][]*Popular, error) {
	limit := 25
	out := map[string][]*Popular{}

	i, ok := popularIntervals[interval]
	if !ok {
		return nil, fae.Errorf("invalid interval: %s", interval)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	date := time.Now().AddDate(0, 0, -i)
	for _, t := range []string{"tv", "anime", "movies"} {
		results, err := c.ReleasesPopularType(ctx, t, date, limit)
		if err != nil {
			return nil, fae.Wrap(err, "popular releases")
		}
		out[t] = results
	}

	return out, nil
}

func (c *Connector) ReleasesPopularType(ctx context.Context, t string, date time.Time, limit int) ([]*Popular, error) {
	p := []bson.M{
		{"$project": bson.M{"title": 1, "type": 1, "year": 1, "published": "$published_at"}},
		{"$match": bson.M{"title": bson.M{"$nin": bson.A{"", nil}}, "type": t, "published": bson.M{"$gte": date}}},
		{"$group": bson.M{"_id": "$title", "type": bson.M{"$first": "$type"}, "year": bson.M{"$first": "$year"}, "count": bson.M{"$sum": 1}}},
		{"$sort": bson.M{"count": -1}},
		{"$limit": limit},
	}

	cursor, err := c.Release.Collection.Aggregate(ctx, p)
	if err != nil {
		return nil, fae.Wrap(err, "aggregating popular releases")
	}

	results := make([]*Popular, limit)
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fae.Wrap(err, "decoding popular releases")
	}

	return results, nil
}