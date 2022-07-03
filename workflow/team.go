package workflow

import (
	"time"

	"github.com/leaderseek/definition/activity/db"
	db_param "github.com/leaderseek/definition/activity/db/param"
	"github.com/leaderseek/definition/workflow/param"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func TeamCreate(ctx workflow.Context, req *param.TeamCreateRequest) (*param.TeamCreateResponse, error) {
	ao := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	aoCtx := workflow.WithActivityOptions(ctx, ao)

	cfg := &db.Config{
		ConnectionString: req.DBConnectionString,
	}

	teamCreateReq := &db_param.TeamCreateRequest{
		Team:    req.Team,
		Players: req.Players,
	}

	f := workflow.ExecuteActivity(aoCtx, cfg.TeamCreate, teamCreateReq)

	var teamID string
	if err := f.Get(ctx, &teamID); err != nil {
		return nil, err
	}

	return &param.TeamCreateResponse{TeamID: teamID}, nil
}

func TeamAddMembers(ctx workflow.Context, req *param.TeamAddMembersRequest) error {
	ao := workflow.ActivityOptions{
		ScheduleToCloseTimeout: 5 * time.Second,
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
	}
	aoCtx := workflow.WithActivityOptions(ctx, ao)

	cfg := &db.Config{
		ConnectionString: req.DBConnectionString,
	}

	teamAddMembersReq := &db_param.TeamAddMembersRequest{
		TeamID:  req.TeamID,
		Players: req.Players,
	}

	f := workflow.ExecuteActivity(aoCtx, cfg.TeamAddMembers, teamAddMembersReq)

	if err := f.Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
