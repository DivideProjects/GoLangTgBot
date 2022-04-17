// THIS FILE IS AUTOGENERATED. DO NOT EDIT.
// Regen by running 'go generate' in the repo root.

package gotgbot

// Answer Helper method for Bot.AnswerCallbackQuery
func (cq CallbackQuery) Answer(b *Bot, opts *AnswerCallbackQueryOpts) (bool, error) {
	return b.AnswerCallbackQuery(cq.Id, opts)
}

// Answer Helper method for Bot.AnswerInlineQuery
func (iq InlineQuery) Answer(b *Bot, results []InlineQueryResult, opts *AnswerInlineQueryOpts) (bool, error) {
	return b.AnswerInlineQuery(iq.Id, results, opts)
}

// Answer Helper method for Bot.AnswerPreCheckoutQuery
func (pcq PreCheckoutQuery) Answer(b *Bot, ok bool, opts *AnswerPreCheckoutQueryOpts) (bool, error) {
	return b.AnswerPreCheckoutQuery(pcq.Id, ok, opts)
}

// Answer Helper method for Bot.AnswerShippingQuery
func (sq ShippingQuery) Answer(b *Bot, ok bool, opts *AnswerShippingQueryOpts) (bool, error) {
	return b.AnswerShippingQuery(sq.Id, ok, opts)
}

// ApproveJoinRequest Helper method for Bot.ApproveChatJoinRequest
func (c Chat) ApproveJoinRequest(b *Bot, userId int64) (bool, error) {
	return b.ApproveChatJoinRequest(c.Id, userId)
}

// BanMember Helper method for Bot.BanChatMember
func (c Chat) BanMember(b *Bot, userId int64, opts *BanChatMemberOpts) (bool, error) {
	return b.BanChatMember(c.Id, userId, opts)
}

// BanSenderChat Helper method for Bot.BanChatSenderChat
func (c Chat) BanSenderChat(b *Bot, senderChatId int64) (bool, error) {
	return b.BanChatSenderChat(c.Id, senderChatId)
}

// Copy Helper method for Bot.CopyMessage
func (m Message) Copy(b *Bot, chatId int64, opts *CopyMessageOpts) (*MessageId, error) {
	return b.CopyMessage(chatId, m.Chat.Id, m.MessageId, opts)
}

// CreateInviteLink Helper method for Bot.CreateChatInviteLink
func (c Chat) CreateInviteLink(b *Bot, opts *CreateChatInviteLinkOpts) (*ChatInviteLink, error) {
	return b.CreateChatInviteLink(c.Id, opts)
}

// DeclineJoinRequest Helper method for Bot.DeclineChatJoinRequest
func (c Chat) DeclineJoinRequest(b *Bot, userId int64) (bool, error) {
	return b.DeclineChatJoinRequest(c.Id, userId)
}

// DeletePhoto Helper method for Bot.DeleteChatPhoto
func (c Chat) DeletePhoto(b *Bot) (bool, error) {
	return b.DeleteChatPhoto(c.Id)
}

// DeleteStickerSet Helper method for Bot.DeleteChatStickerSet
func (c Chat) DeleteStickerSet(b *Bot) (bool, error) {
	return b.DeleteChatStickerSet(c.Id)
}

// Delete Helper method for Bot.DeleteMessage
func (m Message) Delete(b *Bot) (bool, error) {
	return b.DeleteMessage(m.Chat.Id, m.MessageId)
}

// EditInviteLink Helper method for Bot.EditChatInviteLink
func (c Chat) EditInviteLink(b *Bot, inviteLink string, opts *EditChatInviteLinkOpts) (*ChatInviteLink, error) {
	return b.EditChatInviteLink(c.Id, inviteLink, opts)
}

// EditCaption Helper method for Bot.EditMessageCaption
func (m Message) EditCaption(b *Bot, opts *EditMessageCaptionOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageCaptionOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageCaption(opts)
}

