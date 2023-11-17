package types

import (
	"fmt"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	CryptCost         = 12
	minFirstName byte = 2
	minLastName  byte = 2
	minPassword  byte = 7
)

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"firstName"`
	LastName          string             `bson:"lasttName" json:"lastName"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedpassword" json:"-"`
}

type UpdateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (p *UpdateUserParams) ToBson() bson.M {
	m := bson.M{}
	if len(p.FirstName) > 0 {
		m["firstname"] = p.FirstName
	}
	if len(p.LastName) > 0 {
		m["lastname"] = p.LastName
	}
	return m
}

type CreateUserParams struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (params CreateUserParams) Validate() error {
	if byte(len(params.FirstName)) < minFirstName {
		return fmt.Errorf("Firstname length short, must be %d char min", minFirstName)
	}
	if byte(len(params.LastName)) < minLastName {
		return fmt.Errorf("Lastname length short, must be %d char min", minLastName)
	}
	if byte(len(params.Password)) < minPassword {
		return fmt.Errorf("Password length short, must be %d char min", minPassword)
	}
	if isEmailValid(params.Email) == false {
		fmt.Errorf("the insert email is incorrect or invalid")
	}
	return nil
}

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	toByte := []byte(params.Password) // convert strings to byte array

	encrypted, err := bcrypt.GenerateFromPassword(toByte, CryptCost)
	if err != nil {
		log.Fatalln("cant encrypt")
		return nil, err
	}
	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encrypted),
	}, nil
}
