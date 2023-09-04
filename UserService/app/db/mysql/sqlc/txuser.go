package db

import (
	"context"
	"errors"
	"reflect"
	"server/UserService/pkg/hasher"

	"github.com/jinzhu/copier"
)

func (u *TxStore) RegisterUser(ctx context.Context, arg RegisterUserParams) (RegisterUserResult, error) {
	var result RegisterUserResult
	err := u.enableTx(ctx, func(q *Queries) error {
		user, err := q.CreateUser(ctx, CreateUserParams{
			Fullname: arg.Fullname,
			Email:    arg.Email,
			Gender:   arg.Gender,
		})

		if err != nil {
			return err
		}

		id, err := user.LastInsertId()
		if err != nil {
			return err
		}

		hashed, err := hasher.HashPassword(arg.Password)
		if err != nil {
			return err
		}

		err = q.CreateUserPassword(ctx, CreateUserPasswordParams{
			UserID:   int32(id),
			Password: hashed,
		})

		if err != nil {
			return err
		}

		result.ID = int32(id)

		return err
	})
	return result, err
}

func (u *TxStore) DeleteUser(ctx context.Context, id int32) error {
	return u.enableTx(ctx, func(q *Queries) error {
		err := q.DeleteUser(ctx, id)
		if err != nil {
			return err
		}

		err = q.DeleteUserPassword(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})
}

func (u *TxStore) LoginUser(ctx context.Context, arg LoginUserParams) (LoginUserResult, error) {
	var result LoginUserResult
	err := u.enableTx(ctx, func(q *Queries) error {
		user, err := q.GetUserByEmail(ctx, arg.Email)
		if err != nil {
			return err
		}

		userPassword, err := q.GetUserPassword(ctx, user.ID)
		if err != nil {
			return err
		}

		isRightPhone := user.Email == arg.Email

		if !isRightPhone {
			return errors.New("wrong phone or password")
		}

		isRightPassword, err := hasher.ComparePassword(userPassword.Password, arg.Password)
		if err != nil {
			return err
		}

		if !isRightPassword {
			return errors.New("wrong phone or password")
		}

		result.ID = user.ID

		return nil
	})

	return result, err
}

func (u *TxStore) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	return u.enableTx(ctx, func(q *Queries) error {
		var params UpdateUserInfoParams
		err := copier.Copy(&params, reflect.ValueOf(arg).Interface())
		if err != nil {
			return err
		}

		err = q.UpdateUserInfo(ctx, params)

		if err != nil {
			return err
		}

		hashed, err := hasher.HashPassword(arg.Password)
		if err != nil {
			return err
		}

		err = q.UpdateUserPassword(ctx, UpdateUserPasswordParams{
			UserID:   arg.ID,
			Password: hashed,
		})

		if err != nil {
			return err
		}

		return nil
	})
}
