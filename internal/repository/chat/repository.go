package chat

import (
	"context"
	"errors"

	"github.com/DenisCom3/m-chat-server/internal/client/db"
	"github.com/DenisCom3/m-chat-server/internal/model"
	"github.com/DenisCom3/m-chat-server/internal/repository"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

var _ repository.ChatRepository = (*repo)(nil)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

const (
	tableChats = "chats"
	chat_ID = "id"
	chat_Name = "name"
	chat_IsActive = "is_active"
)

const (
	tableChatUsers = "chat_users"
	chatUser_ChatID = "chat_id"
	chatUser_UserID = "user_id"
)


type repo struct {
	db db.Client
}

func New(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (r *repo) Create(ctx context.Context, chat model.CreateChat) (*model.Chat, error) {


	sql, args, err := psql.Insert(tableChats).
		Columns(chat_Name).
		Values(chat.Name).
		Suffix("RETURNING *").
		ToSql()

	if err != nil {
		return nil, err
	}

	q := db.Query{
		QueryRaw: sql,
		Name: "chatRepo.Create",
	}

	var c model.Chat
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&c)
	if err != nil {
		return nil, err
	}

	for _, user := range chat.Users {
		err = r.addUser(ctx, c.ID, user.ID)
			if err != nil {
			return nil, err
		}
	}

	return &c, nil
	
}

func (r *repo) addUser(ctx context.Context, chatID, userID int64) error {

	sql, args, err := psql.Insert(tableChatUsers).
		Columns(chatUser_ChatID, chatUser_UserID).
		Values(chatID, userID).
		Suffix("RETURNING *").
		ToSql()


	if err != nil {
		return err
	}

	q := db.Query{
		QueryRaw: sql,
		Name: "chatRepo.addUser",
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan()

	return err

}

func (r *repo) CheckUserInChat(ctx context.Context, chatID, userID int64) (bool, error) {
	sql, args, err := psql.Select(chatUser_ChatID).
		From(tableChatUsers).
		Where(squirrel.Eq{chatUser_ChatID: chatID, chatUser_UserID: userID}).
		Limit(1).
		ToSql()

	if err != nil {
		return false, err
	}

	q := db.Query{
		QueryRaw: sql,
		Name: "chatRepo.CheckUserInChat",
	}

	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan()
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}

		return false, err
	}

	return true, nil

}

