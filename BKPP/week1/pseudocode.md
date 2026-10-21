# Pseudocode dan Dokumentasi 6 Program BKPP Week 1

Dokumen ini berisi konversi enam program Go menjadi pseudocode dengan struktur:
`program` → `kamus` → `algoritma` → `endprogram`.

## Daftar Program

| No  | Program       | Fungsi Singkat                                                               |
| --- | ------------- | ---------------------------------------------------------------------------- |
| 1   | cacahUang     | Menghitung jumlah lembar pecahan 10.000, 5.000, dan 1.000 dari sejumlah uang |
| 2   | main          | Menjumlahkan dua bilangan bulat                                              |
| 3   | matematikaA   | Menghitung ekspresi 1/(3x² + 10) + 10y + 7 dengan bilangan bulat             |
| 4   | matematikaB   | Menghitung (x³ + 3x) / (x⁴ − 3x² + 4) dengan bilangan real                   |
| 5   | mencariPosisi | Menghitung posisi akhir dari posisi awal, kecepatan, dan waktu               |
| 6   | suhu          | Mengonversi suhu Celsius ke Reamur, Fahrenheit, dan Kelvin                   |

## Konvensi Penulisan

| Notasi      | Arti                               | Padanan di Go         |
| ----------- | ---------------------------------- | --------------------- |
| `input(x)`  | membaca nilai ke variabel x        | `fmt.Scan(&x)`        |
| `output(x)` | menampilkan nilai x                | `fmt.Println(x)`      |
| `x <- y`    | memberi nilai y ke x               | `x = y`               |
| `div`       | pembagian bilangan bulat           | `/` pada tipe int     |
| `mod`       | sisa bagi                          | `%`                   |
| `/`         | pembagian real (jika operand real) | `/` pada tipe float64 |

---

## 1. cacahUang

**Deskripsi:** Program membaca sejumlah uang, lalu menghitung berapa lembar pecahan 10.000, 5.000, dan 1.000 yang dibutuhkan, dengan mendahulukan pecahan terbesar (cara greedy).

**Input:** `uang` (bilangan bulat, kelipatan 1.000)

**Output:** Tiga baris: jumlah lembar 10.000, jumlah lembar 5.000, jumlah lembar 1.000.

**Cara kerja:**

1. Lembar 10.000 = `uang div 10000`. Sisa uang = `uang mod 10000`.
2. Lembar 5.000 = `sisa div 5000`. Sisa uang = `sisa mod 5000`.
3. Lembar 1.000 = `sisa div 1000`.

**Catatan:** Variabel `uang` ditimpa dengan sisanya di setiap langkah (sama seperti kode Go aslinya).

**Pseudocode:**

```
program cacahUang

kamus
    uang : integer    { jumlah uang; setelah tiap langkah berisi sisa uang }

algoritma
    input(uang)
    output(uang div 10000)    { lembar 10.000 }
    uang <- uang mod 10000
    output(uang div 5000)     { lembar 5.000 }
    uang <- uang mod 5000
    output(uang div 1000)     { lembar 1.000 }

endprogram
```

**Contoh:**

| Input   | Output      | Penjelasan                         |
| ------- | ----------- | ---------------------------------- |
| `17000` | `1` `1` `2` | 1 × 10.000 + 1 × 5.000 + 2 × 1.000 |

---

## 2. main

**Deskripsi:** Program membaca dua bilangan bulat `a` dan `b`, lalu menampilkan hasil penjumlahannya.

**Input:** `a` dan `b` (bilangan bulat)

**Output:** `a + b`

**Pseudocode:**

```
program main

kamus
    a     : integer    { bilangan pertama }
    b     : integer    { bilangan kedua }
    hasil : integer    { hasil penjumlahan a dan b }

algoritma
    input(a, b)
    hasil <- a + b
    output(hasil)

endprogram
```

**Contoh:**

| Input | Output |
| ----- | ------ |
| `3 4` | `7`    |

---

## 3. matematikaA

**Deskripsi:** Program membaca dua bilangan bulat `x` dan `y`, lalu menghitung:

hasil = 1 / (3x² + 10) + 10y + 7

