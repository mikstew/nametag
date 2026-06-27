# nametag

A small desktop app that shows a customizable nametag in a window.

## Project layout

```
cmd/nametag/          Application entrypoint
internal/
  app/                Window lifecycle and wiring
  config/             Hardcoded nametag settings and version
  platform/           OS-specific helpers (restart)
  ui/nametag/         Nametag view
  update/             GitHub Releases self-update
```

## Requirements

- Go 1.22+
- macOS: Xcode command line tools (for native window support)

## Run

```bash
go run ./cmd/nametag
```

## Build

```bash
go build \
  -ldflags="-s -w -extldflags=-Wl,-no_warn_duplicate_libraries" \
  -o nametag \
  ./cmd/nametag
```

## Customize

Edit `DisplayName` and `TagColor` in `internal/config/config.go`, then rebuild.

## Releases and self-update

The refresh button on the nametag checks [GitHub Releases](https://github.com/mikio/nametag/releases) for a newer version. If one exists, it downloads the binary for your OS/arch and restarts.

Release assets must be named like `nametag-darwin-arm64` (also supports `_` instead of `-`).

### Publish a release (recommended)

Fyne uses CGO, so each OS must be built on a native runner. Push a tag and GitHub Actions builds all four platforms and publishes the release:

```bash
git tag v1.0.1
git push origin v1.0.1
```

The workflow in `.github/workflows/release.yml` handles the rest.

### Publish macOS-only from your Mac

GoReleaser is configured for local macOS builds only:

```bash
export GITHUB_TOKEN=ghp_...
git tag v1.0.1
git push origin v1.0.1
goreleaser release --clean
```

Linux binaries cannot be cross-compiled from macOS with CGO. Use the GitHub Actions workflow for full multi-platform releases.

Self-update only works when running a built binary (not `go run`).
