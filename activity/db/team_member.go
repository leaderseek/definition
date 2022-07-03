package db

import (
	"context"
	"database/sql"

	"github.com/leaderseek/definition/activity"
	"github.com/leaderseek/sqlboiler/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func teamMemberAndPlayerInsert(ctx context.Context, tx *sql.Tx, teamID string, reqPlayer *repository.Player) error {
	player, err := playerInsert(ctx, tx, reqPlayer)
	if err != nil {
		return err
	}

	teamMember := &repository.TeamMember{
		ID:     player.ID,
		TeamID: teamID,
	}

	if err := teamMember.Insert(ctx, tx, boil.Infer()); err != nil {
		return activity.ErrorWithRollback(tx, err, "failed to insert into team_member table")
	}

	return nil
}
