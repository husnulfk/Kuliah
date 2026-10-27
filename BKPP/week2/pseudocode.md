# Pseudocode dan Dokumentasi BKPP Week 2

Dokumen ini berisi konversi tujuh program Go menjadi pseudocode dengan struktur:
`program` → `kamus` → `algoritma` → `endfor` → `endprogram`.

## Daftar Program

| No  | Program        | Fungsi Singkat                                                   |
| --- | -------------- | ---------------------------------------------------------------- |
| 1   | bidangPersegi  | Menghitung luas dan keliling persegi dari beberapa sisi          |
| 2   | cetakBarisan   | Mencetak bilangan dari n sampai m                                |
| 3   | cetakString    | Mencetak sebuah kata sebanyak n kali                             |
| 4   | faktorBilangan | Menampilkan setiap d dari 1..n beserta status apakah d faktor n  |
| 5   | faktorial      | Menghitung n!                                                    |
| 6   | ngoding        | Menghitung rata-rata jam ngoding                                 |
| 7   | sandi          | Menjumlahkan digit pertama dan digit terakhir dari tiap bilangan |
| 8   | lele           | menghitung panjang lele setiap bulan                             |

## Konvensi Penulisan

| Notasi                          | Arti                                  | Padanan di Go             |
| ------------------------------- | ------------------------------------- | ------------------------- |
| `input(x)`                      | membaca nilai ke variabel x           | `fmt.Scan(&x)`            |
| `output(x)`                     | menampilkan nilai x                   | `fmt.Println(x)`          |
| `x <- y`                        | memberi nilai y ke x                  | `x = y`                   |
| `div`                           | pembagian bilangan bulat              | `/` pada int              |
| `mod`                           | sisa bagi                             | `%`                       |
| `for i <- a to b do ... endfor` | perulangan dari a sampai b (inklusif) | `for i := a; i <= b; i++` |

---

## 1. bidangPersegi

**Deskripsi:** Program membaca jumlah persegi `n`, lalu untuk setiap persegi membaca panjang sisinya dan menampilkan luas serta kelilingnya.

**Input:**

- Baris pertama: `n` (jumlah persegi)
- Berikutnya: `n` buah bilangan bulat (panjang sisi)

**Output:** Untuk tiap sisi, satu baris berisi `luas keliling` dipisah spasi.

**Rumus:** luas = sisi × sisi, keliling = 4 × sisi

**Pseudocode:**

```
program bidangPersegi

kamus
    n        : integer    { banyaknya persegi }
    sisi     : integer    { panjang sisi persegi }
    luas     : integer    { luas persegi }
    keliling : integer    { keliling persegi }

algoritma
    input(n)
    for i <- 1 to n do
        input(sisi)
        luas <- sisi * sisi
        keliling <- sisi * 4
        output(luas, keliling)
    endfor

endprogram
```

**Contoh:**

| Input       | Output              |
| ----------- | ------------------- |
| `2` `3` `5` | `9 12` lalu `25 20` |

---

## 2. cetakBarisan

**Deskripsi:** Program membaca dua bilangan `n` dan `m`, lalu mencetak semua bilangan bulat dari `n` sampai `m`, satu bilangan per baris.

**Input:** `n` dan `m` (bilangan bulat)

**Output:** Bilangan `n, n+1, ..., m`, masing-masing satu baris. Jika `n > m`, tidak ada yang dicetak.

**Pseudocode:**

```
program cetakBarisan

kamus
    n : integer    { bilangan awal }
    m : integer    { bilangan akhir }
    i : integer    { penghitung perulangan }

algoritma
    input(n, m)
    for i <- n to m do
        output(i)
    endfor

endprogram
```

**Contoh:**

| Input | Output                                     |
| ----- | ------------------------------------------ |
| `3 6` | `3` `4` `5` `6` (masing-masing satu baris) |

---

## 3. cetakString

**Deskripsi:** Program membaca jumlah pengulangan `n` dan sebuah kata, lalu mencetak kata tersebut sebanyak `n` kali, satu per baris.

**Input:** `n` (bilangan bulat) lalu `kata` (satu kata tanpa spasi)

**Output:** `kata` dicetak `n` baris.

**Catatan:** `input` untuk string hanya membaca satu kata (berhenti di spasi).

**Pseudocode:**

```
program cetakString

kamus
    n    : integer    { jumlah pengulangan }
    kata : string     { kata yang dicetak }
    i    : integer    { penghitung perulangan }

algoritma
    input(n)
    input(kata)
    for i <- 1 to n do
        output(kata)
    endfor

endprogram
```

**Contoh:**

| Input      | Output                                          |
| ---------- | ----------------------------------------------- |
| `3` `halo` | `halo` `halo` `halo` (masing-masing satu baris) |

---

## 4. faktorBilangan

**Deskripsi:** Program membaca bilangan `n`, lalu untuk setiap `d` dari 1 sampai `n` menampilkan `d` dan status apakah `d` merupakan faktor dari `n` (`true` jika `n` habis dibagi `d`, `false` jika tidak).

**Input:** `n` (bilangan bulat positif)

**Output:** `n` baris berisi `d status`.

**Pseudocode:**

```
program faktorBilangan

kamus
    n : integer    { bilangan yang diperiksa }
    d : integer    { calon pembagi / faktor }
    s : boolean    { true jika d faktor dari n }

algoritma
    input(n)
    for d <- 1 to n do
        s <- (n mod d = 0)
        output(d, s)
    endfor

endprogram
```

**Contoh:**

