package url_shortener

import (
	"context"
	"errors"

	urlShortener "github.com/gurbaj5124871/url-shortener/src/proto/url_shortener"
	"github.com/rs/zerolog/log"
)

type UrlShortenerGrpcHandler struct {
	ShortURLDomain string
}

func (grpc *UrlShortenerGrpcHandler) CreateShortURL(ctx context.Context, req *urlShortener.CreateShortURLRequest) (*urlShortener.CreateShortURLResponse, error) {
	destinationURL := &req.DestinationURL

	shortStr, err := GenerateShortURL(*destinationURL)
	if err != nil {
		log.Err(err).Msgf("Failed to generate short url for %s", *destinationURL)
		return nil, errors.New("failed to generate short url, try again sometime later")
	}

	log.Info().Str("shortStr", *shortStr).Msg("")

	shortURL, err := CreateShortURL(*shortStr, *destinationURL)
	if err != nil {
		log.Err(err).Msgf("Failed to generate short url for %s", *destinationURL)
		return nil, errors.New("failed to generate short url, try again sometime later")
	}

	return &urlShortener.CreateShortURLResponse{
		URL: &urlShortener.ShortUrlDTO{
			DestinationURL: shortURL.DestinationURL,
			ShortURL:       grpc.ShortURLDomain + shortURL.ShortString,
			CreatedAt:      shortURL.CreatedAt.String(),
		},
	}, nil
}
