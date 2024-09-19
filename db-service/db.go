package main

import (
    "context"
    "os"

    "github.com/jackc/pgx/v5"
)

type DBService struct{}

func NewDBService() *DBService {
    return &DBService{}
}

func (db *DBService) SaveOrder(ctx context.Context, userId int64, orderText string) error {
    conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
    if err != nil {
        return err
    }
    defer conn.Close(ctx)

    _, err = conn.Exec(ctx, "INSERT INTO orders (user_id, order_text) VALUES ($1, $2)", userId, orderText)
    return err
}
