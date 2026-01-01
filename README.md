# Tubes Alfiko Fidanzato – Analisis Kompleksitas Algoritma

Repository ini berisi implementasi dan analisis perbandingan efisiensi
algoritma iteratif dan rekursif dalam menyelesaikan permasalahan
perhitungan pangkat bilangan (x^n) menggunakan bahasa pemrograman Go.

## Deskripsi Studi Kasus
Studi kasus yang diangkat adalah menghitung nilai x^n untuk bilangan
bulat positif n. Permasalahan ini dipilih karena dapat diselesaikan
menggunakan pendekatan iteratif maupun rekursif, sehingga cocok untuk
menganalisis perbedaan efisiensi waktu eksekusi dari kedua algoritma.

## Algoritma yang Digunakan
1. **Algoritma Iteratif**  
   Menghitung nilai x^n dengan melakukan perkalian berulang sebanyak n
   kali menggunakan perulangan. Pendekatan ini sederhana dan efisien
   untuk nilai n yang besar.  
   Kompleksitas waktu: **O(n)**

2. **Algoritma Rekursif**  
   Menghitung nilai x^n dengan memanggil fungsi secara berulang hingga
   mencapai kondisi dasar. Pendekatan ini lebih dekat dengan definisi
   matematis, namun memiliki overhead pemanggilan fungsi.  
   Kompleksitas waktu: **O(n)**

## Cara Menjalankan Program
Pastikan bahasa pemrograman Go telah terpasang di sistem, kemudian
jalankan perintah berikut melalui terminal:

```bash
go run tubes.go
