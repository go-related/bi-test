package repository

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/user"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

func (db *DB) GetAllUsers(ctx context.Context, name string) ([]*ent.User, error) {
	query := db.client.User.Query()

	if name != "" {
		query = query.Where(user.NameEQ(name))
	}

	result, err := query.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load users")
	}
	return result, err
}

func (db *DB) GetUserById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	result, err := db.client.User.Get(ctx, id)

	if err != nil {
		return nil, errors.Wrap(err, "failed to load users")
	}
	return result, err
}

func (db *DB) Delete(ctx context.Context, id uuid.UUID) error {
	err := db.client.User.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (db *DB) Create(ctx context.Context, usr *ent.User) (*ent.User, error) {
	userCreated, err := db.client.User.Create().
		SetEmail(usr.Email).
		SetName(usr.Name).
		SetAge(usr.Age).
		SetNickname(usr.Nickname).
		Save(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create user")
	}
	return userCreated, nil
}

func (db *DB) Update(ctx context.Context, usr *ent.User) error {
	_, err := db.client.User.Update().Where(user.ID(usr.ID)).
		SetEmail(usr.Email).
		SetName(usr.Name).
		SetAge(usr.Age).
		SetNickname(usr.Nickname).
		SetUpdatedAt(time.Now()).
		Save(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (db *DB) GetUserFriends(ctx context.Context, id uuid.UUID) ([]*ent.User, error) {
	friends, err := db.client.User.Query().Where(user.IDEQ(id)).WithFriends().All(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "failed to load users friends")
	}

	return friends, err
}

func (db *DB) UpdateFriends(ctx context.Context, id uuid.UUID, newFriendIds []uuid.UUID) error {
	friends, err := db.client.User.Query().Where(user.IDIn(newFriendIds...)).All(ctx)

	recordsAffected, err := db.client.User.Update().Where(user.ID(id)).
		AddFriends(friends...).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if recordsAffected == 0 || err != nil {
		return errors.Wrap(err, "failed to update user friends")
	}
	return nil
}
