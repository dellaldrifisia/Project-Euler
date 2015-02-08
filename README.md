# Programming in GO

##BAB 1. PENDAHULUAN

Go adalah bahasa pemrogaman yang dirancang oleh Google dan mudah untuk dipelajari oleh anda yang baru pertama kali belajar bahasa pemrogaman. Selain itu, Go adalah bahasa pemrogaman yang mudah untuk membangun perangkat lunak sederhana, handal dan efisien. Keunggulan dari Go ini adalah memiliki fitur yang canggih dan ketersediaan yang luas pada berbagai *platform*.

Tahapan yang digunakan dalam membuat aplikasi dengan Go cukup mudah, yaitu :

1.	Mengumpulkan persyaratan
2.	Mencari solusi
3.	Menuliskan *source code* yang mengimplementasikan solusi
4.	Mengompilasi *source code*
5.	Menjalankan program untuk memastikan apakah program tersebut bekerja

###File dan Folder

File adalah kumpulan data yang disimpan sebagai unit dengan nama tertentu. Semua jenis file disimpan dengan cara yang sama pada komputer, memiliki nama, ukuran dan jenis yang terkait. Biasanya jenis file ditandai dengan ekstensi file setelah nama pada bagian terakhir. Contoh : 

    contoh.txt
  
* contoh : nama file
* txt : ekstensi yang digunakan untuk mewakili data tekstual

Folder atau direktori digunakan untuk sekumpulan file secara bersama-sama. Dapat juga berisi folder lainnya. Contoh :

    D:\apps\golang\contoh.txt

* contoh.txt adalah nama file yang terkandung didalam folder ‘golang’, folder ‘golang tersebut terkandung didalam folder ‘apps’ yang disimpan di drive D

Pada dasarnya dalam pengaturan file dan folder yang terkait dengan pembuatan aplikasi menggunakan bahasa pemrograman Go berhubungan dengan proses preparasi yang terdiri dari beberapa tahapan, yaitu :

1.	Penginstallan Go
2.	Pengaturan properties pada komputer
3.	Penginstallan Git
4.	Penginstallan Mercurial (hg)
5.	Penginstallan Gorilla

###Terminal

Sebagian besar interaksi kita dengan komputer dilakukan melalui interaksi antarmuka sebagai pengguna grafis (GUI). Digunakan keyboard, mouse dan layar sentuh untuk berinteraksi dengan tombol visual atau jenis lainnya dari kontrol yang ditampilkan pada layar. Sebelum GUI berkembang, terlebih dahuku digunakan terminal yaitu sebuah tekstual antarmuka  sederhana untuk komputer yang digunakan dengan memberikan perintah dan menerima balasan untuk memanipulasi tombol pada layar komputer. 

Walaupun sebagian besar dunia komputasi telah meninggalkan terminal, kebenarannya adalah terminal masih fundamental yang digynakan oleh sebagian besar bahasa pemrogaman. Bahasa pemrogaman Go tidak berbeda, sehingga sebelum kita menulis program pada Go, terlebih dahulu kita perlu memiliki pemahaman dasar bagaimana terminal bekerja.  

###*Text Editor*

Sebagai programmer, alat utama yang digunakan untuk menulis program adalah Text Editor. Text Editor mirip dengan pengolahan kata seperti Microsoft Word. Text Editor yang biasa digunakan dalam bahasa Go adalah Sublime Text. 

###*Tools*

Go adalah bahasa pemrogaman yang dikompilasi, artinya source code diterjemahkan ke dalam bahasa yang dapat dimengerti oleh komputer. Terkait akan hal tersebut, sebelum program Go dibuat, kita membutuhkan Go compiler. Installer akan mengatur Go untuk anda secara otomatis. Informasi lebih lanjut dapat ditemukan di http://www.golang.org


##BAB 2. PROGRAM PERTAMA

