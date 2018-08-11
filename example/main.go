package main

import (
	"context"

	"code.jogchat.internal/go-schemaless/utils"
	"github.com/satori/go.uuid"
	"encoding/json"
	"fmt"
	"code.jogchat.internal/go-schemaless/models"
	"time"
	"code.jogchat.internal/go-schemaless"
)


func newBusiness(id uuid.UUID, colKey string, domain string, name string) models.Cell {
	rowKey := utils.NewUUID().Bytes()
	refKey := time.Now().UnixNano()
	blob, err := json.Marshal(map[string]interface{} {
		"id": id,
		"domain": domain,
		"name": name,
	})
	utils.CheckErr(err)
	return models.Cell{RowKey: rowKey, ColumnName: colKey, RefKey: refKey, Body: blob}
}

// run this to make sure nothing breaks
func main() {
	dataStore := schemaless.InitDataStore()
	defer dataStore.Destroy(context.TODO())

	UIUC := newBusiness(utils.NewUUID(), "schools", "illinois.edu", "UIUC")
	err := dataStore.PutCell(context.TODO(), UIUC.RowKey, UIUC.ColumnName, UIUC.RefKey, UIUC)
	utils.CheckErr(err)

	CMU := newBusiness(utils.NewUUID(), "schools", "andrew.cmu.edu", "CMU")
	err = dataStore.PutCell(context.TODO(), CMU.RowKey, CMU.ColumnName, CMU.RefKey, CMU)

	Sift := newBusiness(utils.NewUUID(), "companies", "siftscience.com", "Sift Science")
	err = dataStore.PutCell(context.TODO(), Sift.RowKey, Sift.ColumnName, Sift.RefKey, Sift)
	utils.CheckErr(err)

	Yahoo := newBusiness(utils.NewUUID(), "companies", "yahoo-inc.com", "Yahoo!")
	err = dataStore.PutCell(context.TODO(), Yahoo.RowKey, Yahoo.ColumnName, Yahoo.RefKey, Yahoo)
	utils.CheckErr(err)

	cells, _, err := dataStore.GetCellsByFieldLatest(context.TODO(), "schools", "domain", "illinois.edu", "=")
	utils.CheckErr(err)
	for _, cell := range cells {
		fmt.Println(cell.String())
	}

	cells, _, err = dataStore.GetCellsByColumnLatest(context.TODO(), "companies")
	utils.CheckErr(err)
	for _, cell := range cells {
		fmt.Println(cell.String())
	}
}
