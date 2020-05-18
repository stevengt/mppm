package versioning

type Versioner interface {
	Init() (err error)
	Git(args ...string) (err error)
}
