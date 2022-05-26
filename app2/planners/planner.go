package planners

import (
	"errors"
	"fmt"

	"github.com/kudrykv/latex-yearly-planner/app2/devices"
)

const (
	MonthsOnSidesTemplate = "mos"
)

type Planner interface {
	GenerateFor(devices.Device) error
	WriteTo(dir string) error
}

var UnknownTemplateName = errors.New("unknown planner name")

func New(params Params) (Planner, error) {
	switch params.Name {
	case MonthsOnSidesTemplate:
		return newMonthsOnSides(params)
	default:
		return nil, fmt.Errorf("%s: %w", params.Name, UnknownTemplateName)
	}
}
