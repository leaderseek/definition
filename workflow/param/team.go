package param

import "github.com/leaderseek/sqlboiler/repository"

type TeamCreateRequest struct {
	DBConnectionString string
	Team               *repository.Team
	Players            []*repository.Player
}

type TeamCreateResponse struct {
	TeamID string
}

type TeamAddMembersRequest struct {
	DBConnectionString string
	TeamID             string
	Players            []*repository.Player
}
