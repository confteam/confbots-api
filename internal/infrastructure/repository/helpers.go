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

func ptrPgInt8(v *int64) pgtype.Int8 {
	if v == nil {
		return pgtype.Int8{Valid: false}
	}
	return pgtype.Int8{Int64: *v, Valid: true}
}

func ptrPgText(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *v, Valid: true}
}
