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
