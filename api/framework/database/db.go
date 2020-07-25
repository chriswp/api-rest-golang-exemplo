package database

import (
	"api/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstace := NewDb()
	dbInstace.Env = "test"
	dbInstace.DbTypeTest = "sqlite3"
	dbInstace.DsnTest = ":memory:"
	dbInstace.AutoMigrateDb = true
	dbInstace.Debug = true

	connection, err := dbInstace.Connect()

	if err != nil {
		log.Fatalf("Test db erro %v", err)
	}

	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(d.DbType, d.Dsn)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Categoria{}, &domain.Produto{}, &domain.ProdutoDigital{}, &domain.Cliente{}, &domain.Compra{}, &domain.CompraItem{})
		d.Db.AutoMigrate(&domain.Venda{}, &domain.VendaItem{})
		d.Db.Model(domain.Produto{}).AddForeignKey("categoria_id", "categorias(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.CompraItem{}).AddForeignKey("compra_id", "compras(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.CompraItem{}).AddForeignKey("produto_id", "produtos(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.ProdutoDigital{}).AddForeignKey("produto_id", "produtos(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.Venda{}).AddForeignKey("cliente_id", "clientes(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.VendaItem{}).AddForeignKey("produto_id", "produtos(id)", "NO ACTION", "NO ACTION")
		d.Db.Model(domain.VendaItem{}).AddForeignKey("venda_id", "vendas(id)", "NO ACTION", "NO ACTION")
	}

	return d.Db, nil
}
