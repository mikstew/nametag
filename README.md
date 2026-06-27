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
- Linux: X11 or Wayland, OpenGL/Mesa dev libraries (for local builds)
- Windows: MinGW-w64 GCC (included with Go on Windows)

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

The `-extldflags` flag is macOS-only (silences a harmless `-lobjc` linker warning).

## Customize

Edit `DisplayName` and `TagColor` in `internal/config/config.go`, then rebuild.

## Releases and self-update

The app checks [GitHub Releases](https://github.com/mikstew/nametag/releases) every minute for a newer version. If one exists, it downloads the binary and restarts.

Release assets are named `nametag-darwin-arm64`, `nametag-linux-amd64`, `nametag-windows-amd64.exe`, etc.

### Publish a release

Commit your changes, push, tag, and push the tag. GitHub Actions builds all five platform binaries and publishes the release:

```bash
git push origin main
git tag v1.0.2
git push origin v1.0.2
```

The workflow in `.github/workflows/release.yml` runs on every `v*` tag push.

Self-update only works when running a built binary (not `go run`).
