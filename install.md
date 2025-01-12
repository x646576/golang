# Install

1. Install [gvm](https://github.com/moovweb/gvm)
2. List all Go versions: `gvm listall`
3. Install go: `gvm install goX.Y.Z --binary`
4. Go versions: `gvm list`
5. Select: `gvm use goX.Y.Z`
6. Env vars: `go env`
7. gopls: `go install golang.org/x/tools/gopls@latest`
8. Editor extension

## Go Version Manager

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

### Restart a terminal session

`~/.zshrc`:

```bash
# Languages
## Go
source /Users/$USER/.gvm/scripts/gvm
```

### Install

```bash
gvm install go1.19.4 -B
gvm use go1.19.4 --default
```

### List installed

```bash
gvm list

gvm gos (installed)

=> go1.19.4
```

---

## Visual Studio Code Extension

Extension: [Go](https://marketplace.visualstudio.com/items?itemName=golang.go)

```bash
go env

GOPATH="~/.gvm/pkgsets/go1.19.4/global"
GOROOT="~/.gvm/gos/go1.19.4"
```

**Preferences: Configure Language Specific Settings** `⇧⌘P` → Go

`settings.json`

```json
  "go.gopath": "~/.gvm/pkgsets/go1.19.4/global",
  "go.goroot": "~/.gvm/gos/go1.19.4",
  "[go]": {
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    },
    "editor.suggest.snippetsPreventQuickSuggestions": false,
    "editor.defaultFormatter": "golang.go"
  },
  "gopls": {
    "experimentalWorkspaceModule": true
  }
```

### gopls

VS Code should handle that step for you.

```bash
Tools environment: GOPATH=~/.gvm/pkgsets/go1.19.4/global
Installing 1 tool at ~/.gvm/pkgsets/go1.19.4/global/bin in module mode.
  gopls
  gopkgs
  go-outline
  dlv
  staticcheck

Installing golang.org/x/tools/gopls (~/.gvm/pkgsets/go1.19.4/global/bin/gopls) SUCCEEDED
# ...

All tools successfully installed. You are ready to Go :).
```

---

## SpaceVim

- [Use Vim as a Go IDE](https://spacevim.org/use-vim-as-a-go-ide/)

### Configuration

```toml
[[layers]]
  name = "format"

[[layers]]
  name = "lang#go"
```

- go run: `SPC l r`
- go build: `SPC l b`
- go test: `SPC l t`
- code coverage: `SPC l c`
- gofmt:`SPC b f`

### Install packages

in `nvim`

```bash
:GoInstallBinaries
```
