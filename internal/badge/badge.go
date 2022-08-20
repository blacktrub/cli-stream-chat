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

type Badges []Badge

func (b Badges) Len() int {
	return len(b)
}

func (b Badges) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Badges) Less(i, j int) bool {
	return b[i].Name[0] < b[j].Name[0]
}

var BadgePath = "./pic/badges"

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

func Show(badges map[string]int) string {
	supported := getSupported()
	var out string
	for _, badge := range supported {
		if _, ok := badges[string(badge.Name)]; ok {
			err := downloadBadge(badge)
			if err != nil {
				// TODO: do something
				continue
			}
			out = out + image.Build(badge.Name, badge.path(), 2)
		}
	}
	return out
}

// TODO: cache me
func getSupported() Badges {
	resp, err := http.Get("https://badges.twitch.tv/v1/badges/global/display")
	if err != nil {
		return Badges{}
	}
	defer resp.Body.Close()
	data := map[string]map[string]map[string]map[string]BadgeResponseItem{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return Badges{}
	}

	var badges Badges
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

func downloadBadge(badge Badge) error {
	resp, err := http.Get(badge.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.Create(badge.path())
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
