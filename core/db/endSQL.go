package db

type endSQL interface {
	updateSQL
	insertSQL
	findSQL
	deleteSQL
}
