//go:build !darwin && !(openbsd && !mips64) && !windows && !linux

package mid

func UUID() string {
	return newUUID()
}
