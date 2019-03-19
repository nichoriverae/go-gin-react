package migrate

import (
	"github.com/jinzhu/gorm"
	gormigrate "gopkg.in/gormigrate.v1"

	"go-gin-react/server/db/models"
)

func main() {

}

// Start initializes the migrations
func Start(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				return tx.CreateTable(&models.BlogPost{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable(&models.BlogPost{}).Error
			},
		},
	})
	return m.Migrate()
}
