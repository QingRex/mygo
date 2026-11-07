//go:build !(linux && (amd64 || arm64)) && !darwin && !windows && !js

package video

import "github.com/egoist/mygo/internal/media"

func newEngine(source) (media.Engine, error) { return nil, ErrUnsupported }
