package main

import "fmt"

func main() {
 fmt.Println("Selamat datang di kalkulator sederhana")
 fmt.Println("— — — — — — — — — — — — — — — — — — —" )
 var a int
 fmt.Println("1. Penjumlahan")
 fmt.Println("2. Pengurangan")
 fmt.Println("3. Perkalian")
 fmt.Println("4. Pembagian")
 fmt.Println("5. Akar")
 fmt.Println("6. Pangkat")
 fmt.Println("7. Luas Persegi")
 fmt.Println("8. Luas Lingkaran")
 fmt.Println("9. Volume Balok")
 fmt.Println("10. Exit")
 fmt.Print("Masukkan menu yang anda inginkan: ")
 fmt.Scanln(&a)
 switch a {
  case 1:
     fmt.Println("— — — — — — — — — — — — — — — — — — —")
     fmt.Println("Menu Penjumlahan")
     var b int
     var c int
     fmt.Print("Masukkan angka pertama: ")
     fmt.Scanln(&b)
     fmt.Print("Masukkan angka kedua: ")
     fmt.Scanln(&c)
     var jumlah int = b+c
     fmt.Print("Hasilnya adalah ")
     fmt.Println(jumlah)
     fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
  case 2:
     fmt.Println("— — — — — — — — — — — — — — — — — — —")
     fmt.Println("Menu Pengurangan")
     var b int
     var c int
     fmt.Print("Masukkan angka pertama: ")
     fmt.Scanln(&b)
     fmt.Print("Masukkan angka kedua: ")
     fmt.Scanln(&c)
     var jumlah int = b-c
     fmt.Print("Hasilnya adalah ")
     fmt.Println(jumlah)
     fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 3:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Menu Perkalian")
      var b int
      var c int
      fmt.Print("Masukkan angka pertama: ")
      fmt.Scanln(&b)
      fmt.Print("Masukkan angka kedua: ")
      fmt.Scanln(&c)
      var jumlah int = b*c
      fmt.Print("Hasilnya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 4:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Menu Pembagian")
      var b int
      var c int
      fmt.Print("Masukkan angka pertama: ")
      fmt.Scanln(&b)
      fmt.Print("Masukkan angka kedua: ")
      fmt.Scanln(&c)
      var jumlah int = b/c
      fmt.Print("Hasilnya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 5:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Menu Pengurangan")
      var b int
      var c int
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&b)
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&c)
      var jumlah int = b*b
      fmt.Print("Hasilnya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 6:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Menu Pengurangan")
      var b int
      var c int
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&b)
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&c)
      var jumlah int = b*b*b
      fmt.Print("Hasilnya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 7:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Las Persegi")
      var b int
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&b)
      var jumlah int = b*b
      fmt.Print("Luas Perseginya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 8:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Las Lingkaran")
      var b float32
      var phi float32 
      phi = 3.14
      fmt.Print("Masukkan angka: ")
      fmt.Scanln(&b)
      var jumlah = phi*b*b
      fmt.Print("Luas Lingkarannya adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
   case 9:
      fmt.Println("— — — — — — — — — — — — — — — — — — —")
      fmt.Println("Las Lingkaran")
      var b float32
      var c float32 
      var d float32
      fmt.Print("Masukkan panjang: ")
      fmt.Scanln(&b)
      fmt.Print("Masukkan lebar: ")
      fmt.Scanln(&c)
      fmt.Print("Masukkan tinggi: ")
      fmt.Scanln(&d)
      var jumlah = b*c*d
      fmt.Print("Volume balok adalah ")
      fmt.Println(jumlah)
      fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
  case 12:
     fmt.Println(" — — — — — — — — — — — — — — — — — — — ")
     fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
  default:
     fmt.Println("Menu yang anda masukkan salah. Aplikasi akan keluar")
     fmt.Println("Terima kasih telah menggunakan aplikasi kalkulator")
  }
 
}