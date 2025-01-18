# Asset Scanner

A command-line tool that scans for unused assets in your project and optionally removes them.

## Description

Asset Scanner helps you maintain a clean codebase by identifying assets (like images, fonts, etc.) that are no longer referenced in your target directories. It can either report unused assets or automatically remove them based on your preference.

## Installation

Requires Go 1.23.0 or higher.

bash
go install github.com/pawelataman/asset-scanner@latest

## Usage

Basic command structure:

bash
asset-scanner [assetsDirectory] [targetDirectory] [flags]

### Parameters

- `assetsDirectory`: Root directory containing the assets to be checked
- `targetDirectory`: Root directory where the tool should search for asset references

### Flags

- `-ext`: Specify the asset file extension to scan for (e.g., ".png", ".jpg")
- `-remove`: When set, removes unused assets instead of just reporting them

### Examples

#### Report unused PNG files:

```bash
asset-scanner ./assets ./src -ext=.png
```

#### Remove unused PNG files:

```bash
asset-scanner ./assets ./src -ext=.png -remove
```

## How It Works

1. The scanner first creates an inventory of all assets in the specified assets directory that match the given file extension.
2. It then recursively searches through all files in the target directory for references to these assets.
3. Finally, it either reports or removes assets that have zero occurrences in the target directory.

## Features

- Recursive directory scanning
- File extension filtering
- Optional automatic removal of unused assets
- Clear reporting of unused assets
- Safe scanning (read-only) by default

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.