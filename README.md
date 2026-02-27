# git-pick

A simple tool to pick files to add to staging area in Git

## Installation

### Quick Install (Global)

```bash
# Clone the repository
git clone https://github.com/MarcosSiega/git-pick.git
cd git-pick

# Build and install globally
make install
```

This will install the binary to `/usr/local/bin/git-pick`, making it available system-wide.

### Manual Installation

```bash
# Clone the repository
git clone https://github.com/MarcosSiega/git-pick.git
cd git-pick

# Install dependencies
go mod tidy

# Build the binary
go build -o git-pick main.go

# Move to a directory in your PATH
sudo mv git-pick /usr/local/bin/
sudo chmod +x /usr/local/bin/git-pick
```

### Using as Git Alias

Once installed, you can use it directly as `git pick` (yes, Git automatically recognizes `git-*` binaries!):

```bash
git pick
```

Or, if you prefer a custom alias, add this to your `~/.gitconfig`:

```ini
[alias]
    pick = !git-pick
```

## Usage

Simply run in any git repository:

```bash
git-pick
# or
git pick
```

Then use:
- `↑/↓` or `j/k` to navigate
- `Space` to select/deselect files
- `Enter` to add selected files to staging
- `q` or `Esc` to quit

