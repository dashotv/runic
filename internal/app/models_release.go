package app

import (
	"cmp"
	"context"
	"slices"
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

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
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
		{"$match": bson.M{"title": bson.M{"$ne": ""}, "type": t, "published_at": bson.M{"$gte": date}}},
		{"$project": bson.M{"title": 1, "type": 1, "year": 1}},
		{"$group": bson.M{"_id": "$title", "title": bson.M{"$first": "$title"}, "type": bson.M{"$first": "$type"}, "year": bson.M{"$first": "$year"}, "count": bson.M{"$sum": 1}}},
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

	// HACK: sort by count and title in aggregate sometimes fucks up and sorts by title only
	slices.SortFunc(results, func(a, b *Popular) int {
		if n := cmp.Compare(b.Count, a.Count); n != 0 {
			return n
		}
		// If counts are equal, order by title
		return cmp.Compare(a.Title, b.Title)
	})

	return results, nil
}

func (c *Connector) ReleasesPopularMovies() ([]*PopularMovie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	defer TickTock("ReleasesPopularMovies")()

	date := time.Now().AddDate(0, 0, -30)
	limit := 25
	t := "movies"
	p := []bson.M{
		{"$match": bson.M{"type": t, "published_at": bson.M{"$gte": date}, "resolution": "1080"}},
		{"$project": bson.M{"title": 1, "type": 1, "year": 1, "published": "$published_at", "verified": 1}},
		{"$group": bson.M{"_id": bson.M{"title": "$title", "year": "$year"}, "count": bson.M{"$sum": 1}, "verified": bson.M{"$sum": bson.M{"$cond": bson.M{"if": "$verified", "then": 1, "else": 0}}}}},
		{"$sort": bson.M{"verified": -1, "count": -1}},
		{"$limit": limit},
	}

	cursor, err := c.Release.Collection.Aggregate(ctx, p)
	if err != nil {
		return nil, fae.Wrap(err, "aggregating popular releases")
	}

	results := make([]*PopularMovie, limit)
	if err = cursor.All(ctx, &results); err != nil {
		return nil, fae.Wrap(err, "decoding popular releases")
	}

	return results, nil
}
