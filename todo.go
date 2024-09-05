package todo

type TodoList struct {
	Id          int    `json:"id" db:"id" `
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int `db:"user_id"`
	ListId int `db:"list_id"`
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (i *UpdateItemInput) IsValid() bool {
	if i.Title != nil || i.Description != nil || i.Done != nil {
		return true
	}
	return false
}

func (i *UpdateListInput) IsValid() bool {
	if i.Title != nil || i.Description != nil {
		return true
	}
	return false
}
