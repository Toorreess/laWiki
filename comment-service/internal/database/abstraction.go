package database

import (
	"context"
	"fmt"

	"github.com/Toorreess/laWiki/comment-service/internal/database/clients/firestore"
)

type Connection struct {
	Client interface{}
	Type   string
	Ctx    context.Context
}

type DBClient interface {
	Get(index, id string, entity interface{}) (map[string]interface{}, error)
	Create(index string, entity interface{}) (map[string]interface{}, error)
	Update(index, id string, entity interface{}, updates map[string]interface{}) (map[string]interface{}, error)
	Delete(index, id string) error
	List(index string, query map[string]string, limit, offset int, orderBy, order string, entity interface{}) ([]map[string]interface{}, error)
	Close() error
}

func NewDBClient(dbType, user, passwd, addr, dbName string) (*Connection, error) {
	var conn Connection
	ctx := context.Background()
	if dbType == "firestore" {
		fsClient := firestore.Client{Project: addr}

		if err := fsClient.Init(ctx); err != nil {
			return nil, err
		}
		conn.Client = fsClient
	}
	conn.Type = dbType
	conn.Ctx = ctx
	return &conn, nil
}

func (conn *Connection) Close() error {
	return conn.Client.(DBClient).Close()
}

func (conn *Connection) Create(index string, entity interface{}) (map[string]interface{}, error) {
	if conn.Client == nil {
		return make(map[string]interface{}), fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).Create(index, entity)
}

func (conn *Connection) Get(index string, id string, entity interface{}) (map[string]interface{}, error) {
	if conn.Client == nil {
		return make(map[string]interface{}), fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).Get(index, id, entity)
}

func (conn *Connection) Update(index, id string, entity interface{}, updates map[string]interface{}) (map[string]interface{}, error) {
	if conn.Client == nil {
		return make(map[string]interface{}), fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).Update(index, id, entity, updates)
}

func (conn *Connection) Delete(index, id string) error {
	if conn.Client == nil {
		return fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).Delete(index, id)
}

func (conn *Connection) List(index string, query map[string]string, limit, offset int, orderBy, order string, entity interface{}) ([]map[string]interface{}, error) {
	if conn.Client == nil {
		return nil, fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).List(index, query, limit, offset, orderBy, order, entity)
}
