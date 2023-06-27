package google

import (
	"context"
	"fmt"
	"log"

	"github.com/mohfahrur/interop-service-c/entity"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

type GoogleAgent interface {
	UpdateSheetPenjualan(data entity.UpdateSheetRequest) (err error)
}

type GoogleDomain struct {
	SheetService  *sheets.Service
	SpreadsheetID string
}

func NewGoogleDomain(credentialsFile []byte, SpreadsheetID string) *GoogleDomain {
	ctx := context.Background()
	srv, err := sheets.NewService(ctx,
		option.WithScopes(sheets.SpreadsheetsScope),
		option.WithCredentialsJSON(credentialsFile))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	return &GoogleDomain{
		SheetService:  srv,
		SpreadsheetID: SpreadsheetID,
	}
}

func (d *GoogleDomain) UpdateSheetPenjualan(data entity.UpdateSheetRequest) (err error) {

	// Create the value range to update
	valueRange := &sheets.ValueRange{
		Values: [][]interface{}{
			{data.User, data.Email, data.Hp, data.Item},
		},
	}
	lastRow, err := d.SheetService.Spreadsheets.Values.Get(d.SpreadsheetID, "PMG Penjualan!A:A").Do()
	if err != nil {
		log.Printf("Failed to get last row: %v", err)
		return
	}
	rowIndex := len(lastRow.Values) + 1

	targetRange := fmt.Sprintf("PMG Penjualan!A%d:D%d", rowIndex, rowIndex)

	_, err = d.SheetService.Spreadsheets.Values.Update(d.SpreadsheetID, targetRange, valueRange).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Printf("Failed to update values: %v", err)
		return
	}

	fmt.Println("Values updated successfully!")
	return
}
