# check-git

**One command. Every repo. Green or red.**

`ckg` scans a folder of git projects and tells you which ones are clean, which ones still have dirt on them, and optionally pulls them all without opening each directory by hand.

```
ckg ~/code
```

You get a wall of green checkmarks for repos with nothing to commit, and yellow names where `git status` would complain. No clicking through twenty tabs.

## Why bother

| You want | `ckg` does |
|----------|------------|
| Health check across a monorepo folder | `ckg Projects` |
| See what changed | `ckg Projects details` |
| Update everything | `ckg Projects pull` |

## Install

Homebrew:

```bash
brew tap tchoupinax/brew
brew install tchoupinax/brew/ckg
```

## Usage

```bash
ckg <folder>              # clean / dirty per repo
ckg <folder> details      # porcelain lines for dirty repos
ckg <folder> pull         # git pull in parallel
```

`<folder>` can be a path, or a name searched under your home directory (up to a few levels deep).

Examples:

```bash
ckg Documents
ckg ~/work/oss details
ckg Code pull
```

## Development

```bash
go get ./...
go install github.com/onsi/ginkgo/v2/ginkgo@latest
ginkgo -r ./...
just lint   # golangci-lint
```

Issues and ideas: https://github.com/Tchoupinax/check-git/issues
