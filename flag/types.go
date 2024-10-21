package flag

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
