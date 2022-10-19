package services

type lifetime int

// Dependency Injections Lifetime out-of-box Dependency Injection styles
const (
	Transient lifetime = iota
	Singleton
	Scoped
)
