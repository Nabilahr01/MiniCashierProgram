package main

import (
	"fmt"
	"strings"
)

type Produk struct {
	ID    int
	Nama  string
	Harga float64
}
type Pelanggan struct {
	ID   int
	Nama string
}
type Transaksi struct {
	ID          int
	IDPelanggan int
	IDProduks   []int
	Total       float64
}

// Urut harga besar ke kecil (selection sort)
func sortProduk(produk []Produk) {
	n := len(produk)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if produk[j].Harga < produk[idxMin].Harga {
				idxMin = j
			}
		}
		produk[i], produk[idxMin] = produk[idxMin], produk[i]
	}
}

// Urut total belanja besar ke kecil (insertion sort)
func sortTransaksi(transaksi []Transaksi) {
	for i := 1; i < len(transaksi); i++ {
		kunci := transaksi[i]
		j := i - 1
		for j >= 0 && transaksi[j].Total < kunci.Total {
			transaksi[j+1] = transaksi[j]
			j--
		}
		transaksi[j+1] = kunci
	}
}

// Cari produk dengan nama / unsur nama
func cariProdukNama(produk []Produk, keyword string) []Produk {
	var hasil []Produk
	for i := 0; i < len(produk); i++ {
		if strings.Contains(strings.ToLower(produk[i].Nama), strings.ToLower(keyword)) {
			hasil = append(hasil, produk[i])
		}
	}
	return hasil
}

// Cari transaksi berdasarkan ID
func cariTransaksiID(transaksi []Transaksi, id int) *Transaksi {
	for i := 0; i < len(transaksi); i++ {
		if transaksi[i].ID == id {
			return &transaksi[i]
		}
	}
	return nil
}

// Ambil nama produk berdasarkan ID
func dataProdukByID(produk []Produk, id int) string {
	for i := 0; i < len(produk); i++ {
		if produk[i].ID == id {
			return produk[i].Nama
		}
	}
	return "Produk tidak ditemukan"
}

// Ambil nama pelanggan berdasarkan ID
func dataPelangganByID(pelanggan []Pelanggan, id int) string {
	for i := 0; i < len(pelanggan); i++ {
		if pelanggan[i].ID == id {
			return pelanggan[i].Nama
		}
	}
	return "Pelanggan tidak ditemukan"
}