Didalam berbagai bahasa pemrogaman, biasanya anda diminta untuk membuat program sederhana “Hello World”. Yaitu sebuah program yang output nya hanya menampilkan “Hello World” pada terminal anda.

Untuk menjadi awalan dalam mempelajari Go, mari kita menulis program menggunakan Go. Hal pertama yang dilakukan adalah membuat folder baru untuk menyimpan program yang akan ditulis. Untuk mempermudah, pada teriminal anda berilah perintah dengan memasukkan : 

    mkdir apps/src/golang

Artinya, anda memerintahkan untuk membuat direktori baru (mkdir) misal pada drive D, dengan folder ‘apps’, didalamnya terdapat pula folder ‘src’, didalam folder src dibuat juga folder ‘golang’.

Pada text editor anda, buatlah file baru didalam folder golang, kemudian simpan file tersebut dengan nama main.go, lalu buatlah program berikut :

```go
package main
    
import "fmt"
// this is a comment
    
func main() {
  fmt.Println("Hello World")
}
```

Kemudian jalankan program tersebut pada terminal anda, tapi arahkan terlebih dahulu terminal anda ke direktori file anda disimpan.
Untuk menjalankan program, berikan perintah berikut pada terminal anda :

    go run main.go

Output yang akan ditampilkan adalah :

    Hello World

Perintah go run artinya menjalankan program dengan memanggil file yang dipisahkan dengan spasi, terlebih dahulu di compile menjadi executable yang disimpan dalam direktori sementara program berjalan. Apabila program anda tidak berjalan, maka anda membuat kesalahan dan Go compiler akan memberi petunjuk dimana letak kesalahan anda. Compiler Go tidak memiliki toleransi untuk kesalahan sekecil apapun.

###Cara Membaca Program Go

Selanjutnya, kita akan melihat program ini secara lebih terperinci. Program Go dibaca dari atas ke bawah dan dari kiri ke kanan. Baris pertama menyatakan ini :

```go
package main
```

Package main dikenal sebagai “package declaration”. Setiap program Go harus dimulai dengan mendeklarasikan sebuah package. Dimana package adalah cara pengorganisasian oleh Go dan code yang digunakan kembali. Ada dua jenis program Go, yaitu :

1.	Executables yaitu jenis program yang dapat dijalankan langsung dari terminal
2.	Libraries yaitu koleksi code dimana package yang kita buat dapat bersama-sama digunakan pada program lain

Selanjutnya :

```go
import "fmt"
```

Kaca kunci import adalah bagaimana kita menyertakan code dari package lain yang digunakan pada program ini. Package fmt (singkatan untuk format) adalah menerapkan format untuk input dan output. Tanda kutip ganda pada fmt (“fmt”) dikenal sebagai string literal” yang merupakan jenis ekspresi.

Selanjutnya, kita akan melihat deklarasi fungsi berikut :

```go
func main() {
  fmt.Println("Hello World")
}
```

*Functions* adalah blok bangunan dari program Go, dimana memiliki input, output dan serangkaian langkah-langkah yang disebut pernyataan kemudian dijalankan dalam kerangka tertentu.


##BAB 5. *CONTROL STRUCTURES*

*Control structures* dibutuhkan dalam mengerjakan suatu proses yang berulang-ulang (looping).

###For

For digunakan untuk mengulang daftar pernyataan (blok) beberapa kali sehingga kita tidak perlu menuliskannya secara berulang, cukup tuliskan perintahnya sekali saja. Contoh program untuk menampilkan angka 1 sampai 10:

```go
func main() {
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }
}
```

* i merupakan variabel yang digunakan untuk mengontrol bilangan atau angka yang ingin di tampilkan.
* Alur iterasi program di atas adalah: Untuk i=1 di cek apakah nilai i masih kecil atau sama dengan 10. Jika i bernilai kecil atau sama dengan 10 (pernyataan *TRUE*) maka akan dijalankan perintah selanjutnya yaitu `fmt.Println(i)` yang artinya nilai i akan ditampilkan selama masih memenuhi syarat kecil atau sama dengan 10. Setelah bilangan pertama ditampilakan maka nilai i akan bertambah melalui perintah `i++`. Nilai i selanjutnya akan menjalankan perintah yang sama dengan yang pertama. *Looping* akan berhenti saat i bernilai lebih dari 10.

