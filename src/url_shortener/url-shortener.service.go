package url_shortener

import (
	"time"

	ignite "github.com/amsokol/ignite-go-client/binary/v1"
	"github.com/gurbaj5124871/url-shortener/src/db"
	"github.com/gurbaj5124871/url-shortener/src/utils"
)

const (
	ShortURLSTable string = "SHORTURLS"
	ShortURLSCache string = "SQL_PUBLIC_SHORTURLS"
)

type ShortURLS struct {
	ShortString    string    `json:"shortString"`
	DestinationURL string    `json:"destinationURL"`
	CreatedAt      time.Time `json:"createdAt"`
}

func GenerateShortURL(url string) (*string, error) {
	id, err := utils.Snowflake.NextID()
	if err != nil {
		return nil, err
	}

	base62EncodedStr, err := utils.Base62Encode(int64(id))
	return &base62EncodedStr, err
}

func CreateShortURL(shortString string, destinationURL string) (*ShortURLS, error) {
	url := ShortURLS{
		ShortString:    shortString,
		DestinationURL: destinationURL,
		CreatedAt:      time.Now(),
	}

	sql := `INSERT INTO ` + ShortURLSTable + ` (SHORTSTRING, DESTINATIONURL, CREATEDAT) VALUES (?, ?, ?);`
	_, err := db.GetIgniteDB().QuerySQLFields(ShortURLSCache, false, ignite.QuerySQLFieldsData{
		PageSize: 10,
		Query:    sql,
		QueryArgs: []interface{}{
			url.ShortString,
			url.DestinationURL,
			url.CreatedAt.Format(time.RFC3339),
		},
	})
	if err != nil {
		return nil, err
	}

	return &url, nil
}
