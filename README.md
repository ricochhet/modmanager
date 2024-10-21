# modmanager
A CLI based mod manager for Capcom games.

> [!Warning]
> #### Under Development
> ###### There is no stable version.

## Privacy
`modmanager` is an open source project. Your commit credentials as author of a commit will be visible by anyone. Please make sure you understand this before submitting a PR.
Feel free to use a "fake" username and email on your commits by using the following commands:
```bash
git config --local user.name "USERNAME"
git config --local user.email "USERNAME@SOMETHING.com"
```

## Requirements (Building)
- Go 1.23 or later.
- GNU Make

### Requirements (Development)
- All of the above requirements.
- golangci-lint 
    - `go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest` or `v1.61.0`
- gofumpt
    - `go install mvdan.cc/gofumpt@latest`
- Optionally: deadcode
    - `go install golang.org/x/tools/cmd/deadcode@latest`

## Build
1. Run `make modmanager-[platform]` where platform is one of the following:
    - linux, linux-arm, darwin, darwin-arm, windows
    - Optionally: run `make syso` first if you have `windres` in your PATH.
2. Extras are included as supplementary tools for modding and debugging, view the additional `Makefiles` to build them.

## Supported Games (Tested)
- `MONSTER HUNTER WORLD: ICEBORNE`
- `MONSTER HUNTER RISE: SUNBREAK` 
    - By extension all RE Engine games as of 2024 should be supported.

### Information

#### Config

The `modmanager_config.txt` file expects a list of `key=value` pairs. The configuration file overwrites all present CLI flags.

#### Common CLI flags

- `--game [game_name]` (Default: `mhr`)
    - Specify the game to process mod files for. The `game_name` should match the modmanager data folder *and* the mods folder.
    - `--game mhr`, `data/mods/mhr`, `data/modmanager/mhr`
    - `--game mhw`, `data/mods/mhw`, `data/modmanager/mhw`

View [main.go](./main.go) to see a full list of command line flags.

**Do not use command line flags found in [main.go](./main.go) unless you know 100% what you are doing.**

#### Adding Mods

Mods should be placed under `data/mods/[game_name]/`.

#### Adding Games

Each game has multiple files to control how it handles mods:
- `data.json`
    - The `data` file controls how mods are copied and handles special cases.
- `formats.json`
    - The `formats` file controls what file archive formats should be read, by default, this includes 7z, zip, and rar.

Adding a game can be accomplished by copying or or modifying a `data.json` and `formats.json` file and placing them under `data/modmanager/[game_name]/*.json`.

#### Addons
`data/user/[game_name]/addons.json`

The mod manager will attempt to copy the specified source path to the destination path when it finds the named mod.

```
{
    "addons": [{
      "name": "ExampleMod1",
      "source": "AddonDirectory",
      "destination": "AddonDirectory"
    },
    {
      "name": "ExampleMod2",
      "source": "ExampleMod2Option1/AddonDirectory",
      "destination": "AddonDirectory"
    },
    {
      "name": "ExampleMod2",
      "source": "ExampleMod2Option2/AddonDirectory",
      "destination": "AddonDirectory"
    }]
  }
```

#### Exclusions
`data/user/[game_name]/exclusions.json`

When the mod manager comes across a file or directory in the specified mod, it will not copy it over.

```
{
  "exclusions": [{
    "name": "ExampleMod1",
    "path": "ExcludedDirectory"
  },
  {
    "name": "ExampleMod2",
    "path": "ExcludedDirectory/ExcludedFile.txt"
  }]
}
```

#### Load Order
`data/user/[game_name]/loadorder.json`

By default the mod manager will copy files over in alphabetical order, specifying the index overrides this behavior, allowing you to load a mod earlier, or later.

```
{
    "loadOrder": [{
      "name": "ExampleMod1",
      "index": 5
    }, {
      "name": "ExampleMod2",
      "index": -1
    }]
  }
```
**Note: '-1' places the mod at the end of the list.**

#### Renames
`data/user/[game_name]/renames.json`

To resolve naming conflicts while allowing the ability to keep both modifications active, you can specify a part of the path or file name to replace when it gets copied.

```
{
    "renames": [{
      "name": "ExampleMod1",
      "old": "000_0000",
      "new": "001_0000"
    },
    {
      "name": "ExampleMod2",
      "old": "000_0000",
      "new": "002_0000"
    }]
}
```

## Platforms

|        | Windows|Linux (Untested)|Mac OS (Untested)|
|--------|--------|----------------|-----------------|
| x86-64 | ✅ | ❌ | ❌ |
| x86    | ❌ | ❌ | ❌ |
| ARM64  | ❌ | ❌ | ❌ |

## Contribution Guidelines
If you would like to contribute to `modmanager` please take the time to carefully read the guidelines below.

### Commit Workflow
- Run `make lint` and ensure ALL diagnostics are fixed.
- Run `make fmt` to ensure consistent formatting.
- Create concise, descriptive commit messages to summarize your changes.
    - Optionally: use `git cz` with the [Commitizen CLI](https://github.com/commitizen/cz-cli#conventional-commit-messages-as-a-global-utility) to prepare commit messages.
- Provide *at least* one short sentence or paragraph in your commit message body to describe your thought process for the changes being committed.

### Pull Requests (PRs) should only contain one feature or fix.
It is very difficult to review pull requests which touch multiple unrelated features and parts of the codebase.

Please do not submit pull requests like this; you will be asked to separate them into smaller PRs that deal only with one feature or bug fix at a time.

### Codebase refactors must have prior approval.
Refactors to the structure of the codebase are not taken lightly and require prior discussion and approval.

Please do not start refactoring the codebase with the expectation of having your changes integrated until you receive an explicit approval or a request to do so.

Similarly, when implementing features and bug fixes, please stick to the structure of the codebase as much as possible and do not take this as an opportunity to do some "refactoring along the way".

It is extremely difficult to review PRs for features and bug fixes if they are lost in sweeping changes to the structure of the codebase.

## License
See [LICENSE](./LICENSE) file.