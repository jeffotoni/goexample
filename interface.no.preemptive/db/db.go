//db.go 
package db

import "database/sql"

type Store struct { 
   db *sql.DB 
}

func NewDB() *Store {...} // função para inicializar DB

func (s *Store) Insert (item interface {}) error {...} // inserir item

func (s *Store) Get (id int) error {...} // get item by id
