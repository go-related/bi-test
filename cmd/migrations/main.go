package main

import (
	"entdemo/blockchain/migration"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	RunMigrations2()

}

func RunMigrations2() {

	// deploy contracts
	ethUrl := "http://localhost:8545"
	adminPrivateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"

	err := migration.DeployContract(ethUrl, adminPrivateKeyHex)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info("finished running migrations")
}
