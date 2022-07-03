package db

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	"github.com/leaderseek/definition/activity"
	"github.com/leaderseek/definition/activity/db/param"
	"github.com/leaderseek/sqlboiler/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func (cfg *Config) TeamCreate(ctx context.Context, req *param.TeamCreateRequest) (*param.TeamCreateResponse, error) {
	tx, dbClose, err := cfg.beginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer dbClose()

	team, err := teamInsert(ctx, tx, req.Team)
	if err != nil {
		return nil, err
	}

	for _, player := range req.Players {
		if err := teamMemberAndPlayerInsert(ctx, tx, team.ID, player); err != nil {
			return nil, err
		}
	}

	return &param.TeamCreateResponse{TeamID: team.ID}, tx.Commit()
}

func (cfg *Config) TeamAddMembers(ctx context.Context, req *param.TeamAddMembersRequest) error {
	tx, dbClose, err := cfg.beginTx(ctx)
	if err != nil {
		return err
	}
	defer dbClose()

	teamExists, err := repository.TeamExists(ctx, tx, req.TeamID)
	if err != nil {
		return activity.ErrorWithRollback(tx, err, "failed to query from team table")
	}

	if !teamExists {
		return activity.NewErrorWithRollback(tx, "failed team existence check")
	}

	for _, player := range req.Players {
		if err := teamMemberAndPlayerInsert(ctx, tx, req.TeamID, player); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func teamInsert(ctx context.Context, tx *sql.Tx, reqTeam *repository.Team) (*repository.Team, error) {
	team, err := repository.Teams(repository.TeamWhere.DisplayName.EQ(reqTeam.DisplayName)).One(ctx, tx)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, activity.ErrorWithRollback(tx, err, "failed to query from team table")
		}

		if err := reqTeam.Insert(ctx, tx, boil.Infer()); err != nil {
			return nil, activity.ErrorWithRollback(tx, err, "failed to insert into team table")
		}

		team = reqTeam
	}

	return team, nil
}
