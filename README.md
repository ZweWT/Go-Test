# GO TEST (Todo app with auth system)

## MakeFile
Check my makefile for operations, setting up database, running the migrations, starting the server

```bash
$ make help
```

## .envrc
Update the variables in the .envrc file 
```
export GOTEST_DB_DSN='postgres://gotest:gotest@localhost/gotest'

export JWT_SECRET=pei3einoh0Beem6uM6Ungohn2heiv5lah1ael4joopie5JaigeikoozaoTew2Eh6
```
** GO_TEST_DB_DSN='postgres://username:password@localhost/dbname'

** Your secret key should be a cryptographically secure random string with an
underlying entropy of at least 32 bytes (256 bits).

