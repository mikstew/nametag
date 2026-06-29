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
  log/                    Structured logging helpers
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

The app checks [GitHub Releases](https://github.com/mikstew/nametag/releases) every minute for a newer version. If one exists, it verifies the binary against `checksums.txt` published with the release, then downloads and restarts.

Release assets are named `nametag-darwin-arm64`, `nametag-linux-amd64`, `nametag-windows-amd64.exe`, etc., plus a `checksums.txt` manifest.

### Publish a release

Commit your changes, push, tag, and push the tag. GitHub Actions builds all five platform binaries and publishes the release:

```bash
git push origin main
git tag v1.0.2
git push origin v1.0.2
```

The workflow in `.github/workflows/release.yml` runs on every `v*` tag push.

Self-update only works when running a built binary (not `go run`).

## Tests

```bash
go test ./...
```

## Write-up

# Overview

For this challenge I built a small cross-platform desktop app in Go using Fyne that displays a nametag with configurable name and background color. Name and color are set at build time in config, then baked into each release binary. Despite this being a very simple app I tried to structure it like a standard Go app layout.

Source lives on GitHub ([mikstew/nametag](https://github.com/mikstew/nametag)). I use GitHub Actions for CI/CD: pushing a version tag (e.g. v1.0.8) triggers a workflow that builds native binaries for macOS (arm64, amd64), Linux (arm64, amd64), and Windows (amd64). Because Fyne uses CGO, each platform is built on its own runner rather than cross-compiled from a single machine. Artifacts are published to GitHub Releases ([here](https://github.com/mikstew/nametag/releases)).

The app polls GitHub Releases every minute for a newer version. When one is found, it downloads the correct binary for the host OS/arch, replaces the running executable on disk, and performs a handoff restart (the new instance starts and waits for the old process to exit before showing the window). Later launches use the updated binary automatically. A true seamless hot-swap of the Go running binary didn't seem possible, so I opted to coordinate a handoff from the old to new binary with a brief gap during the transition.

Downloads are not applied blindly. Each release includes a `checksums.txt` file with SHA256 hashes for every platform binary. The app uses `go-selfupdate`'s `ChecksumValidator` to verify the downloaded executable against that manifest before replacing itself. If the checksum is missing or does not match, the update is rejected and the running version is left unchanged. This protects against corrupted downloads; it does not defend against a compromised release pipeline (that would require signed checksums or code signing as a further step).

# Questions

1) Q: To AI or not to AI?
   A: I chose to use AI. AI tools are common in modern software development, the spec didn’t prohibit them, and this challenge combined several unfamiliar pieces—Fyne, GitHub Actions for CGO builds, and self-updating binaries—where AI was useful for getting started quickly.

   I don’t treat AI output as final. I structured the project like a production Go app, chose GitHub Actions over local cross-compilation, tested releases on macOS and Windows, and fixed real bugs (wrong repo slug, Linux linker flags, update handoff). The result is a working app I understand and can maintain.