// EditLiveLocation Helper method for Bot.EditMessageLiveLocation
func (m Message) EditLiveLocation(b *Bot, latitude float64, longitude float64, opts *EditMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageLiveLocation(latitude, longitude, opts)
}

// EditMedia Helper method for Bot.EditMessageMedia
func (m Message) EditMedia(b *Bot, media InputMedia, opts *EditMessageMediaOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageMediaOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageMedia(media, opts)
}

// EditReplyMarkup Helper method for Bot.EditMessageReplyMarkup
func (m Message) EditReplyMarkup(b *Bot, opts *EditMessageReplyMarkupOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageReplyMarkupOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageReplyMarkup(opts)
}

// EditText Helper method for Bot.EditMessageText
func (m Message) EditText(b *Bot, text string, opts *EditMessageTextOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &EditMessageTextOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.EditMessageText(text, opts)
}

// ExportInviteLink Helper method for Bot.ExportChatInviteLink
func (c Chat) ExportInviteLink(b *Bot) (string, error) {
	return b.ExportChatInviteLink(c.Id)
}

// Forward Helper method for Bot.ForwardMessage
func (m Message) Forward(b *Bot, chatId int64, opts *ForwardMessageOpts) (*Message, error) {
	return b.ForwardMessage(chatId, m.Chat.Id, m.MessageId, opts)
}

// Get Helper method for Bot.GetChat
func (c Chat) Get(b *Bot) (*Chat, error) {
	return b.GetChat(c.Id)
}

// GetAdministrators Helper method for Bot.GetChatAdministrators
func (c Chat) GetAdministrators(b *Bot) ([]ChatMember, error) {
	return b.GetChatAdministrators(c.Id)
}

// GetMember Helper method for Bot.GetChatMember
func (c Chat) GetMember(b *Bot, userId int64) (ChatMember, error) {
	return b.GetChatMember(c.Id, userId)
}

// GetMemberCount Helper method for Bot.GetChatMemberCount
func (c Chat) GetMemberCount(b *Bot) (int64, error) {
	return b.GetChatMemberCount(c.Id)
}

// GetMenuButton Helper method for Bot.GetChatMenuButton
func (c Chat) GetMenuButton(b *Bot, opts *GetChatMenuButtonOpts) (MenuButton, error) {
	if opts == nil {
		opts = &GetChatMenuButtonOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = c.Id
	}

	return b.GetChatMenuButton(opts)
}

// Get Helper method for Bot.GetFile
func (f File) Get(b *Bot) (*File, error) {
	return b.GetFile(f.FileId)
}

// GetProfilePhotos Helper method for Bot.GetUserProfilePhotos
func (u User) GetProfilePhotos(b *Bot, opts *GetUserProfilePhotosOpts) (*UserProfilePhotos, error) {
	return b.GetUserProfilePhotos(u.Id, opts)
}

// Leave Helper method for Bot.LeaveChat
func (c Chat) Leave(b *Bot) (bool, error) {
	return b.LeaveChat(c.Id)
}

// PinMessage Helper method for Bot.PinChatMessage
func (c Chat) PinMessage(b *Bot, messageId int64, opts *PinChatMessageOpts) (bool, error) {
	return b.PinChatMessage(c.Id, messageId, opts)
}

// Pin Helper method for Bot.PinChatMessage
func (m Message) Pin(b *Bot, opts *PinChatMessageOpts) (bool, error) {
	return b.PinChatMessage(m.Chat.Id, m.MessageId, opts)
}

// PromoteMember Helper method for Bot.PromoteChatMember
func (c Chat) PromoteMember(b *Bot, userId int64, opts *PromoteChatMemberOpts) (bool, error) {
	return b.PromoteChatMember(c.Id, userId, opts)
}

// RestrictMember Helper method for Bot.RestrictChatMember
func (c Chat) RestrictMember(b *Bot, userId int64, permissions ChatPermissions, opts *RestrictChatMemberOpts) (bool, error) {
	return b.RestrictChatMember(c.Id, userId, permissions, opts)
}