###If

If merupakan *statement* logika yang biasanya diikuti oleh *statement* `else`. Perintah `if` akan dijalankan apabila pernyataan bernilai *TRUE*, namun jika pernyataan bernilai *FALSE* maka yang akan dijalankan adalah perintah pada bagian `else`. `if` dan `else` biasanya digunakan bersamaan dengan `for`. Contohnya adalah untuk menampilkan angka 1 sampai 10 serta jenisnya apakan bilangan ganjil atau genap.

```go
func main() {
    for i := 1; i <= 10; i++ {
        if i % 2 == 0 {
            fmt.Println(i, "genap")
        } else {
            fmt.Println(i, "ganjil")
        }
    }
}
```
* Alur iterasi program ini adalah: saat i=1 maka akan di cek terlebih dahulu apakan nilai i lebih kecil atau sama dengan 10. Apabila i kecil atau sama dengan 10 maka aka dilakukan perintah selanjutnya yaitu `if i % 2 == 0`. Jika pernyataan ini bernilai *TRUE* maka akan dijalankan perintah `fmt.Println(i, "genap")` yang akan menampilkan nilai i tersebut dan disertai dengan *string* 'genap'. Namun apabila pernyataan bernilai *FALSE* (i tidak habis dibagi 2) maka akan dijalankan perintah `fmt.Println(i, "ganjil")`. Begitu seterusnya hingga program akan berhenti apabila nilai i lebih besar dari 10.

###Switch

Misalkan kita ingin membuat program yang dapat menampilkan nama dari suatu angka. Agar lebih memudahkan kita dapat menuliskan programnya dengan menggunakan format `switch case`. Contoh: menampilkan nama angka dari 1 sampai 5.

```go
switch i {
case 0: fmt.Println("Nol")
case 1: fmt.Println("Satu")
case 2: fmt.Println("Dua")
case 3: fmt.Println("Tiga")
case 4: fmt.Println("Empat")
case 5: fmt.Println("Lima")
default: fmt.Println("Angka tidak diketahui")
}
```

*Statement* switch dimulai dengan kata kunci `switch` dan diikuti oleh sebuah ekspresi (dalam hal ini `i`) dan beberapa `case`. Nilai `i` akan dibandingkan dengan `case` yang cocok. Jika terdapat `case` yang sesuai dengan nilai `i` maka perintah dalam `case` tersebut akan dieksekusi. Namun apabila tidak ada `case` yang sesuai dengan nilai `i` maka yang akan dieksekusi adalah `default`.

##BAB 6. ARRAY, *SLICES*, DAN *MAPS*

###*Maps*

*Map* adalah koleksi atau sekumpulan pasangan *key-value*. Juga dikenal sebagai array asosiatif atau kamus. *Maps* digunakan untuk melihat *value* berdasarkan *key* yang berhubungan. Contoh: *Maps* beberapa unsur kimia yang di-indeks-kan oleh simbol nya.

```go
package main
import "fmt"
func main() {
    elements := make(map[string]string)
    elements["H"] = "Hydrogen"
    elements["He"] = "Helium"
    elements["Li"] = "Lithium"
    elements["Be"] = "Beryllium"
    elements["B"] = "Boron"
    elements["C"] = "Carbon"
    elements["N"] = "Nitrogen"
    elements["O"] = "Oxygen"
    elements["F"] = "Fluorine"
    elements["Ne"] = "Neon"
    fmt.Println(elements["Li"])
}
```


##BAB 7. FUNGSI

Fungsi merupakan bagian independen yang memetakan nol atau lebih parameter input ke nol atau lebih parameter output. Fungsi juga dikenal sebagai prosedur atau subrutin.

