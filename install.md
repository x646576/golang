# Install

1. Install [gvm](https://github.com/moovweb/gvm)
2. List all Go versions: `gvm listall`
3. Install go: `gvm install goX.Y.Z --binary`
4. Go versions: `gvm list`
5. Select: `gvm use goX.Y.Z`
6. Env vars: `go env`

## Go Version Manager

```bash
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

### Restart a terminal session

```bash
source /Users/dodo/.gvm/scripts/gvm
```

### Install

```bash
gvm install go1.16 -B
gvm alias create 16 go1.16
gvm use 16 # --default
```

### List installed

```bash
gvm list

gvm gos (installed)

   16
=> go1.16
```
