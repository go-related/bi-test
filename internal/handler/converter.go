package handler

import "entdemo/ent"

func convertUserToDomainModel(in *UserDTO) *ent.User {
	if in == nil {
		return nil
	}
	user := ent.User{
		Age:      in.Age,
		Name:     in.Name,
		Email:    in.Email,
		Nickname: in.Nickname,
	}

	if in.ID != nil {
		user.ID = *in.ID
	}
	return &user
}

func convertUserToDTO(in *ent.User) *UserDTO {
	if in == nil {
		return nil
	}
	return &UserDTO{
		ID:       &in.ID,
		Age:      in.Age,
		Name:     in.Name,
		Email:    in.Email,
		Nickname: in.Nickname,
	}
}
