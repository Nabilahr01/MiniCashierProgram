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

// Hitung total belanja
func hitungTotal(produk []Produk, idProduks []int) float64 {
	total := 0.0
	for i := 0; i < len(idProduks); i++ {
		id := idProduks[i]
		for j := 0; j < len(produk); j++ {
			if id == produk[j].ID {
				total += produk[j].Harga
			}
		}
	}
	return total
}

func main() {
	var jumlahProduk, jumlahPelanggan, jumlahTransaksi int
	fmt.Print("Masukkan jumlah produk: ")
	fmt.Scan(&jumlahProduk)
	daftarProduk := make([]Produk, jumlahProduk)
	for i := 0; i < jumlahProduk; i++ {
		fmt.Printf("Produk %d:\n", i+1)
		daftarProduk[i].ID = i + 1
		fmt.Print("  Nama produk: ")
		fmt.Scan(&daftarProduk[i].Nama)
		fmt.Print("  Harga produk: ")
		fmt.Scan(&daftarProduk[i].Harga)
	}

	fmt.Print("\nMasukkan jumlah pelanggan: ")
	fmt.Scan(&jumlahPelanggan)
	daftarPelanggan := make([]Pelanggan, jumlahPelanggan)
	for i := 0; i < jumlahPelanggan; i++ {
		daftarPelanggan[i].ID = i + 1
		fmt.Printf("Nama pelanggan %d: ", i+1)
		fmt.Scan(&daftarPelanggan[i].Nama)
	}

	fmt.Print("\nMasukkan jumlah transaksi: ")
	fmt.Scan(&jumlahTransaksi)
	daftarTransaksi := make([]Transaksi, jumlahTransaksi)
	for i := 0; i < jumlahTransaksi; i++ {
		trans := &daftarTransaksi[i]
		trans.ID = i + 1
		fmt.Printf("\nTransaksi %d:\n", i+1)
		fmt.Print("  ID Pelanggan: ")
		fmt.Scan(&trans.IDPelanggan)
		var jumlahBeli int
		fmt.Print("  Jumlah produk yang dibeli: ")
		fmt.Scan(&jumlahBeli)
		trans.IDProduks = make([]int, jumlahBeli)
		for j := 0; j < jumlahBeli; j++ {
			fmt.Printf("    Masukkan ID produk ke-%d: ", j+1)
			fmt.Scan(&trans.IDProduks[j])
		}
		trans.Total = hitungTotal(daftarProduk, trans.IDProduks)
	}
	sortProduk(daftarProduk)
	sortTransaksi(daftarTransaksi)
