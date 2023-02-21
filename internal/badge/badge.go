/*
Fetch badges
https://badges.twitch.tv/v1/badges/global/display
*/
package badge

import (
	"cli-stream-chat/internal/image"
	"encoding/json"
	"sort"

	"io"
	"net/http"
	"os"
	"path/filepath"
)

type Badge struct {
	Name string
	url  string
}

func (b Badge) path() string {
	return filepath.Join(BadgePath, b.Name)
}

func (b Badge) download() error {
	resp, err := http.Get(b.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(b.path())
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

type BadgeList []Badge

func (b BadgeList) Len() int {
	return len(b)
}

func (b BadgeList) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b BadgeList) Less(i, j int) bool {
	return b[i].Name[0] < b[j].Name[0]
}

type Badges struct {
	data BadgeList
}

func (b *Badges) get() BadgeList {
	if len(b.data) > 0 {
		return b.data
	}
	badges := b.fetch()
	b.data = badges
	return b.fetch()
}

type BadgeResponseItem struct {
	ClickAction string      `json:"click_action"`
	ClickURL    string      `json:"click_url"`
	Description string      `json:"description"`
	ImageURL1x  string      `json:"image_url_1x"`
	ImageURL2x  string      `json:"image_url_2x"`
	ImageURL4x  string      `json:"image_url_4x"`
	LastUpdated interface{} `json:"last_updated"`
	Title       string      `json:"title"`
}

func (b *Badges) fetch() BadgeList {
	resp, err := http.Get("https://badges.twitch.tv/v1/badges/global/display")
	if err != nil {
		return BadgeList{}
	}
	defer resp.Body.Close()

	// TODO: we can use json.RawMessage here to make it prettier
	data := map[string]map[string]map[string]map[string]BadgeResponseItem{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return BadgeList{}
	}

	var badges BadgeList
	for key, value := range data["badge_sets"] {
		versions := value["versions"]
		one, exists := versions["1"]
		if !exists {
			continue
		}
		badges = append(badges, Badge{Name: key, url: one.ImageURL2x})
	}
	sort.Sort(badges)
	return badges
}

var BadgePath = "./pic/badges"
var supportedBadges = Badges{}

func Show(badges map[string]int) string {
	var out string
	for _, badge := range supportedBadges.get() {
		if _, ok := badges[string(badge.Name)]; ok {
			err := badge.download()
			if err != nil {
				// TODO: do something
				continue
			}
			out = out + image.Build(badge.Name, badge.path(), 2)
		}
	}
	return out
}
