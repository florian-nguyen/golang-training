package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"gitlab.com/florian-nguyen/training/podcast-graphql/feeds"
	"gitlab.com/florian-nguyen/training/podcast-graphql/graph/generated"
	"gitlab.com/florian-nguyen/training/podcast-graphql/graph/model"
	"gitlab.com/florian-nguyen/training/podcast-graphql/graph/model/utils"
	"gitlab.com/florian-nguyen/training/podcast-graphql/itunes"
)

func (r *queryResolver) Search(ctx context.Context, term string) ([]*model.Podcast, error) {
	ias := itunes.NewItunesApiServices()

	res, err := ias.Search(term)
	if err != nil {
		return nil, err
	}

	var podcasts []*model.Podcast

	for _, res := range res.Results {
		podcast := &model.Podcast{
			Artist:       res.Artistname,
			PodcastName:  res.Trackname,
			FeedURL:      res.Feedurl,
			Thumbnail:    res.Artworkurl100,
			EpisodeCount: res.Trackcount,
			Genres:       res.Genres,
		}
		podcasts = append(podcasts, podcast)
	}

	return podcasts, nil
}

func (r *queryResolver) Feed(ctx context.Context, feedURL string) ([]*model.FeedItem, error) {
	res, err := feeds.GetFeed(feedURL)
	if err != nil {
		return nil, err
	}

	var feedItems []*model.FeedItem

	for _, item := range res.Channel.Item {
		feedItem := &model.FeedItem{
			PubDate:     item.PubDate,
			Text:        item.Text,
			Title:       item.Title,
			Subtitle:    item.Subtitle,
			Description: item.Description,
			Image:       utils.CheckNilString(item.Image.Href),
			Summary:     item.Summary,
			LinkURL:     item.Enclosure.URL,
			Duration:    item.Duration,
		}

		feedItems = append(feedItems, feedItem)
	}

	return feedItems, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
