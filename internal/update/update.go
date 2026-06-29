package update

import (
	"context"
	"fmt"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"

	"github.com/mikio/nametag/internal/log"
)

const checksumsFile = "checksums.txt"

// Result describes the outcome of an update check.
type Result struct {
	Updated bool
	Version string
	Message string
}

// Service checks GitHub Releases and applies updates.
type Service struct {
	repo    string
	version string
	updater *selfupdate.Updater
}

// New creates an update service for the given repository slug and version.
func New(repo, version string) (*Service, error) {
	updater, err := selfupdate.NewUpdater(selfupdate.Config{
		Validator: &selfupdate.ChecksumValidator{
			UniqueFilename: checksumsFile,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("create updater: %w", err)
	}

	log.Info("update service ready", "repo", repo, "version", version)

	return &Service{
		repo:    repo,
		version: version,
		updater: updater,
	}, nil
}

// CheckAndApply fetches the latest release and replaces the running binary when newer.
func (s *Service) CheckAndApply(ctx context.Context) (Result, error) {
	repo := selfupdate.ParseSlug(s.repo)

	log.Debug("checking for updates", "current", s.version)

	latest, found, err := s.updater.DetectLatest(ctx, repo)
	if err != nil {
		return Result{}, fmt.Errorf("check for updates: %w", err)
	}
	if !found {
		return Result{}, fmt.Errorf("no release found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(s.version) {
		log.Debug("already on latest version", "current", s.version, "latest", latest.Version())
		return Result{
			Message: fmt.Sprintf("Already on the latest version (v%s).", s.version),
		}, nil
	}

	log.Info("downloading update", "current", s.version, "latest", latest.Version(), "asset", latest.AssetName)

	if _, err := s.updater.UpdateSelf(ctx, s.version, repo); err != nil {
		return Result{}, fmt.Errorf("download update: %w", err)
	}

	log.Info("update installed", "version", latest.Version())

	return Result{
		Updated: true,
		Version: latest.Version(),
		Message: fmt.Sprintf("Updated to v%s.", latest.Version()),
	}, nil
}
