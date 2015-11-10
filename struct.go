package gochatwork

import (
	"encoding/json"
)

// Me is /me response struct
type Me struct {
	AccountID        int64 `json:"account_id"`
	RoomID           int64 `json:"room_id"`
	Name             string
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int64  `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string
	Title            string
	URL              string
	Introduction     string
	Mail             string
	TelOrganization  string `json:"tel_organization"`
	TelExtension     string `json:"tel_extension"`
	TelMobile        string `json:"tel_mobile"`
	Skype            string
	Facebook         string
	Twitter          string
	AvatarImageURL   string `json:"avatar_image_url"`
}

// Status is /me response struct
type Status struct {
	UnreadRoomNum  int64 `json:"unread_room_num"`
	MentionRoomNum int64 `json:"mention_room_num"`
	MytaskRoomNum  int64 `json:"mytask_room_num"`
	UnreadNum      int64 `json:"unread_num"`
	MentionNum     int64 `json:"mention_num"`
	MytaskNum      int64 `json:"mytask_num"`
}

// Room is room struct
type Room struct {
	RoomID         int64 `json:"room_id"`
	Name           string
	Type           string
	Role           string
	Sticky         bool
	UnreadNum      int64  `json:"unread_num"`
	MentionNum     int64  `json:"mention_num"`
	MytaskNum      int64  `json:"mytask_num"`
	MessageNum     int64  `json:"message_num"`
	FileNum        int64  `json:"file_num"`
	TaskNum        int64  `json:"task_num"`
	IconPath       string `json:"icon_path"`
	Description    string
	LastUpdateTime int64 `json:"last_update_time"`
}

// Account is task struct
type Account struct {
	AccountID        int64 `json:"account_id"`
	Name             string
	Role             string
	ChatworkID       string `json:"chatwork_id"`
	OrganizationID   int64  `json:"organization_id"`
	OrganizationName string `json:"organization_name"`
	Department       string
	AvatarImageURL   string `json:"avatar_image_url"`
}

// Task is task struct
type Task struct {
	TaskID            int64 `json:"task_id"`
	Room              Room
	Account           Account
	AssignedByAccount Account `json:"assigned_by_account"`
	MessageID         int64   `json:"message_id"`
	Body              string
	LimitTime         int64 `json:"limit_time"`
	Status            string
}

// Message is message struct
type Message struct {
	MessageID  int64 `json:"message_id"`
	Account    Account
	Body       string
	SendTime   int64 `json:"send_time"`
	UpdateTime int64 `json:"update_time"`
}

// File is file struct
type File struct {
	FileID      int64 `json:"file_id"`
	Account     Account
	MessageID   int64 `json:"message_id"`
	Filename    string
	Filesize    int64
	UploadTime  int64  `json:"upload_time"`
	DownloadURL string `json:"download_url"`
}

func setSturctFromJSON(b []byte, v interface{}, err error) error {
	if err != nil {
		return err
	}

	return json.Unmarshal(b, &v)
}
