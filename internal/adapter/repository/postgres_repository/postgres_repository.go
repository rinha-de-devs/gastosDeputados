package postgres_repository

import (
	"database/sql"
	"deputySpending/internal/domain"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type postgresDB struct {
}

func New() *postgresDB {
	return &postgresDB{}
}

func (repo *postgresDB) SaveDeputy(deputy domain.Deputy) (domain.Deputy, error) {

	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO public.deputy 
	(id, nome, partido, estado, cota, verba_de_gabinete_disponivel, porcentagem_disponivel, verba_de_gabinete_gasto, porcentagem_gasto)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);
	`

	err := db.QueryRow(sqlStatement, deputy.ID, deputy.Nome, deputy.Partido, deputy.Estado, deputy.Cota, deputy.VerbaDeGabineteDisponivel, deputy.PorcentagemDisponivel, deputy.VerbaDeGabineteGasto, deputy.PorcentagemGasto).Scan()
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
		return domain.Deputy{}, err
	}

	return deputy, nil

}

func createConnection() *sql.DB {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error to load env file")
	}

	var (
		host   = os.Getenv("DATABASE_HOST")
		port   = os.Getenv("DATABASE_PORT")
		user   = os.Getenv("DATABASE_USER")
		pass   = os.Getenv("DATABASE_PASSWORD")
		dbname = os.Getenv("DATABASE_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected")

	return db
}

// func insert(deputy domain.Deputy, db *sql.DB) (domain.Deputy, error) {
// 	sqlStatement := `INSERT INTO public.deputy
// 	(id, nome, partido, estado, cota, verba_de_gabinete_disponivel, porcentagem_disponivel, verba_de_gabinete_gasto, porcentagem_gasto)
// 	VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9);
// 	`

// 	err := db.QueryRow(sqlStatement, deputy.ID, deputy.Nome, deputy.Partido, deputy.Estado, deputy.Cota, deputy.VerbaDeGabineteDisponivel, deputy.PorcentagemDisponivel, deputy.VerbaDeGabineteGasto, deputy.PorcentagemGasto).Scan()
// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 		return domain.Deputy{}, err
// 	}

// 	return deputy, nil
// }

// func update(deputy domain.Deputy, db *sql.DB) (domain.Deputy, error) {
// 	sqlStatement := `UPDATE public.deputy
// 	SET nome=$2, partido=$3, estado=$4, cota=$5, verba_de_gabinete_disponivel=$6, porcentagem_disponivel=$7, verba_de_gabinete_gasto=$8, porcentagem_gasto=$9
// 	WHERE id=$1
// 	`

// 	result, err := db.Exec(sqlStatement, deputy.ID, deputy.Nome, deputy.Partido, deputy.Estado, deputy.Cota, deputy.VerbaDeGabineteDisponivel, deputy.PorcentagemDisponivel, deputy.VerbaDeGabineteGasto, deputy.PorcentagemGasto)
// 	if err != nil {
// 		log.Fatalf("Unable to execute the query. %v", err)
// 		return domain.Deputy{}, err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		log.Fatalf("Error while checking the affected rows. %v", err)
// 		return domain.Deputy{}, err
// 	}

// 	fmt.Printf("Total rows/record affected %v", rowsAffected)

// 	return deputy, nil
// }
