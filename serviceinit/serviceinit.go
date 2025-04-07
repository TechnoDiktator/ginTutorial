package serviceinit

import (
	"github.com/TechnoDiktator/ginTutorial/lib/pg"

	"github.com/TechnoDiktator/ginTutorial/utils/models/sharedpackets"
	"github.com/sirupsen/logrus"
)

func LocalServerInit()(*sharedpackets.LocalServiceInit, error)  {

	logAndReturnError := func(err error, context string) (*sharedpackets.LocalServiceInit, error) {
		logrus.Warnf("Error while %s: %v", context)
		return nil, err
	}

	// Initialize the database connection
	db, err := pg.NewLocalPGInstance()
	if err != nil {
		return logAndReturnError(err, "initializing database connection")

	}

	pgService := pg.NewLocalPGService(db)
	if pgService == nil {
		return logAndReturnError(err, "initializing pg service")
	}

	return &sharedpackets.LocalServiceInit{
		PgConnect: db,
		PgService: pgService,
	}, nil

}
