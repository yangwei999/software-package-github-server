package repo

type Repo interface {
	CreateRepo(repo string) (string, error)
}
