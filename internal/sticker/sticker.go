/*
Kitty image protocol - https://sw.kovidgoyal.net/kitty/graphics-protocol/
*/

package sticker

import (
	b64 "encoding/base64"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// TODO: do not use relative path
var StickersPath = "./pic"

func stringToBase64(content []byte) string {
	return b64.StdEncoding.EncodeToString(content)
}

func readStickerFile(name string) ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(StickersPath, name+".png"))
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}

func buildKittyStiker(name string) string {
	content, err := readStickerFile(name)
	if err != nil {
		return name
	}

	var out string
	for {
		var chunk []byte
		var m string
		chunkSize := 4096

		if len(content) > chunkSize {
			chunk = content[:chunkSize]
			content = content[chunkSize:]
			m = "1"
		} else {
			chunk = content
			content = []byte{}
			m = "0"
		}

		// TODO: delete hardcode
		out = out + "\033_G"
		out = out + fmt.Sprintf("m=%s,a=T,f=100,r=1;", m)
		out = out + stringToBase64(chunk)
		out = out + "\033\\"

		if len(content) == 0 {
			break
		}
	}
	return out
}

func getSupportedNames() []string {
	return []string{
		"Kappa",
		"LUL",
		"PogChamp",
		"VoHiYo",
		"NotLikeThis",
		"<3",
		"BibleThump",
		"WutFace",
		"ResidentSleeper",
		"Kreygasm",
		"SeemsGood",
		"SirPrise",
		"SirSad",
		"SirMad",
		"SirSword",
		"SirShield",
		"Shush",
		"PizzaTime",
		"LaundryBasket",
		"ModLove",
		"PotFriend",
		"Jebasted",
		"PogBones",
		"PoroSad",
		"KEKHeim",
		"CaitlynS",
		"HarleyWink",
		"WhySoSerious",
		"DarkKnight",
		"FamilyMan",
		"RyuChamp",
		"HungryPaimon",
		"TransgenderPride",
		"PansexualPride",
		"NonbinaryPride",
		"LesbianPride",
		"IntersexPride",
		"GenderFluidPride",
		"GayPride",
		"BisexualPride",
		"AsexualPride",
		"NewRecord",
		"PogChamp",
		"GlitchNRG",
		"GlitchLit",
		"StinkyGlitch",
		"GlitchCat",
		"FootGoal",
		"FootYellow",
		"FootBall",
		"BlackLivesMatter",
		"ExtraLife",
		"VirtualHug",
		"BOP",
		"SingsNote",
		"SingsMic",
		"TwitchSings",
		"SoonerLater",
		"HolidayTree",
		"HolidaySanta",
		"HolidayPresent",
		"HolidayLog",
		"HolidayCookie",
		"GunRun",
		"PixelBob",
		"FBPenalty",
		"FBChallenge",
		"FBCatch",
		"FBBlock",
		"FBSpiral",
		"FBPass",
		"FBRun",
		"MaxLOL",
		"TwitchRPG",
		"PinkMercy",
		"MercyWing2",
		"MercyWing1",
		"PartyHat",
		"EarthDay",
		"TombRaid",
		"PopCorn",
		"FBtouchdown",
		"TPFufun",
		"TwitchVotes",
		"DarkMode",
		"HSWP",
		"HSCheers",
		"PowerUpL",
		"PowerUpR",
		"LUL",
		"EntropyWins",
		"TPcrunchyroll",
		"TwitchUnity",
		"Squid4",
		"Squid3",
		"Squid2",
		"Squid1",
		"CrreamAwk",
		"CarlSmile",
		"TwitchLit",
		"TehePelo",
		"TearGlove",
		"SabaPing",
		"PunOko",
		"KonCha",
		"Kappu",
		"InuyoFace",
		"BigPhish",
		"BegWan",
		"ThankEgg",
		"MorphinTime",
		"TheIlluminati",
		"TBAngel",
		"MVGame",
		"NinjaGrumpy",
		"PartyTime",
		"RlyTho",
		"UWot",
		"YouDontSay",
		"KAPOW",
		"ItsBoshyTime",
		"CoolStoryBob",
		"TriHard",
		"SuperVinlin",
		"FreakinStinkin",
		"Poooound",
		"CurseLit",
		"BatChest",
		"BrainSlug",
		"PrimeMe",
		"StrawBeary",
		"RaccAttack",
		"UncleNox",
		"WTRuck",
		"TooSpicy",
		"Jebaited",
		"DogFace",
		"BlargNaut",
		"TakeNRG",
		"GivePLZ",
		"imGlitch",
		"pastaThat",
		"copyThis",
		"UnSane",
		"DatSheffy",
		"TheTarFu",
		"PicoMause",
		"TinyFace",
		"DxCat",
		"RuleFive",
		"VoteNay",
		"VoteYea",
		"PJSugar",
		"DoritosChip",
		"OpieOP",
		"FutureMan",
		"ChefFrank",
		"StinkyCheese",
		"NomNom",
		"SmoocherZ",
		"cmonBruh",
		"KappaWealth",
		"MikeHogu",
		"VoHiYo",
		"KomodoHype",
		"SeriousSloth",
		"OSFrog",
		"OhMyDog",
		"KappaClaus",
		"KappaRoss",
		"MingLee",
		"SeemsGood",
		"twitchRaid",
		"bleedPurple",
		"duDudu",
		"riPepperonis",
		"NotLikeThis",
		"DendiFace",
		"CoolCat",
		"KappaPride",
		"ShadyLulu",
		"ArgieB8",
		"CorgiDerp",
		"PraiseIt",
		"TTours",
		"mcaT",
		"NotATK",
		"HeyGuys",
		"Mau5",
		"PRChase",
		"WutFace",
		"BuddhaBar",
		"PermaSmug",
		"panicBasket",
		"BabyRage",
		"HassaanChop",
		"TheThing",
		"EleGiggle",
		"RitzMitz",
		"YouWHY",
		"PipeHype",
		"BrokeBack",
		"ANELE",
		"PanicVis",
		"GrammarKing",
		"PeoplesChamp",
		"SoBayed",
		"BigBrother",
		"Keepo",
		"Kippa",
		"RalpherZ",
		"TF2John",
		"ThunBeast",
		"WholeWheat",
		"DAESuppy",
		"FailFish",
		"HotPokket",
		"4Head",
		"ResidentSleeper",
		"FUNgineer",
		"PMSTwin",
		"ShazBotstix",
		"BibleThump",
		"AsianGlow",
		"DBstyle",
		"BloodTrail",
		"OneHand",
		"FrankerZ",
		"SMOrc",
		"ArsonNoSexy",
		"PunchTrees",
		"SSSsss",
		"Kreygasm",
		"KevinTurtle",
		"PJSalt",
		"SwiftRage",
		"DansGame",
		"GingerPower",
		"BCWarrior",
		"MrDestructoid",
		"JonCarnage",
		"Kappa",
		"RedCoat",
		"TheRinger",
		"StoneLightning",
		"OptimizePrime",
		"JKanStyle",
		"R)",
		";p",
		":p",
		";)",
		":\\",
		"<3",
		":O",
		"B)",
		"O_o",
		":|",
		">(",
		":D",
		":(",
		":)",
	}
}

func FindAndReplace(text string) string {
	if !isKitty() {
		return text
	}

	stickers := getSupportedNames()
	for _, name := range stickers {
		if !strings.Contains(text, name) {
			continue
		}
		buildedSticker := buildKittyStiker(name)
		text = strings.ReplaceAll(text, name, buildedSticker)
	}
	return text
}

func isKitty() bool {
	term := os.Getenv("TERM")
	return term == "xterm-kitty"
}
