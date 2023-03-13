package humanizer

import "github.com/gosimple/slug"

type Url interface {
	Generate(text string) string
}

type ManagerHumanizer struct {
	Sub map[string]string
}

func NewUrlHumanizer(sub map[string]string) *ManagerHumanizer {
	return &ManagerHumanizer{
		Sub: sub,
	}
}
func (h *ManagerHumanizer) Generate(text string) string {
	slug.CustomSub = h.Sub
	return slug.Make(text)
}
