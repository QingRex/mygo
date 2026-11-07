//go:build darwin || windows || js

package video

import "github.com/egoist/mygo/internal/media"

// The engines of these systems are next.
func newEngine(source) (media.Engine, error) { return nil, ErrUnsupported }
