package url_shortener

import (
	"context"
	"errors"
	"net/url"
	"strings"

	urlShortener "github.com/gurbaj5124871/url-shortener/src/proto/url_shortener"
	"github.com/rs/zerolog/log"
)

type UrlShortenerGrpcHandler struct {
	ShortURLDomain string
}

func (grpc *UrlShortenerGrpcHandler) CreateShortURL(
	ctx context.Context,
	req *urlShortener.CreateShortURLRequest,
) (*urlShortener.CreateShortURLResponse, error) {
	destinationURL := &req.DestinationURL

	shortStr, err := GenerateShortURL()
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

func (grpc *UrlShortenerGrpcHandler) GetDestinationURLFromShortURL(
	ctx context.Context,
	req *urlShortener.GetDestinationURLFromShortURLRequest,
) (*urlShortener.GetDestinationURLFromShortURLResponse, error) {
	shortURL := &req.ShortURL

	parsedURL, err := url.Parse(*shortURL)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to get destination url for short url: %s", shortURL)
		return nil, errors.New("invalid short url")
	}

	if serverDomainURL, err := url.Parse(grpc.ShortURLDomain); err != nil || parsedURL.Host != serverDomainURL.Host {
		log.Error().Err(err).Msg("Request short url host mismatch")
		return nil, errors.New("invalid short url")
	}

	paths := strings.Split(parsedURL.Path, "/")
	if len(paths) != 2 {
		log.Err(err).Msgf("Failed to get destination url for short url: %s", shortURL)
		return nil, errors.New("invalid short url")
	}
	shortString := paths[1]

	destinationURL, err := GetDestinationURLFromShortString(shortString)
	if err != nil {
		log.Err(err).Msgf("Failed to get destination url for short url: %s", shortURL)
		return nil, err
	}

	return &urlShortener.GetDestinationURLFromShortURLResponse{
		DestinationURL: *destinationURL,
	}, nil
}
