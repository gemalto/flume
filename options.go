package flume

// An CoreOption configures a Core.
type CoreOption interface {
	apply(*Core)
}

// coreOptionFunc wraps a func so it satisfies the CoreOption interface.
type coreOptionFunc func(*Core)

func (f coreOptionFunc) apply(c *Core) {
	f(c)
}

// AddCallerSkip increases the number of callers skipped by caller annotation
// (as enabled by the AddCaller option). When building wrappers around a
// Core, supplying this CoreOption prevents Core from always
// reporting the wrapper code as the caller.
func AddCallerSkip(skip int) CoreOption {
	return coreOptionFunc(func(c *Core) {
		c.callerSkip += skip
	})
}
