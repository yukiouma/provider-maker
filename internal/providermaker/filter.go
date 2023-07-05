package providermaker

type FilterFunc func(string) bool

var DefaultFilters = []FilterFunc{}
