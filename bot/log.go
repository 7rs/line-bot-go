package bot

import (
	"log"

	"github.com/comail/colog"
)

func SetColog() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Lshortfile,
	})
	colog.Register()
}
