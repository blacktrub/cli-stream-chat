package platform

type Platform struct {
	Name string
}

const Twitch string = "TW"
const Youtube string = "YT"

var TwitchPlatform = Platform{Name: Twitch}
var YoutubePlatform = Platform{Name: Youtube}
