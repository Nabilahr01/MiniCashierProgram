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
	for {
		fmt.Print("\nMasukkan jumlah produk: ")
		fmt.Scan(&jumlahProduk)
		if jumlahProduk < 0 {
			fmt.Println("Jumlah produk tidak boleh negatif. Silakan coba lagi.")
		} else {
			break
		}
	daftarProduk := make([]Produk, jumlahProduk)
	for i := 0; i < jumlahProduk; i++ {
		fmt.Printf("Produk %d:\n", i+1)
		daftarProduk[i].ID = i + 1
		fmt.Print("  Nama produk: ")
		fmt.Scan(&daftarProduk[i].Nama)
		fmt.Print("  Harga produk: ")
		fmt.Scan(&daftarProduk[i].Harga)
	}

	for {
		fmt.Print("\nMasukkan jumlah pelanggan: ")
		fmt.Scan(&jumlahPelanggan)
		if jumlahPelanggan < 0 {
			fmt.Println("Jumlah pelanggan tidak boleh negatif. Silakan coba lagi.")
		} else {
			break
		}
	}
	daftarPelanggan := make([]Pelanggan, jumlahPelanggan)
	for i := 0; i < jumlahPelanggan; i++ {
		daftarPelanggan[i].ID = i + 1
		fmt.Printf("Nama pelanggan %d: ", i+1)
		fmt.Scan(&daftarPelanggan[i].Nama)
	}

	for {
		fmt.Print("\nMasukkan jumlah transaksi: ")
		fmt.Scan(&jumlahTransaksi)
		if jumlahTransaksi < 0 {
			fmt.Println("Jumlah transaksi tidak boleh negatif. Silakan coba lagi.")
		} else {
			break
		}
	}
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
	var pilihan int
	for {
		fmt.Println("\n===<< MENU >>===")
		fmt.Println("1. Tampilkan daftar produk (terurut harga)")
		fmt.Println("2. Tampilkan daftar transaksi (terurut total)")
		fmt.Println("3. Cari produk berdasarkan nama")
		fmt.Println("4. Cari transaksi berdasarkan ID")
		fmt.Println("5. Keluar")
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilihan)
		switch pilihan {
		case 1:
			fmt.Println("\nDaftar Produk (terurut berdasarkan harga):")
			for i := 0; i < len(daftarProduk); i++ {
				p := daftarProduk[i]
				fmt.Printf("- ID %d: %s - Rp%.0f\n", p.ID, p.Nama, p.Harga)
			}
		case 2:
			fmt.Println("\nDaftar Transaksi (terurut berdasarkan total):")
			for i := 0; i < len(daftarTransaksi); i++ {
				t := daftarTransaksi[i]
				namaPelanggan := dataPelangganByID(daftarPelanggan, t.IDPelanggan)
				// Ambil daftar nama produk
				var daftarNamaProduk []string
				for j := 0; j < len(t.IDProduks); j++ {
					namaProduk := dataProdukByID(daftarProduk, t.IDProduks[j])
					daftarNamaProduk = append(daftarNamaProduk, namaProduk)
				}
				namaProdukGabung := strings.Join(daftarNamaProduk, ", ")
				fmt.Printf("- ID Transaksi %d | ID Pelanggan %d (%s)\n", t.ID, t.IDPelanggan, namaPelanggan)
				fmt.Printf("  Produk dibeli: %s\n", namaProdukGabung)
				fmt.Printf("  Total: Rp%.0f\n\n", t.Total)
			}
		case 3:
			var keyword string
			fmt.Print("Masukkan kata kunci nama produk: ")
			fmt.Scan(&keyword)
			hasil := cariProdukNama(daftarProduk, keyword)
			if len(hasil) == 0 {
				fmt.Println("Produk tidak ditemukan.")
			} else {
				fmt.Println("Hasil pencarian:")
				for i := 0; i < len(hasil); i++ {
					p := hasil[i]
					fmt.Printf("- ID %d: %s - Rp%.0f\n", p.ID, p.Nama, p.Harga)
				}
			}
		case 4:
			var cariID int
			fmt.Print("Masukkan ID transaksi: ")
			fmt.Scan(&cariID)
			trans := cariTransaksiID(daftarTransaksi, cariID)
			if trans != nil {
				namaPelanggan := dataPelangganByID(daftarPelanggan, trans.IDPelanggan)

				var daftarNamaProduk []string
				for i := 0; i < len(trans.IDProduks); i++ {
					namaProduk := dataProdukByID(daftarProduk, trans.IDProduks[i])
					daftarNamaProduk = append(daftarNamaProduk, namaProduk)
				}
				namaProdukGabung := strings.Join(daftarNamaProduk, ", ")

				fmt.Printf("\nDitemukan:\n")
				fmt.Printf("ID Transaksi %d | ID Pelanggan %d (%s)\n", trans.ID, trans.IDPelanggan, namaPelanggan)
				fmt.Printf("Produk dibeli: %s\n", namaProdukGabung)
				fmt.Printf("Total: Rp%.0f\n", trans.Total)
			} else {
				fmt.Println("Transaksi tidak ditemukan.")
			}
		case 5:
			fmt.Println("Terima kasih telah belanja di kitamart")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}
	}
}
}

