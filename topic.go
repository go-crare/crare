package crare

import "gopkg.in/crare.v1/pkg/params"

type Topic struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	ThreadID          int    `json:"message_thread_id"`
}

// CreateTopic creates a topic in a forum supergroup chat.
func (b *Bot) CreateTopic(chat *Chat, topic *Topic) (*Topic, error) {
	params := map[string]any{
		"chat_id": chat.Recipient(),
		"name":    topic.Name,
	}

	if topic.IconColor != 0 {
		params["icon_color"] = topic.IconColor
	}
	if topic.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = topic.IconCustomEmojiID
	}

	data, err := b.Raw("createForumTopic", params)
	if err != nil {
		return nil, err
	}
	defer ReleaseBuffer(data)
	var resp Response[*Topic]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, err
}

// EditTopic edits name and icon of a topic in a forum supergroup chat.
func (b *Bot) EditTopic(chat *Chat, topic *Topic) error {
	params := map[string]any{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	if topic.Name != "" {
		params["name"] = topic.Name
	}
	if topic.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = topic.IconCustomEmojiID
	}

	data, err := b.Raw("editForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// CloseTopic closes an open topic in a forum supergroup chat.
func (b *Bot) CloseTopic(chat *Chat, topic *Topic) error {
	params := params.Topic{
		ChatID: chat.Recipient(),
		ID:     topic.ThreadID,
	}

	data, err := b.Raw("closeForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// ReopenTopic reopens a closed topic in a forum supergroup chat.
func (b *Bot) ReopenTopic(chat *Chat, topic *Topic) error {
	params := params.Topic{
		ChatID: chat.Recipient(),
		ID:     topic.ThreadID,
	}

	data, err := b.Raw("reopenForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// DeleteTopic deletes a forum topic along with all its messages in a forum supergroup chat.
func (b *Bot) DeleteTopic(chat *Chat, topic *Topic) error {
	params := params.Topic{
		ChatID: chat.Recipient(),
		ID:     topic.ThreadID,
	}

	data, err := b.Raw("deleteForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// UnpinAllTopicMessages clears the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllTopicMessages(chat *Chat, topic *Topic) error {
	params := params.Topic{
		ChatID: chat.Recipient(),
		ID:     topic.ThreadID,
	}

	data, err := b.Raw("unpinAllForumTopicMessages", &params)
	ReleaseBuffer(data)
	return err
}

// TopicIconStickers gets custom emoji stickers, which can be used as a forum topic icon by any user.
func (b *Bot) TopicIconStickers() ([]Sticker, error) {
	data, err := b.Raw("getForumTopicIconStickers")
	if err != nil {
		return nil, err
	}
	defer ReleaseBuffer(data)

	var resp Response[[]Sticker]
	if err := b.json.NewDecoder(data).Decode(&resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// EditGeneralTopic edits name of the 'General' topic in a forum supergroup chat.
func (b *Bot) EditGeneralTopic(chat *Chat, topic *Topic) error {
	params := map[string]any{
		"chat_id": chat.Recipient(),
		"name":    topic.Name,
	}

	data, err := b.Raw("editGeneralForumTopic", params)
	ReleaseBuffer(data)
	return err
}

// CloseGeneralTopic closes an open 'General' topic in a forum supergroup chat.
func (b *Bot) CloseGeneralTopic(chat *Chat, topic *Topic) error {
	params := params.OnlyID{
		ChatID: chat.Recipient(),
	}

	data, err := b.Raw("closeGeneralForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// ReopenGeneralTopic reopens a closed 'General' topic in a forum supergroup chat.
func (b *Bot) ReopenGeneralTopic(chat *Chat) error {
	params := params.OnlyID{
		ChatID: chat.Recipient(),
	}

	data, err := b.Raw("reopenGeneralForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// HideGeneralTopic hides the 'General' topic in a forum supergroup chat.
func (b *Bot) HideGeneralTopic(chat *Chat) error {
	params := params.OnlyID{
		ChatID: chat.Recipient(),
	}

	data, err := b.Raw("hideGeneralForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// UnhideGeneralTopic unhides the 'General' topic in a forum supergroup chat.
func (b *Bot) UnhideGeneralTopic(chat *Chat) error {
	params := params.OnlyID{
		ChatID: chat.Recipient(),
	}

	data, err := b.Raw("unhideGeneralForumTopic", &params)
	ReleaseBuffer(data)
	return err
}

// Use this method to clear the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot) UnpinAllGeneralForumTopicMessages(chat *Chat) error {
	params := params.OnlyID{
		ChatID: chat.Recipient(),
	}

	data, err := b.Raw("unpinAllGeneralForumTopicMessages", &params)
	ReleaseBuffer(data)
	return err
}
