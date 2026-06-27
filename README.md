# nametag

A small desktop app that shows a customizable nametag in a window.

## Project layout

```
cmd/nametag/              Application entrypoint
.github/workflows/        CI and release automation
internal/
  app/                    Window lifecycle and wiring
  config/                 Hardcoded nametag settings and version
  platform/               OS-specific helpers (restart)
  ui/nametag/             Nametag view
  update/                 GitHub Releases self-update
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

Release assets are named `nametag-darwin-arm64`, `nametag-linux-amd64`, etc.

### Publish a release

Commit your changes, push, tag, and push the tag. GitHub Actions builds all four platforms and publishes the release:

```bash
git push origin main
git tag v1.0.2
git push origin v1.0.2
```

The workflow in `.github/workflows/release.yml` runs on every `v*` tag push.

Self-update only works when running a built binary (not `go run`).