**Input:** `x` dan `y` (bilangan bulat)

**Output:** `hasil` (bilangan bulat)

**Catatan penting:** Semua variabel bertipe integer, sehingga `1 / (3x² + 10)` adalah pembagian bilangan bulat (`div`). Karena `3x² + 10` selalu minimal 10 (lebih besar dari 1), hasil bagi itu selalu `0`. Jadi hasil akhirnya sama dengan `10y + 7`. Jika ingin hasil pecahan yang sebenarnya, variabel harus bertipe real.

**Pseudocode:**

```
program matematikaA

kamus
    x     : integer    { bilangan pertama }
    y     : integer    { bilangan kedua }
    hasil : integer    { hasil perhitungan }

algoritma
    input(x, y)
    hasil <- (1 div (3 * (x * x) + 10)) + (10 * y) + 7
    output(hasil)

endprogram
```

**Contoh:**

| Input | Output | Penjelasan                    |
| ----- | ------ | ----------------------------- |
| `2 3` | `37`   | 1 div 22 = 0, lalu 0 + 30 + 7 |

---

## 4. matematikaB

**Deskripsi:** Program membaca bilangan real `x`, lalu menghitung:

hasil = (x³ + 3x) / (x⁴ − 3x² + 4)

**Input:** `x` (bilangan real)

**Output:** `hasil` (bilangan real)

**Catatan:** Penyebut `x⁴ − 3x² + 4` tidak pernah bernilai 0 untuk semua `x` real (nilai minimumnya 0,75), jadi tidak ada risiko pembagian dengan nol.

**Pseudocode:**

```
program matematikaB

kamus
    x     : real    { masukan }
    hasil : real    { hasil perhitungan }

algoritma
    input(x)
    hasil <- ((x * x * x) + 3 * x) / ((x * x * x * x) - 3 * (x * x) + 4)
    output(hasil)

endprogram
```

**Contoh:**

| Input | Output | Penjelasan                    |
| ----- | ------ | ----------------------------- |
| `1`   | `2`    | (1 + 3) / (1 − 3 + 4) = 4 / 2 |

---

## 5. mencariPosisi

**Deskripsi:** Program membaca posisi awal, kecepatan, dan waktu, lalu menghitung posisi akhir.

**Rumus:** posisi = posisi0 + (s × t)

**Input:** `posisi0` (posisi awal), `s` (kecepatan), `t` (waktu), semuanya bilangan bulat

**Output:** `posisi` (posisi akhir)

**Pseudocode:**

```
program mencariPosisi

kamus
    posisi0 : integer    { posisi awal }
    s       : integer    { kecepatan }
    t       : integer    { waktu }
    posisi  : integer    { posisi akhir }

algoritma
    input(posisi0, s, t)
    posisi <- posisi0 + (s * t)
    output(posisi)

endprogram
```

**Contoh:**

| Input    | Output | Penjelasan   |
| -------- | ------ | ------------ |
| `10 5 2` | `20`   | 10 + (5 × 2) |

---

## 6. suhu

**Deskripsi:** Program membaca suhu dalam Celsius, lalu mengonversinya ke Reamur, Fahrenheit, dan Kelvin.

**Rumus:**

- Reamur = Celsius × 4 / 5
- Fahrenheit = Celsius × 9 / 5 + 32
- Kelvin = Celsius + 273,15

**Input:** `celsius` (bilangan real)

**Output:** Satu baris berisi `reamur fahrenheit kelvin` dipisah spasi.

**Pseudocode:**

```
program suhu

kamus
    celsius    : real    { suhu dalam Celsius }
    reamur     : real    { suhu dalam Reamur }
    fahrenheit : real    { suhu dalam Fahrenheit }
    kelvin     : real    { suhu dalam Kelvin }

algoritma
    input(celsius)
    reamur <- celsius * 4 / 5
    fahrenheit <- celsius * 9 / 5 + 32
    kelvin <- celsius + 273.15
    output(reamur, fahrenheit, kelvin)

endprogram
```

**Contoh:**

| Input | Output          |
| ----- | --------------- |
| `100` | `80 212 373.15` |
