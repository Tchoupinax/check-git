<p align="center">
  <strong>check-git</strong><br>
  <em>Your folder of repos, one glance.</em>
</p>

<<<<<<< HEAD
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
=======
<p align="center">
  Stop opening twenty terminals.<br>
  <code>ckg</code> walks a directory, finds every <code>.git</code>, and shows clean vs dirty in color — then <strong>details</strong> or <strong>pull</strong> when you are ready.
</p>

<p align="center">
  <a href="https://github.com/Tchoupinax/check-git/issues">Report a bug</a>
  ·
  <a href="#install">Install</a>
  ·
  <a href="#usage">Usage</a>
</p>

---

## At a glance

```bash
ckg ~/code
```

```
  clean-repo      everything committed
  dirty-repo      uncommitted changes (yellow)
```

One folder. Every nested repo. No scripts to maintain.

## What you get

| Mode | Command | Best for |
|------|---------|----------|
| **Scan** | `ckg <folder>` | Morning sanity check before standup |
| **Details** | `ckg <folder> details` | See porcelain lines without `cd` |
| **Pull all** | `ckg <folder> pull` | Refresh every repo in parallel |

`<folder>` can be a full path **or** a folder name — `ckg` searches a few levels under your home directory if the path is not absolute.

## Install

**macOS (Homebrew)**
>>>>>>> 07671e1 (Polish README for onboarding.)

```bash
brew tap tchoupinax/brew
brew install tchoupinax/brew/ckg
```

<<<<<<< HEAD
## Usage

```bash
ckg <folder>              # clean / dirty per repo
ckg <folder> details      # porcelain lines for dirty repos
ckg <folder> pull         # git pull in parallel
```

`<folder>` can be a path, or a name searched under your home directory (up to a few levels deep).

Examples:
=======
Then run:

```bash
ckg Code
```

## Examples
>>>>>>> 07671e1 (Polish README for onboarding.)

```bash
ckg Documents
ckg ~/work/oss details
<<<<<<< HEAD
ckg Code pull
```
=======
ckg Projects pull
```

## Why ckg

- **Fast** — only repos that contain `.git` are checked
- **Readable** — green / red / yellow output, optional change lines in `details`
- **Parallel pulls** — one command instead of a shell loop
- **Small** — single Go binary, no runtime to install
>>>>>>> 07671e1 (Polish README for onboarding.)

## Development

```bash
go get ./...
go install github.com/onsi/ginkgo/v2/ginkgo@latest
ginkgo -r ./...
<<<<<<< HEAD
just lint   # golangci-lint
```

Issues and ideas: https://github.com/Tchoupinax/check-git/issues
=======
just lint
```

---

<p align="center">
  Made for people who own more than one git clone.<br>
  <a href="https://github.com/Tchoupinax/check-git/issues">Ideas & issues welcome</a>
</p>
>>>>>>> 07671e1 (Polish README for onboarding.)