###Bekerja dengan Lebih dari Satu Fungsi

Pada program yang sebelumnya kita bekerja dengan hanya satu fungsi. Pada bagian ini akan dijelaskan bagaimana membuat program dengan lebih dari satu fungsi.

```go
package main
import "fmt"
func average(xs []float64) float64 {
    total := 0.0
    for _, v := range xs {
        total += v
    }
    return total / float64(len(xs))
}
func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
}
```

Program di atas akan memberikan output berupa nilai rata-rata dari angka yang ada di dalam array `xs`. Fungsi main merupakan fungsi utama yang dijalankan pertama kali. Di dalam fungsi main diinisialisasi sebuah array `xs` yang berisi lima bilangan. Lalu di dalam fungsi main juga terdapat perintah `fmt.Println(average(xs))` yang memanggil fungsi average dan memasukkan nilai `xs` ke dalamnya. Setelah perintah di dalam fungsi average dieksekusi maka akan dijalankan perintah untuk menampilkan nilai rata-rata dari kelima bilangan tersebut.

###Mengembalikan Lebih dari Satu Nilai

Dalam bahasa Go juga memungkinkan kita untuk  mengembalikan beberapa nilai dari suatu fungsi. Contoh:

```go
package main
import "fmt"
func f() (int, int) {
    return 5, 6
}
func main() {
    x, y := f()
}
```

Fungsi f mengembalikan dua nilai `(int, int)` yaitu 5 dan 6 yang selanjutnya dipanggil dalam fungsi main.

###Fungsi *Variadic*

Fungsi jenis ini merupakan fungsi yang hanya dapat dijalankan pada bahasa Go. Dengan fungsi ini kita dapat membuat suatu variabel yang dapat berisi nol atau lebih nilai dengan menggunakan `...` sebelum tipe variabel.

```go
func add(args ...int) int {
    total := 0
    for _, v := range args {
        total += v
    }
    return total
}
func main() {
    fmt.Println(add(1,2,3))
}
```

Dengan program diatas kita dapat memasukkan nilai ke dalam fungsi dalam jumlah yang tidak perlu ditentukan terlebih dahulu.

###*Closure*

*Closure* merupakan suatu fungsi yang berisi fungsi lain di dalamnya. Contoh:

```go
package main
import "fmt"
func main() {
    add := func(x, y int) int {
        return x + y
    }
    fmt.Println(add(1,1))
}
```

Di dalam fungsi main terdapat variabel `add` dengan tipe `func(int,int)int` (sebuah fungsi yang berisi dua `int` dan mengembalikan satu `int`.

###Rekursi

Suatu fungsi dapat memanggil dirinya sendiri dalam fungsi rekursi. Contohnya adalah dalam program penghitung nilai faktorial berikut:

```go
func factorial(x uint) uint {
    if x == 0 {
        return 1
    }
    return x * factorial(x-1)
}
func main(){
    fmt.Println(factorial(2))
}
```

Misalnya kita ingin menghitung nilai faktorial dari angka 2 melalui perintah: `fmt.Println(factorial(2))`

Fingsi main memanggil fungsi factorial dan memberikan nilai x=2. Dalam fungsi factorial nilai x tidak sama dengan 0 maka perintah selanjutnya yang akan dieksekusi adalah `return x * factorial(x-1)` yang berarti fungsi factorial tersebut akan memanggil dirinya sendiri dengan nilai x=1. Saat x=1 fungsi factorial akan memanggil kembali dirinya dengan x=0. Saat x=0 maka pernyataan `if` bernilai *TRUE* sehingga akan mengembalikan nilai 1. Pada akhirnya fungsi factorial akan mengembalikan nilai `2*1*1=2`. Jadi dapat dibuktikan bahwa nilai faktorial dari 2 adalah sama dengan 2.


##BAB 9. *STRUCTS* DAN *INTERFACES*

























