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

func ptrInt64(n pgtype.Int8) *int64 {
	if n.Valid {
		return &n.Int64
	}

	return nil
}

func ptrString(n pgtype.Text) *string {
	if n.Valid {
		return &n.String
	}

	return nil
}
