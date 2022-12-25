package url_shortener

import (
	"context"

	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	urlShortener "github.com/gurbaj5124871/url-shortener/src/proto/url_shortener"
)

type UrlShortenerGrpcHandler struct {
	Ignite ignite.Client
}

func (grpc *UrlShortenerGrpcHandler) CreateShortURL(ctx context.Context, _ *urlShortener.CreateShortURLRequest) (*urlShortener.CreateShortURLResponse, error) {
	return &urlShortener.CreateShortURLResponse{
		URL: &urlShortener.ShortUrlDTO{
			DestinationURL: "",
			ShortURL:       "",
		},
	}, nil
}
