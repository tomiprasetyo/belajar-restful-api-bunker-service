package helper

import "database/sql"

// method untuk pengecekan transaksi di database
func CommitOrRollback(tx *sql.Tx) {
	// cek apakah terjadi error atau tidak menggunakan recover
	err := recover()
	// jika ada error
	if err != nil {
		// rollback transaksinya
		errorRollback := tx.Rollback()
		// cek jika error
		PanicIfError(errorRollback)

		panic(err)

		//jika tidak ada error
	} else {
		// lakukan commit transaksi
		errorCommit := tx.Commit()
		// cek jika error
		PanicIfError(errorCommit)
	}
}
