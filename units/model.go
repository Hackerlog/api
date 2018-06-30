package units

import (
	"time"

	"github.com/jinzhu/gorm/dialects/postgres"
)

// Unit This stuff
type Unit struct {
	ID              uint           `json:"id" gorm:"primary_key"`
	UserID          uint           `json:"user_id"`
	EditorType      string         `json:"editor_type"`
	ProjectName     string         `json:"project_name"`
	Languages       postgres.Jsonb `json:"languages"`
	LocWritten      int            `json:"loc_written"`
	LocDeleted      int            `json:"loc_deleted"`
	FilesEdited     int            `json:"files_edited"`
	NumberOfCommits int            `json:"number_of_commits"`
	ComputerType    string         `json:"computer_type"`
	Os              string         `json:"os"`
	StartedAt       time.Time      `json:"started_at"`
	StoppedAt       time.Time      `json:"stopped_at"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       *time.Time     `json:"-" sql:"index"`
}
