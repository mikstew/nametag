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
go build -o nametag ./cmd/nametag
```

## Customize

Edit `displayName` and `tagColor` in `internal/config/config.go`, then rebuild.

## Releases and self-update

The refresh button on the nametag checks [GitHub Releases](https://github.com/mikio/nametag/releases) for a newer version. If one exists, it downloads the binary for your OS/arch and restarts.

Release assets must be named like `nametag-darwin-arm64` (also supports `_` instead of `-`). GoReleaser is configured in `.goreleaser.yaml` to produce these names.

To publish a release:

```bash
git tag v1.0.1
git push origin v1.0.1
goreleaser release --clean
```

Bump `Version` in `internal/config/config.go` before tagging, or rely on the GoReleaser ldflag that sets it at build time.

Self-update only works when running a built binary (not `go run`).
