package flag

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
