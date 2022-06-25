CREATE TABLE bunker_services (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    guid VARCHAR(55) UNIQUE NOT NULL,
    no_so VARCHAR(255) NOT NULL,
    nama_perusahaan VARCHAR(255) NOT NULL,
    nama_kapal VARCHAR(255) NOT NULL,
    nama_produk VARCHAR(255) NOT NULL,
    jumlah_pengisian INTEGER(55) NOT NULL,
    pelabuhan VARCHAR(255) NOT NULL,
    nopol_truk VARCHAR(255) NOT NULL,
    nama_operator VARCHAR(255) NOT NULL,
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME
);