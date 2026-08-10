package maxon

type (
	UpdateType    string
	MessageFormat string
)

const (
	UpdateTypeBotAdded         UpdateType = "bot_added"
	UpdateTypeBotStarted       UpdateType = "bot_started"
	UpdateTypeBotStopped       UpdateType = "bot_stopped"
	UpdateTypeBotRemoved       UpdateType = "bot_removed"
	UpdateTypeChatTitleChanged UpdateType = "chat_title_changed"
	UpdateTypeDialogCleared    UpdateType = "dialog_cleared"
	UpdateTypeDialogMuted      UpdateType = "dialog_muted"
	UpdateTypeDialogUnmuted    UpdateType = "dialog_unmuted"
	UpdateTypeDialogRemoved    UpdateType = "dialog_removed"
	UpdateTypeMessageCallback  UpdateType = "message_callback"
	UpdateTypeMessageCreated   UpdateType = "message_created"
	UpdateTypeMessageEdited    UpdateType = "message_edited"
	UpdateTypeMessageRemoved   UpdateType = "message_removed"
	UpdateTypeUserAdded        UpdateType = "user_added"
	UpdateTypeUserRemoved      UpdateType = "user_removed"

	MessageFormatMarkdown MessageFormat = "markdown"
	MessageFormatHtml     MessageFormat = "html"
)

type User struct {
	ID               int64  `json:"user_id"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Username         string `json:"username"`
	IsBot            bool   `json:"is_bot"`
	LastActivityTime int64  `json:"last_activity_time"`
	Description      string `json:"description"`
	AvatarUrl        string `json:"avatar_url"`
	FullAvatarUrl    string `json:"full_avatar_url"`
}

type Message struct {
	MessageID string `json:"message_id"`
	Timestamp int64  `json:"timestamp"`
	Sender    User   `json:"sender"`
	Recipient struct {
		ChatID   int64  `json:"chat_id"`
		ChatType string `json:"chat_type"`
		UserID   int64  `json:"user_id"`
	} `json:"recipient"`
	Body struct {
		Mid  string `json:"mid"`
		Seq  int64  `json:"seq"`
		Text string `json:"text"`
	} `json:"body"`
}

type CallbackPayload struct {
	CallbackID string `json:"callback_id"`
	Payload    string `json:"payload,omitempty"`
	User       User   `json:"user"`
}

type Update struct {
	Type       UpdateType       `json:"update_type"`
	Timestamp  int64            `json:"timestamp"`
	ChatID     int64            `json:"chat_id,omitempty"`
	UserID     int64            `json:"user_id,omitempty"`
	User       *User            `json:"user,omitempty"`
	UserLocale string           `json:"user_locale,omitempty"`
	IsChannel  bool             `json:"is_channel,omitempty"`
	Title      string           `json:"title,omitempty"`
	MessageID  string           `json:"message_id,omitempty"`
	Payload    string           `json:"payload,omitempty"`
	MutedUntil int64            `json:"muted_until,omitempty"`
	Message    *Message         `json:"message,omitempty"`
	Callback   *CallbackPayload `json:"callback,omitempty"`
}

type ResponseUpdates struct {
	Market  int64    `json:"market"`
	Updates []Update `json:"updates"`
}

type MessageSendOptions struct {
	DisableLinkPreview bool          `json:"disable_link_preview,omitempty"`
	NotifyDisable      bool          `json:"notify_disable,omitempty"`
	Format             MessageFormat `json:"format,omitempty"`
}
