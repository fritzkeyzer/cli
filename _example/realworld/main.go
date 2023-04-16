package main

import (
	"log"

	"github.com/fritzkeyzer/cli"
)

// This example shows a real world example of a database management cli tool.
// Functional code is commented out to ensure this example will compile.
// Notice that for each command the flags are specified as required.
// This ensures that the flags are loaded before the command is executed.
// If any required flags are missing, the cli will print the docs, a list of errors and exit.

var DBConnFlag = &cli.StringFlag{
	Name:        "db-conn",
	EnvVar:      "DB_CONN",
	Description: "Database connection string",
}

type GCPCreds struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project_id"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
}

var GCPCredsFlag = &cli.JSONFlag[GCPCreds]{
	Name:        "gcp-creds",
	EnvVar:      "GCP_CREDS",
	Description: "GCP credentials in JSON format",
}

var BackupBucketFlag = &cli.StringFlag{
	Name:        "bucket",
	EnvVar:      "DB_BACKUP_BUCKET",
	Description: "GCS bucket to backup and restore from",
}

func main() {
	app := &cli.App{
		Name: "db-cli",
		Description: `cli tool for managing the application database.
This tool is used in a CI/CD pipeline as well as for local development.
Flags are provided with environment variables, but can be overridden with command line flags.
`,
		SubCmds: []cli.Cmd{
			backupCmd,
			restoreCmd,
			dropCmd,
			migrateCmd,
		},
	}

	app.Run()
}

var backupCmd = cli.Cmd{
	Name:        "backup",
	Description: "Create a database backup and upload it to GCS",
	ReqFlags: []cli.Flag{
		DBConnFlag,
		GCPCredsFlag,
		BackupBucketFlag,
	},
	Action: func(args map[string]string) {
		log.Println("Perform database backup")

		// database.BackupToGCS(context.Background(),
		// 	DBConnFlag.Value,
		// 	GCPCredsFlag.Value,
		// 	BackupBucketFlag.Value,
		// )
	},
}

var restoreCmd = cli.Cmd{
	Name:        "restore",
	Description: "Restore database from backup from GCS",
	ReqFlags: []cli.Flag{
		DBConnFlag,
		GCPCredsFlag,
		BackupBucketFlag,
	},
	Action: func(args map[string]string) {
		log.Println("Perform database restore")

		// if err := database.RestoreFromGCS(context.Background(),
		// 	DBConnFlag.Value,
		// 	GCPCredsFlag.Value,
		// 	BackupBucketFlag.Value,
		// ); err != nil {
		// 	log.Fatal("ERROR:", err)
		// }
	},
}

var dropCmd = cli.Cmd{
	Name:        "drop",
	Description: "drop and recreate public schema - WARNING!!! ALL DATA WILL BE LOST!",
	ReqFlags: []cli.Flag{
		DBConnFlag,
	},
	Action: func(args map[string]string) {
		log.Println("Perform database drop")

		// db, err := database.NewConn(DBConnFlag.Value)
		// if err != nil {
		// 	log.Fatal("ERROR: new database connection pool:", err)
		// }
		//
		// log.Println("Dropping db")
		// if err := database.Drop(context.Background(), db); err != nil {
		// 	log.Fatal("ERROR: drop database:", err)
		// }
		//
		// log.Println("Creating schema")
		// if err := database.CreateSchema(context.Background(), db); err != nil {
		// 	log.Fatal("ERROR: create schema:", err)
		// }
		//
		// log.Println("Adding extensions")
		// if err := database.AddExtensions(context.Background(), db); err != nil {
		// 	log.Fatal("ERROR: add extensions:", err)
		// }
	},
}

var migrateCmd = cli.Cmd{
	Name:        "migrate",
	Description: "run all migrations",
	ReqFlags: []cli.Flag{
		DBConnFlag,
	},
	Action: func(args map[string]string) {
		log.Println("Perform database migration")

		// db, err := database.NewConn(DBConnFlag.Value)
		// if err != nil {
		// 	log.Fatal("ERROR: new database connection pool:", err)
		// }
		//
		// if err := migration.RunAll(context.Background(), db); err != nil {
		// 	log.Fatal("ERROR: run migrations:", err)
		// }
	},
}
