package db

import (
	"context"
	"database/sql"

	"github.com/friendsofgo/errors"
	"github.com/leaderseek/definition/activity"
	"github.com/leaderseek/sqlboiler/repository"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func playerInsert(ctx context.Context, tx *sql.Tx, reqPlayer *repository.Player) (*repository.Player, error) {
	player, err := repository.Players(repository.PlayerWhere.EmailAddress.EQ(reqPlayer.EmailAddress)).One(ctx, tx)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, activity.ErrorWithRollback(tx, err, "failed to query from player table")
		}

		if err := reqPlayer.Insert(ctx, tx, boil.Infer()); err != nil {
			return nil, activity.ErrorWithRollback(tx, err, "failed to insert into player table")
		}

		player = reqPlayer
	}

	return player, nil
}
