package domain

import "encoding/json"

const PlatformGithub = "github"

type RepoCreatedEvent struct {
	PkgId        string `json:"pkg_id"`
	Platform     string `json:"platform"`
	RepoLink     string `json:"repo_link"`
	FailedReason string `json:"failed_reason"`
}

func (e *RepoCreatedEvent) Message() ([]byte, error) {
	return json.Marshal(e)
}

func NewRepoCreatedEvent(pkgId, url string) RepoCreatedEvent {
	return RepoCreatedEvent{
		PkgId:    pkgId,
		Platform: PlatformGithub,
		RepoLink: url,
	}
}

func NewCodePushedEvent(pkgId string) RepoCreatedEvent {
	return RepoCreatedEvent{
		PkgId:    pkgId,
		Platform: PlatformGithub,
	}
}
