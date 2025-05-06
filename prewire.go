package prewire

type ProviderSet struct{}

// NewSet is a stub method. It exists to provide an is readable by the prewire command
func NewSet(providers ...interface{}) ProviderSet {
	return ProviderSet{}
}

// Union is a stub method. It exists to provide an is readable by the prewire command
// Use it to join other Provider sets to the current set
func (pm ProviderSet) Union(psets ...ProviderSet) ProviderSet {
	return pm
}
