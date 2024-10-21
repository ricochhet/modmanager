package internal

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ricochhet/simplefs"
)

type Game struct {
	Name   string `json:"name"`
	Engine Engine `json:"engine"`
}

type Engine struct {
	Paths []Data `json:"paths"`
	Hooks []Hook `json:"hooks"`
}

type Files struct {
	Files []string `json:"files"`
	Data  Data     `json:"data"`
}

type Data struct {
	Path        string   `json:"path"`
	IsDir       bool     `json:"isDir"`
	Requires    []string `json:"requires"`
	Unsupported bool     `json:"unsupported"`
}

type Hook struct {
	Name     string   `json:"name"`
	Dll      string   `json:"dll"`
	Arch     string   `json:"arch"`
	Requires []string `json:"requires"`
	Include  []string `json:"include"`
}

type Options struct {
	// user
	Game       string `json:"game"`
	LoadOrder  string `json:"loadOrder"`
	Addons     string `json:"addons"`
	Renames    string `json:"renames"`
	Exclusions string `json:"exclusions"`
	Engine     string `json:"engine"`
	Formats    string `json:"formats"`
	Bin        string `json:"bin"`
	Silent     bool   `json:"silent"`

	// manager
	Data   string `json:"data"`
	Mods   string `json:"mods"`
	Temp   string `json:"temp"`
	Output string `json:"output"`
	User   string `json:"user"`
	Hook   string `json:"hook"`
	Config string `json:"config"`
	Log    string `json:"log"`
}

const requiredData = "modmanager"

var errNoGameDataFound = errors.New("no game data found")

func NewOptions() *Options {
	return &Options{
		Game:       "mhr",
		LoadOrder:  "loadOrder.json",
		Addons:     "addons.json",
		Renames:    "renames.json",
		Exclusions: "exclusions.json",
		Engine:     "data.json",
		Formats:    "formats.json",
		Bin:        "redist/win64/7z.exe",
		Silent:     false,
		Data:       "data",
		Mods:       "mods",
		Temp:       "temp",
		Output:     "output",
		User:       "User",
		Hook:       ".dll",
		Config:     "modmanager_config.txt",
		Log:        "modmanager_log.txt",
	}
}

func ConfigPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Config)
}

func SetFlagString(flag *string, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	*flag = val
}

func SetFlagBool(flag *bool, key string, kvp map[string]string) {
	val := kvp[key]

	if val == "" {
		return
	}

	b, parseErr := strconv.ParseBool(val)

	*flag = b

	if parseErr != nil {
		panic(parseErr)
	}
}

func MapConfig(lines []string) (map[string]string, error) {
	keyvalues := map[string]string{}

	for _, line := range lines {
		keyvalue := strings.SplitN(line, "=", 2)

		if len(keyvalue) != 2 {
			continue // skip
		}

		keyvalues[keyvalue[0]] = keyvalue[1]
	}

	return keyvalues, nil
}

func LogPath(opt Options) string {
	return simplefs.GetRelativePath(filepath.Join(opt.Data, requiredData, opt.Log))
}

func ModPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Mods, opt.Game)
}

func TempPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Temp, opt.Game)
}

func OutputPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.Output, opt.Game)
}

func LoadOrderPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.LoadOrder)
}

func AddonPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Addons)
}

func RenamePath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Renames)
}

func ExclusionPath(opt Options) string {
	return simplefs.GetRelativePath(opt.Data, opt.User, opt.Game, opt.Exclusions)
}

func FindGames(opt Options) ([]Game, error) {
	games := []Game{}

	path := simplefs.GetRelativePath(filepath.Join(opt.Data, requiredData))

	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(dirs) == 0 {
		return nil, errNoGameDataFound
	}

	for _, dir := range dirs {
		data := filepath.Join(path, dir.Name(), opt.Engine)

		if simplefs.Exists(data) {
			engine, err := ReadEngine(data)
			if err != nil {
				return nil, err
			}

			games = append(games, Game{Name: dir.Name(), Engine: engine})
		}
	}

	return games, nil
}

func FindFormats(opt Options) ([]string, error) {
	formats := []string{}

	path := simplefs.GetRelativePath(filepath.Join(opt.Data, requiredData))

	dirs, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	if len(dirs) == 0 {
		return nil, errNoGameDataFound
	}

	for _, dir := range dirs {
		if dir.IsDir() && opt.Game == dir.Name() {
			data := filepath.Join(path, dir.Name(), opt.Formats)

			if simplefs.Exists(data) {
				format, err := ReadFormats(data)
				if err != nil {
					return nil, err
				}

				formats = format.JSON

				break
			}
		}
	}

	return formats, nil
}

func FindGame(opt Options) (Game, error) {
	games, err := FindGames(opt)
	if err != nil {
		return Game{}, err
	}

	for _, game := range games {
		if game.Name == opt.Game {
			return game, nil
		}
	}

	return Game{}, nil //nolint:exhaustruct // wontfix
}
