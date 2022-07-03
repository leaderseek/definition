package param

import (
	"github.com/leaderseek/sqlboiler/repository"
)

type TeamCreateRequest struct {
	Team    *repository.Team
	Players []*repository.Player
}

type TeamCreateResponse struct {
	TeamID string
}

type TeamAddMembersRequest struct {
	TeamID  string
	Players []*repository.Player
}
