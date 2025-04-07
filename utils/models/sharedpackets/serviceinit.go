package sharedpackets

import (
	"github.com/TechnoDiktator/ginTutorial/lib/pg"
	"github.com/jmoiron/sqlx"
)

type LocalServiceInit struct {
	PgConnect *sqlx.DB
	PgService pg.LocalPgService
	// RedisConnect  *redis.Client
	// DefaultGate   *localdqrdbmodels.Gate
	// RedisService  *redisservice.RedisService
	//SharedService localDqrService.LocalDqrService
	//HasuraService gqlsubscription.GqlService

}
