package models

import "time"

type Mahasiswa struct {
	ID        int        `json:"id"`
	Nama      string     `json:"nama"`
	Password  string     `json:"password,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

var MahasiswaQuery = `
CREATE TABLE IF NOT EXISTS mahasiswa (
	id INT AUTO_INCREMENT PRIMARY KEY,
	nama VARCHAR(100) NOT NULL UNIQUE,
	password VARCHAR(255) NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP NULL DEFAULT NULL
);`