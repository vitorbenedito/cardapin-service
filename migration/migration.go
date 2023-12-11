package migration

import (
	"log"

	"cardap.in/lambda/db"
	"cardap.in/lambda/model"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func AutoMigrate() {

	m := gormigrate.New(db.DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				tx.AutoMigrate(&model.AdditionalItem{})
				tx.AutoMigrate(&model.Address{})
				tx.AutoMigrate(&model.Category{})
				tx.AutoMigrate(&model.Company{})
				tx.AutoMigrate(&model.Menu{})
				tx.AutoMigrate(&model.OpeningHours{})
				tx.AutoMigrate(&model.PaymentType{})
				tx.AutoMigrate(&model.Product{})
				tx.AutoMigrate(&model.Section{})
				tx.AutoMigrate(&model.Client{})
				tx.AutoMigrate(&model.User{})
				tx.AutoMigrate(&model.Table{})
				tx.AutoMigrate(&model.AdditionalItemsGroup{})

				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "2",
			Migrate: func(tx *gorm.DB) error {
				return tx.Exec("ALTER TABLE company_section ADD CONSTRAINT company_section_s_fk FOREIGN KEY (section_id) REFERENCES section(ID)").Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Exec("ALTER TABLE company_section DROP CONSTRAINT company_section_fk").Error
			},
		},
		{

			ID: "3",
			Migrate: func(tx *gorm.DB) error {
				return tx.Exec("ALTER TABLE company_section ADD CONSTRAINT company_section_c_fk FOREIGN KEY (company_id) REFERENCES company(ID)").Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Exec("ALTER TABLE company_section DROP CONSTRAINT company_section_c_fk").Error
			},
		},
		{

			ID: "4",
			Migrate: func(tx *gorm.DB) error {
				return tx.Exec("INSERT INTO cardapin_user(created_at, updated_at, deleted_at, \"name\", login, email, \"password\") VALUES('2020-07-04 23:31:53.756', '2020-07-04 23:31:53.734', NULL, 'Cardapin Admin', 'cardapin-admin', 'falecom@cardap.in', 'edbf177c911e9bb38bbc9ab1102a8e81')").Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Exec("DELETE cardapin_user where id = 1").Error
			},
		},
		{

			ID: "5",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec("INSERT INTO section(name) VALUES ('Alimentação')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Lanches')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Vestuário')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Artesanato')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Hamburguers')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Bebidas')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Pizzas')").Error
				err = tx.Exec("INSERT INTO section(name) VALUES ('Lanches Árabes')").Error

				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Mastercard Crédito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Mastercard Débito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Visa Crédito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Visa Débito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Elo Crédito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Elo Débito')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Alelo VR - Voucher')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Alelo VA - Voucher')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Sodexo VR - Voucher')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Sodexo VA - Voucher')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Dinheiro')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('Picpay')").Error
				err = tx.Exec("INSERT INTO payment_type(name) VALUES ('MercadoPago')").Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Exec("DELETE from section").Error
				err = tx.Exec("DELETE from payment_type").Error
				return err
			},
		},
		{

			ID: "6",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec("ALTER TABLE category ADD CONSTRAINT category_menu_fk FOREIGN KEY (menu_id) REFERENCES menu(id) ON DELETE CASCADE;").Error
				err = tx.Exec("ALTER TABLE product ADD CONSTRAINT category_product_fk FOREIGN KEY (category_id) REFERENCES category(id) ON DELETE CASCADE;").Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{

			ID: "7",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec("ALTER TABLE cardapin_table ADD CONSTRAINT table_company_fk FOREIGN KEY (company_id) REFERENCES company(id);").Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{

			ID: "8",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec("ALTER TABLE additional_items_group ADD CONSTRAINT group_company_fk FOREIGN KEY (company_id) REFERENCES company(id);").Error
				err = tx.Exec("ALTER TABLE additional_item ADD CONSTRAINT additional_item_group_fk FOREIGN KEY (additional_items_group_id) REFERENCES additional_items_group(id) ON DELETE CASCADE;").Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{

			ID: "9",
			Migrate: func(tx *gorm.DB) error {
				err := tx.Exec("ALTER TABLE address ADD CONSTRAINT address_company_fk FOREIGN KEY (company_id) REFERENCES company(id);").Error
				err = tx.Exec("ALTER TABLE company_payment_type ADD CONSTRAINT company_payment_type_company_fk FOREIGN KEY (company_id) REFERENCES company(id);").Error
				err = tx.Exec("ALTER TABLE company_payment_type ADD CONSTRAINT company_payment_type_payment_fk FOREIGN KEY (payment_type_id) REFERENCES payment_type(id);").Error
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Exec("ALTER TABLE address DROP CONSTRAINT address_company_fk").Error
				err = tx.Exec("ALTER TABLE company_payment_type DROP CONSTRAINT company_payment_type_company_fk").Error
				err = tx.Exec("ALTER TABLE company_payment_type DROP CONSTRAINT company_payment_type_address_fk").Error
				return err
			},
		},
	})

	err := m.Migrate()
	if err == nil {
		log.Printf("Migration did run successfully")
	} else {
		log.Printf("Could not migrate: %v", err)
	}
}