// RevokeInviteLink Helper method for Bot.RevokeChatInviteLink
func (c Chat) RevokeInviteLink(b *Bot, inviteLink string) (*ChatInviteLink, error) {
	return b.RevokeChatInviteLink(c.Id, inviteLink)
}

// SendAction Helper method for Bot.SendChatAction
func (c Chat) SendAction(b *Bot, action string) (bool, error) {
	return b.SendChatAction(c.Id, action)
}

// SetAdministratorCustomTitle Helper method for Bot.SetChatAdministratorCustomTitle
func (c Chat) SetAdministratorCustomTitle(b *Bot, userId int64, customTitle string) (bool, error) {
	return b.SetChatAdministratorCustomTitle(c.Id, userId, customTitle)
}

// SetDescription Helper method for Bot.SetChatDescription
func (c Chat) SetDescription(b *Bot, opts *SetChatDescriptionOpts) (bool, error) {
	return b.SetChatDescription(c.Id, opts)
}

// SetMenuButton Helper method for Bot.SetChatMenuButton
func (c Chat) SetMenuButton(b *Bot, opts *SetChatMenuButtonOpts) (bool, error) {
	if opts == nil {
		opts = &SetChatMenuButtonOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = c.Id
	}

	return b.SetChatMenuButton(opts)
}

// SetPermissions Helper method for Bot.SetChatPermissions
func (c Chat) SetPermissions(b *Bot, permissions ChatPermissions) (bool, error) {
	return b.SetChatPermissions(c.Id, permissions)
}

// SetPhoto Helper method for Bot.SetChatPhoto
func (c Chat) SetPhoto(b *Bot, photo InputFile) (bool, error) {
	return b.SetChatPhoto(c.Id, photo)
}

// SetStickerSet Helper method for Bot.SetChatStickerSet
func (c Chat) SetStickerSet(b *Bot, stickerSetName string) (bool, error) {
	return b.SetChatStickerSet(c.Id, stickerSetName)
}

// SetTitle Helper method for Bot.SetChatTitle
func (c Chat) SetTitle(b *Bot, title string) (bool, error) {
	return b.SetChatTitle(c.Id, title)
}

// StopLiveLocation Helper method for Bot.StopMessageLiveLocation
func (m Message) StopLiveLocation(b *Bot, opts *StopMessageLiveLocationOpts) (*Message, bool, error) {
	if opts == nil {
		opts = &StopMessageLiveLocationOpts{}
	}

	if opts.ChatId == 0 {
		opts.ChatId = m.Chat.Id
	}
	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.StopMessageLiveLocation(opts)
}

// UnbanMember Helper method for Bot.UnbanChatMember
func (c Chat) UnbanMember(b *Bot, userId int64, opts *UnbanChatMemberOpts) (bool, error) {
	return b.UnbanChatMember(c.Id, userId, opts)
}

// UnbanSenderChat Helper method for Bot.UnbanChatSenderChat
func (c Chat) UnbanSenderChat(b *Bot, senderChatId int64) (bool, error) {
	return b.UnbanChatSenderChat(c.Id, senderChatId)
}

// UnpinAllMessages Helper method for Bot.UnpinAllChatMessages
func (c Chat) UnpinAllMessages(b *Bot) (bool, error) {
	return b.UnpinAllChatMessages(c.Id)
}

// UnpinMessage Helper method for Bot.UnpinChatMessage
func (c Chat) UnpinMessage(b *Bot, opts *UnpinChatMessageOpts) (bool, error) {
	return b.UnpinChatMessage(c.Id, opts)
}

// Unpin Helper method for Bot.UnpinChatMessage
func (m Message) Unpin(b *Bot, opts *UnpinChatMessageOpts) (bool, error) {
	if opts == nil {
		opts = &UnpinChatMessageOpts{}
	}

	if opts.MessageId == 0 {
		opts.MessageId = m.MessageId
	}

	return b.UnpinChatMessage(m.Chat.Id, opts)
}
