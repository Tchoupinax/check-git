<p align="center">
  <strong>check-git</strong><br>
  <em>Your folder of repos, one glance.</em>
</p>

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

```bash
brew tap tchoupinax/brew
brew install tchoupinax/brew/ckg
```

Then run:

```bash
ckg Code
```

## Examples

```bash
ckg Documents
ckg ~/work/oss details
ckg Projects pull
```

## Why ckg

- **Fast** — only repos that contain `.git` are checked
- **Readable** — green / red / yellow output, optional change lines in `details`
- **Parallel pulls** — one command instead of a shell loop
- **Small** — single Go binary, no runtime to install

## Development

```bash
go get ./...
go install github.com/onsi/ginkgo/v2/ginkgo@latest
ginkgo -r ./...
just lint
```

---

<p align="center">
  Made for people who own more than one git clone.<br>
  <a href="https://github.com/Tchoupinax/check-git/issues">Ideas & issues welcome</a>
</p>
