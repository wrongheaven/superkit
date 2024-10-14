package handlers

import (
	"AABBCCDD/app/views/landing"

	"github.com/wrongheaven/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	return kit.Render(landing.Index())
}
