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


##BAB2. PROGRAM PERTAMA

Didalam berbagai bahasa pemrogaman, biasanya anda diminta untuk membuat program sederhana “Hello World”. Yaitu sebuah program yang output nya hanya menampilkan “Hello World” pada terminal anda.

Untuk menjadi awalan dalam mempelajari Go, mari kita menulis program menggunakan Go. Hal pertama yang dilakukan adalah membuat folder baru untuk menyimpan program yang akan ditulis. Untuk mempermudah, pada teriminal anda berilah perintah dengan memasukkan : 

    mkdir apps/src/golang

Artinya, anda memerintahkan untuk membuat direktori baru (mkdir) misal pada drive D, dengan folder ‘apps’, didalamnya terdapat pula folder ‘src’, didalam folder src dibuat juga folder ‘golang’.

Pada text editor anda, buatlah file baru didalam folder golang, kemudian simpan file tersebut dengan nama main.go, lalu buatlah program berikut :

    package main
    
    import "fmt"
    // this is a comment
    
    func main() {
      fmt.Println("Hello World")
    }

Kemudian jalankan program tersebut pada terminal anda, tapi arahkan terlebih dahulu terminal anda ke direktori file anda disimpan.
Untuk menjalankan program, berikan perintah berikut pada terminal anda :

    go run main.go

Output yang akan ditampilkan adalah :

    Hello World

Perintah go run artinya menjalankan program dengan memanggil file yang dipisahkan dengan spasi, terlebih dahulu di compile menjadi executable yang disimpan dalam direktori sementara program berjalan. Apabila program anda tidak berjalan, maka anda membuat kesalahan dan Go compiler akan memberi petunjuk dimana letak kesalahan anda. Compiler Go tidak memiliki toleransi untuk kesalahan sekecil apapun.

###Cara Membaca Program Go

Selanjutnya, kita akan melihat program ini secara lebih terperinci. Program Go dibaca dari atas ke bawah dan dari kiri ke kanan. Baris pertama menyatakan ini :

    package main

Package main dikenal sebagai “package declaration”. Setiap program Go harus dimulai dengan mendeklarasikan sebuah package. Dimana package adalah cara pengorganisasian oleh Go dan code yang digunakan kembali. Ada dua jenis program Go, yaitu :

1.	Executables yaitu jenis program yang dapat dijalankan langsung dari terminal
2.	Libraries yaitu koleksi code dimana package yang kita buat dapat bersama-sama digunakan pada program lain

Selanjutnya :

    import "fmt"

Kaca kunci import adalah bagaimana kita menyertakan code dari package lain yang digunakan pada program ini. Package fmt (singkatan untuk format) adalah menerapkan format untuk input dan output. Tanda kutip ganda pada fmt (“fmt”) dikenal sebagai string literal” yang merupakan jenis ekspresi.

Selanjutnya, kita akan melihat deklarasi fungsi berikut :

    func main() {
      fmt.Println("Hello World")
    }

*Functions* adalah blok bangunan dari program Go, dimana memiliki input, output dan serangkaian langkah-langkah yang disebut pernyataan kemudian dijalankan dalam kerangka tertentu.

Masih belum selesai uni maaf, aku lg ngerjain di rumah diego. Tiba2 mbahnya sakit gt uni gak sadarin diri, 











