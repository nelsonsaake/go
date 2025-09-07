package dto

type InclusionStrategy string

const (
	IncludeAll InclusionStrategy = "include_all"
	ExcludeAll InclusionStrategy = "exclude_all"
)

var (
	inclusionStrategy = IncludeAll
)

func SetInclusionStrategy(v string) {
	inclusionStrategy = InclusionStrategy(v)
}

func GetInclusionStrategy() InclusionStrategy {
	return inclusionStrategy
}

func isIncludeAll() bool {
	return GetInclusionStrategy() == IncludeAll
}
