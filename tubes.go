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
