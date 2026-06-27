package update

import (
	"context"
	"fmt"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
)

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
}

// New creates an update service for the given repository slug and version.
func New(repo, version string) *Service {
	return &Service{repo: repo, version: version}
}

// CheckAndApply fetches the latest release and replaces the running binary when newer.
func (s *Service) CheckAndApply(ctx context.Context) (Result, error) {
	repo := selfupdate.ParseSlug(s.repo)

	latest, found, err := selfupdate.DetectLatest(ctx, repo)
	if err != nil {
		return Result{}, fmt.Errorf("check for updates: %w", err)
	}
	if !found {
		return Result{}, fmt.Errorf("no release found for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	if latest.LessOrEqual(s.version) {
		return Result{
			Message: fmt.Sprintf("Already on the latest version (v%s).", s.version),
		}, nil
	}

	if _, err := selfupdate.UpdateSelf(ctx, s.version, repo); err != nil {
		return Result{}, fmt.Errorf("download update: %w", err)
	}

	return Result{
		Updated: true,
		Version: latest.Version(),
		Message: fmt.Sprintf("Updated to v%s.", latest.Version()),
	}, nil
}
