package core

import (
	session "codeberg.org/vnpower/pixivfe/v2/core/config"
	http "codeberg.org/vnpower/pixivfe/v2/core/http"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type ArtworkBrief struct {
	ID           string `json:"id"`
	Title        string `json:"title"`
	ArtistID     string `json:"userId"`
	ArtistName   string `json:"userName"`
	ArtistAvatar string `json:"profileImageUrl"`
	Thumbnail    string `json:"url"`
	Pages        int    `json:"pageCount"`
	XRestrict    int    `json:"xRestrict"`
	AiType       int    `json:"aiType"`
	Bookmarked   any    `json:"bookmarkData"`
}

func GetNewestArtworks(c *fiber.Ctx, worktype string, r18 string) ([]ArtworkBrief, error) {
	token := session.GetToken(c)
	URL := http.GetNewestArtworksURL(worktype, r18, "0")

	var body struct {
		Artworks []ArtworkBrief `json:"illusts"`
		// LastId string
	}

	resp, err := http.UnwrapWebAPIRequest(URL, token)
	if err != nil {
		return nil, err
	}
	resp = session.ProxyImageUrl(resp)

	err = json.Unmarshal([]byte(resp), &body)
	if err != nil {
		return nil, err
	}

	return body.Artworks, nil
}
