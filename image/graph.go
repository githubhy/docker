package image

import (
	"github.com/dotcloud/docker/graphdriver"
)

type Graph interface {
	Get(id string) (*Image, error)
	ImageRoot(id string) string
	Driver() graphdriver.Driver
}
