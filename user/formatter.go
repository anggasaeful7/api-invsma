package user

type UserFormatter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Nik      string `json:"nik"`
	Nohp     string `json:"nohp"`
	Npwp     string `json:"npwp"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Nik:      user.Nik,
		Nohp:     user.Nohp,
		Npwp:     user.Npwp,
		Token:    token,
		ImageURL: user.AvatarFileName,
	}

	return formatter
}
