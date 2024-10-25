package database

import (
	"context"
	"fmt"

	"github.com/Toorreess/laWiki/wiki-service/pkg/database/clients/firestore"
)

type Connection struct {
	Client interface{}
	Type   string
	Ctx    context.Context
}

type DBClient interface {
	Get()
	Create(index string, entity interface{}) (map[string]interface{}, error)
	Update()
	Delete()
	List()
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

func (conn *Connection) Create(index string, entity interface{}) (map[string]interface{}, error) {
	if conn.Client == nil {
		return make(map[string]interface{}), fmt.Errorf("no client found. Please, init Connection before.")
	}
	return conn.Client.(DBClient).Create(index, entity)
}

func (conn *Connection) Read(index string, id string) (map[string]interface{}, error) {
	return nil, nil
}

func (conn *Connection) Update(index, id string, entity interface{}) (map[string]interface{}, error) {
	return nil, nil
}

func (conn *Connection) Delete(index, id string) error {
	return nil
}

func (conn *Connection) List(index, query string, limit, offset int, orderBy, order string, entity interface{}) ([]map[string]interface{}, error) {
	return nil, nil
}
