package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func ptrInt32(n pgtype.Int4) *int32 {
	if n.Valid {
		return &n.Int32
	}

	return nil
}

func ptrString(n pgtype.Text) *string {
	if n.Valid {
		return &n.String
	}

	return nil
}
