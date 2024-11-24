/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"database/sql"
	"fmt"
	dbInfra "go-hexagonal/adapters/db"
	"go-hexagonal/application"
	"os"

	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite
	"github.com/spf13/cobra"
)

var cfgFile string

// Declaração das variáveis globais para a conexão e o serviço de produtos
var db *sql.DB
var productDb *dbInfra.ProductDb
var productService application.ProductService

// rootCmd representa o base command
var rootCmd = &cobra.Command{
	Use:   "go-hexagonal",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your application.
Cobra is a CLI library for Go that empowers applications.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adiciona todos os subcomandos ao comando root
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Abre a conexão com o banco de dados SQLite
	var err error
	db, err = sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		os.Exit(1)
	}

	// Valida a conexão com o banco
	err = db.Ping()
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		os.Exit(1)
	}

	// Cria a tabela products se não existir
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS products (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		price REAL,
		status TEXT NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Printf("Error creating table: %v\n", err)
		os.Exit(1)
	}

	// Inicializa o banco de dados e o serviço de produtos
	productDb = dbInfra.NewProductDb(db)
	productService = application.ProductService{Persistence: productDb}

	// Adiciona o comando CLI
	rootCmd.AddCommand(cliCmd)

	// Configura os flags para o comando CLI
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
