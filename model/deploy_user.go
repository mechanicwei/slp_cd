package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type DeployUser struct {
	Name      string `json:"name" db:"name"`
	AvatarUrl string `json:"avatar_url"`
	GithubUrl string `json:"github_url"`
}

func (this *DeployUser) Scan(src interface{}) error {
	source, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion .([]byte) failed.")
	}

	err := json.Unmarshal(source, this)
	if err != nil {
		return err
	}
	return nil
}

func (this DeployUser) Value() (driver.Value, error) {
	j, err := json.Marshal(this)
	return j, err
}
