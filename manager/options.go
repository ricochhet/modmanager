package manager

import (
	"errors"
	"os"
	"path/filepath"

	aflag "github.com/ricochhet/modmanager/flag"
	"github.com/ricochhet/modmanager/rules"
	"github.com/ricochhet/simplefs"
)

var errNoGameDataFound = errors.New("no game data found")

func FindGames(opt aflag.Options) ([]aflag.Game, error) {
	games := []aflag.Game{}

	path := simplefs.GetRelativePath(filepath.Join(opt.Data, aflag.RequiredData))

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
			engine, err := rules.ReadEngine(data)
			if err != nil {
				return nil, err
			}

			games = append(games, aflag.Game{Name: dir.Name(), Engine: engine})
		}
	}

	return games, nil
}

func FindFormats(opt aflag.Options) ([]string, error) {
	formats := []string{}

	path := simplefs.GetRelativePath(filepath.Join(opt.Data, aflag.RequiredData))

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
				format, err := rules.ReadFormats(data)
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

func FindGame(opt aflag.Options) (aflag.Game, error) {
	games, err := FindGames(opt)
	if err != nil {
		return aflag.Game{}, err
	}

	for _, game := range games {
		if game.Name == opt.Game {
			return game, nil
		}
	}

	return aflag.Game{}, nil //nolint:exhaustruct // wontfix
}
