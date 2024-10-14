package handlers

import (
	"AABBCCDD/app/types"

	"github.com/wrongheaven/superkit/kit"
)

func HandleAuthentication(kit *kit.Kit) (kit.Auth, error) {
	return types.AuthUser{}, nil
}
