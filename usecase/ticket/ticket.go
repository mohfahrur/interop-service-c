package ticket

import (
	"log"

	googleD "github.com/mohfahrur/interop-service-c/domain/google"
	"github.com/mohfahrur/interop-service-c/entity"
)

type TicketAgent interface {
	UpdateGoogle(user string, item string) (err error)
}

type TicketUsecase struct {
	googleDomain googleD.GoogleDomain
}

func NewTicketUsecase(
	googleD googleD.GoogleDomain) *TicketUsecase {

	return &TicketUsecase{
		googleDomain: googleD}
}

func (uc *TicketUsecase) UpdateSheet(req entity.UpdateSheetRequest) (err error) {

	err = uc.googleDomain.UpdateSheetPenjualan(req)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
