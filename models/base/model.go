package base

import (
	"time"
)

type Entity struct {
	Created   time.Time  `json:"created"`
	Createdby int        `json:"createdby"`
	Updated   *time.Time `json:"updated"`
	Updatedby *int       `json:"updatedby"`
	Deleted   *time.Time `json:"deleted"`
	Deletedby *int       `json:"deletedby"`
}
