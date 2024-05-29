package bot

type BotMode struct {
	adminMode bool
	editMode  bool
}

func NewBotMode() *BotMode {
	return &BotMode{}
}

func (r *BotMode) SetAdminMode(isAdmin bool) {
	r.adminMode = isAdmin
}

func (r *BotMode) SetEditMode(isEdit bool) {
	r.editMode = isEdit
}

func (r *BotMode) IsEditMode() bool {
	return r.editMode
}

func (r *BotMode) IsAdminMode() bool {
	return r.adminMode
}

func (r *BotMode) Reset() {
	r.adminMode = false
	r.editMode = false
}
