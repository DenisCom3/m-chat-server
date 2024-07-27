package message

import (
	"context"

	"github.com/DenisCom3/m-chat-server/internal/client/db"
	"github.com/DenisCom3/m-chat-server/internal/model"
	"github.com/DenisCom3/m-chat-server/internal/repository"
	"github.com/Masterminds/squirrel"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

var _ repository.MessageRepository = (*repo)(nil)

const (
	tableName = "messages"
	message_ID = "id"
	message_Text = "text"
	message_ChatID = "chat_id"
	message_SenderID = "user_id"
	message_created = "created_at"
	message_updated = "updated_at"
	// message_IsRead = "is_read"
	// message_IsEdited = "is_edited"
	// message_IsDeleted = "is_deleted"
	// message_IsSent = "is_sent"
)


func (r *repo) Create(ctx context.Context, message model.CreateMessage) (*model.Message, error) {

	sql, args, err := psql.Insert(tableName).
		Columns(message_Text, message_ChatID, message_SenderID).
		Values(message.Text, message.ChatId, message.UserId).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}
	q := db.Query{
		QueryRaw: sql,
		Name: "messageRepo.Create",
	}
	var m model.Message
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}