package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	GrpcUserService "server/MainService/GrpcClients/UserService"
	"server/MainService/config"
	_struct "server/MainService/struct"
	"server/MainService/utils"
)

type UserApiHanlder struct {
	Repo GrpcUserService.UserServiceRepository
	C    config.Config
}

func NewUserApiHanlder(c config.Config, repo GrpcUserService.UserServiceRepository) UserApiHanlder {
	return UserApiHanlder{
		Repo: repo,
		C:    c,
	}
}

func (u *UserApiHanlder) LoginUser(w http.ResponseWriter, r *http.Request) {
	Phone := r.FormValue("Phone")
	Password := r.FormValue("Password")
	loginRequest := GrpcUserService.LoginUserRequest{
		Phone:    Phone,
		Password: Password,
	}

	res := u.Repo.LoginUser(&loginRequest)
	out, _ := json.Marshal(res)

	// jwt generate token
	jwtToken := utils.NewJwtUtils(u.C)

	token, err := jwtToken.GenerateToken(utils.InfoInJwt{
		UserId: int(res.UserId),
	})

	if err != nil {
		log.Println(err)
	}

	// set cookie for user
	cookie := http.Cookie{
		Name:  "gme",
		Value: token,
	}

	http.SetCookie(w, &cookie)

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Setting content type to json
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(out))
}

func (u *UserApiHanlder) RegisterUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
	}

	Phone := r.FormValue("Phone")
	Password := r.FormValue("Password")
	FullName := r.FormValue("FullName")
	Gender := r.FormValue("Gender")

	log.Printf("%v %v %v\n", Phone, FullName, Gender)
	registerRequest := GrpcUserService.RegisterUserRequest{
		Phone:    Phone,
		Password: Password,
		Fullname: FullName,
		Gender:   Gender,
	}

	response := u.Repo.RegisterUser(&registerRequest)
	message := _struct.ApiMessage{
		ErrorCode: response.ErrorCode,
		Message:   "success",
		Data:      utils.Convert(_struct.GetRegisterUserResponse(response)),
	}

	returnString := utils.Convert(message)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	//Setting content type to json
	w.Header().Set("Content-Type", "application/json")
	// write json response
	w.Write([]byte(returnString))
}