| Input | Output                                                       |
| ----- | ------------------------------------------------------------ |
| `6`   | `1 true`, `2 true`, `3 true`, `4 false`, `5 false`, `6 true` |

---

## 5. faktorial

**Deskripsi:** Program membaca bilangan `n` lalu menghitung n! = n × (n−1) × ... × 1.

**Input:** `n` (bilangan bulat tidak negatif)

**Output:** Nilai n!. Untuk `n = 0` hasilnya `1`.

**Cara kerja:** `jumlah` diawali 1. Pada iterasi ke-`i` (mulai 0), `jumlah` dikali `(n - i)`, sehingga faktor yang dikalikan berurutan `n, n-1, ..., 1`.

**Pseudocode:**

```
program faktorial

kamus
    n      : integer    { bilangan yang dihitung faktorialnya }
    jumlah : integer    { hasil perkalian / faktorial }
    i      : integer    { penghitung perulangan }

algoritma
    input(n)
    jumlah <- 1
    for i <- 0 to n do
        jumlah <- jumlah * (n - i)
    endfor
    output(jumlah)

endprogram
```

**Contoh:**

| Input | Output | Penjelasan        |
| ----- | ------ | ----------------- |
| `5`   | `120`  | 5 × 4 × 3 × 2 × 1 |

---

## 6. ngoding

**Deskripsi:** Program membaca banyaknya data `n`, lalu membaca `n` nilai jam ngoding, menjumlahkannya, dan menampilkan rata-ratanya sebagai bilangan desimal.

**Input:** `n`, lalu `n` bilangan bulat (jam ngoding)

**Output:** Rata-rata jam (bilangan real).

**Catatan:** Konversi ke real dilakukan sebelum pembagian agar hasilnya tidak dibulatkan. Jika `n = 0`, terjadi pembagian dengan nol, jadi `n` harus lebih dari 0.

**Pseudocode:**

```
program ngoding

kamus
    n        : integer    { banyaknya data }
    jam      : integer    { jam ngoding satu data }
    totalJam : integer    { jumlah seluruh jam }
    rata     : real       { rata-rata jam }
    i        : integer    { penghitung perulangan }

algoritma
    input(n)
    totalJam <- 0
    for i <- 1 to n do
        input(jam)
        totalJam <- totalJam + jam
    endfor
    rata <- totalJam / n    { pembagian real }
    output(rata)

endprogram
```

**Contoh:**

| Input           | Output               |
| --------------- | -------------------- |
| `3` `2` `4` `5` | `3.6666666666666665` |

---

## 7. sandi

**Deskripsi:** Program membaca `n` buah bilangan empat digit. Untuk tiap bilangan, diambil digit pertama (ribuan) dan digit terakhir (satuan), lalu semuanya dijumlahkan.

**Input:** `n`, lalu `n` bilangan empat digit.

**Output:** Total jumlah digit pertama dan digit terakhir dari semua bilangan.

**Cara kerja:**

- digit pertama = `bilangan div 1000`
- digit terakhir = `bilangan mod 10`

**Pseudocode:**

```
program sandi

kamus
    n        : integer    { banyaknya bilangan }
    bilangan : integer    { bilangan empat digit }
    d1       : integer    { digit pertama (ribuan) }
    d4       : integer    { digit terakhir (satuan) }
    jumlah   : integer    { total seluruh digit }
    i        : integer    { penghitung perulangan }

algoritma
    input(n)
    jumlah <- 0
    for i <- 1 to n do
        input(bilangan)
        d1 <- bilangan div 1000
        d4 <- bilangan mod 10
        jumlah <- jumlah + d1 + d4
    endfor
    output(jumlah)

endprogram
```

**Contoh:**

| Input             | Output | Penjelasan    |
| ----------------- | ------ | ------------- |
| `2` `1234` `5678` | `20`   | (1+4) + (5+8) |

## 8. lele

**Deskripsi:** Program menghitung panjang ikan lele setelah 14 bulan. Panjang awal lele adalah 3,0 dan bertambah 0,3 setiap bulan. Program juga mencatat jumlah bulan yang sudah berlalu.

**Input:** tidak ada (nilai awal dan jumlah bulan ditulis langsung di program)

**Output:** Satu baris berisi `panjangLele jumlahBulan` dipisah spasi.

**Cara kerja:** Pada setiap bulan (`i` dari 1 sampai 14), `panjangLele` ditambah 0,3 dan `jumlahBulan` ditambah 1.

**Catatan:** Karena `panjangLele` bertipe real, penjumlahan 0,3 berulang kali bisa menghasilkan selisih kecil akibat pembulatan komputer. Nilai matematisnya 3,0 + 14 × 0,3 = 7,2, tetapi program Go menampilkan `7.1999999999999975`.

**Pseudocode:**

```
program lele

kamus
    panjangLele : real       { panjang lele saat ini }
    jumlahBulan : integer    { banyaknya bulan yang sudah berlalu }
    i           : integer    { penghitung perulangan }

algoritma
    panjangLele <- 3.0
    jumlahBulan <- 0
    for i <- 1 to 14 do
        panjangLele <- panjangLele + 0.3
        jumlahBulan <- jumlahBulan + 1
    endfor
    output(panjangLele, jumlahBulan)

endprogram
```

**Contoh:**

| Input       | Output                  | Penjelasan                            |
| ----------- | ----------------------- | ------------------------------------- |
| (tidak ada) | `7.1999999999999975 14` | 3,0 + 14 × 0,3 ≈ 7,2 setelah 14 bulan |
