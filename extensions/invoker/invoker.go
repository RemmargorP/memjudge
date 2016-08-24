package invoker

import (
	"github.com/RemmargorP/mjudge/interfaces"
	"github.com/RemmargorP/mjudge/models"

	_ "github.com/go-mangos/mangos"
)

type Invoker struct {
	LanguagesRepo interfaces.LanguagesRepo
}

func (inv *Invoker) Process() {
	lang, _ := inv.LanguagesRepo.GetById(models.Id("lel"))
	lang.Process(&models.Submission{}, &models.Problem{})
}
