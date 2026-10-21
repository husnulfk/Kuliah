# Pseudocode dan Dokumentasi 10 Program (Bagian 3)

Dokumen ini berisi konversi sepuluh program Go menjadi pseudocode dengan struktur:
`program` → `kamus` → `algoritma` → `endprogram`.

## Daftar Program

| No  | Program      | Fungsi Singkat                                                               |
| --- | ------------ | ---------------------------------------------------------------------------- |
| 1   | cacah_uang   | Menghitung jumlah lembar pecahan 10.000, 5.000, dan 1.000 dari sejumlah uang |
| 2   | fungsi       | Menghitung x² + 2y + 3                                                       |
| 3   | fungsixy     | Menghitung 1/(3x² + 10) + 10y + 7 dengan bilangan bulat                      |
| 4   | jumlah       | Menjumlahkan dua bilangan bulat                                              |
| 5   | kabisat      | Memeriksa apakah suatu tahun adalah tahun kabisat                            |
| 6   | koversi_suhu | Mengonversi suhu Celsius ke Reamur, Fahrenheit, dan Kelvin                   |
| 7   | main_1       | Menampilkan kalimat sapaan                                                   |
| 8   | modulo       | Menghitung sisa bagi dua bilangan bulat                                      |
| 9   | segitiga     | Memeriksa apakah tiga sisi dapat membentuk segitiga                          |
| 10  | tigaDigit    | Memisahkan bilangan tiga digit menjadi digit ratusan, puluhan, dan satuan    |

## Konvensi Penulisan

| Notasi      | Arti                           | Padanan di Go     |
| ----------- | ------------------------------ | ----------------- |
| `input(x)`  | membaca nilai ke variabel x    | `fmt.Scan(&x)`    |
| `output(x)` | menampilkan nilai x            | `fmt.Println(x)`  |
| `x <- y`    | memberi nilai y ke x           | `x = y`           |
| `div`       | pembagian bilangan bulat       | `/` pada tipe int |
| `mod`       | sisa bagi                      | `%`               |
| `=`, `<>`   | sama dengan, tidak sama dengan | `==`, `!=`        |
| `and`, `or` | dan, atau (logika)             | `&&`, `\|\|`      |

---

## 1. cacah_uang

**Deskripsi:** Program membaca sejumlah uang, lalu menghitung berapa lembar pecahan 10.000, 5.000, dan 1.000 yang dibutuhkan, dengan mendahulukan pecahan terbesar.

**Input:** `uang` (bilangan bulat, kelipatan 1.000)

**Output:** Tiga baris: jumlah lembar 10.000, 5.000, dan 1.000.

**Catatan:** Variabel `uang` ditimpa dengan sisanya di setiap langkah, sama seperti kode Go aslinya.

**Pseudocode:**

```
program cacah_uang

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

## 2. fungsi

**Deskripsi:** Program membaca dua bilangan bulat `x` dan `y`, lalu menghitung hasil = x² + 2y + 3.

**Input:** `x` dan `y` (bilangan bulat)

**Output:** `hasil` (bilangan bulat)

**Pseudocode:**

```
program fungsi

kamus
    x     : integer    { bilangan pertama }
    y     : integer    { bilangan kedua }
    hasil : integer    { hasil perhitungan }

algoritma
    input(x, y)
    hasil <- (x * x) + (2 * y) + 3
    output(hasil)

endprogram
```

**Contoh:**

| Input | Output | Penjelasan |
| ----- | ------ | ---------- |
| `2 3` | `13`   | 4 + 6 + 3  |

---

## 3. fungsixy

**Deskripsi:** Program membaca dua bilangan bulat `x` dan `y`, lalu menghitung hasil = 1 / (3x² + 10) + 10y + 7.

**Input:** `x` dan `y` (bilangan bulat)

**Output:** `hasil` (bilangan bulat)

**Catatan penting:** Semua variabel bertipe integer, sehingga `1 / (3x² + 10)` adalah pembagian bilangan bulat (`div`). Karena `3x² + 10` selalu minimal 10, hasil bagi itu selalu `0`, sehingga hasil akhir sama dengan `10y + 7`. Jika ingin hasil pecahan yang sebenarnya, variabel harus bertipe real.

**Pseudocode:**

```
program fungsixy

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

## 4. jumlah

**Deskripsi:** Program membaca dua bilangan bulat `a` dan `b`, lalu menampilkan hasil penjumlahannya.

**Input:** `a` dan `b` (bilangan bulat)

**Output:** `a + b`

**Pseudocode:**

```
program jumlah

kamus
    a     : integer    { bilangan pertama }
    b     : integer    { bilangan kedua }
    hasil : integer    { hasil penjumlahan }

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

## 5. kabisat

**Deskripsi:** Program membaca sebuah tahun `x` lalu menampilkan `true` jika tahun tersebut kabisat dan `false` jika bukan.

**Aturan tahun kabisat:** Tahun kabisat jika habis dibagi 400, atau habis dibagi 4 tetapi tidak habis dibagi 100.

**Input:** `x` (bilangan bulat, tahun)

**Output:** `true` atau `false`

**Catatan:** Operator `and` lebih didahulukan daripada `or`, jadi ekspresinya dibaca: `(x mod 400 = 0) or ((x mod 4 = 0) and (x mod 100 <> 0))`. Tanda kurung ekstra ditambahkan di pseudocode agar maksudnya jelas.

**Pseudocode:**

```
program kabisat

kamus
    x     : integer    { tahun yang diperiksa }
    hasil : boolean    { true jika x tahun kabisat }

algoritma
    input(x)
    hasil <- (x mod 400 = 0) or ((x mod 4 = 0) and (x mod 100 <> 0))
    output(hasil)

endprogram
```

**Contoh:**

| Input  | Output  | Penjelasan                                     |
| ------ | ------- | ---------------------------------------------- |
| `2024` | `true`  | habis dibagi 4, tidak habis dibagi 100         |
| `1900` | `false` | habis dibagi 100 tetapi tidak habis dibagi 400 |
| `2000` | `true`  | habis dibagi 400                               |

---

## 6. koversi_suhu

**Deskripsi:** Program membaca suhu dalam Celsius, lalu mengonversinya ke Reamur, Fahrenheit, dan Kelvin.

**Rumus:**

- Reamur = Celsius × 4 / 5
- Fahrenheit = Celsius × 9 / 5 + 32
- Kelvin = Celsius + 273,15

**Input:** `celcius` (bilangan real)

**Output:** Satu baris berisi `reamur fahreinheit kelvin` dipisah spasi.

**Catatan:** Nama variabel mengikuti penulisan di kode asli (`celcius`, `fahreinheit`).

**Pseudocode:**

```
program koversi_suhu

kamus
    celcius     : real    { suhu dalam Celsius }
    reamur      : real    { suhu dalam Reamur }
    fahreinheit : real    { suhu dalam Fahrenheit }
    kelvin      : real    { suhu dalam Kelvin }

algoritma
    input(celcius)
    reamur <- celcius * 4 / 5
    fahreinheit <- celcius * 9 / 5 + 32
    kelvin <- celcius + 273.15
    output(reamur, fahreinheit, kelvin)

endprogram
```

**Contoh:**

| Input | Output          |
| ----- | --------------- |
| `100` | `80 212 373.15` |

---

## 7. main_1

**Deskripsi:** Program tanpa input yang hanya menampilkan satu kalimat sapaan.

**Input:** tidak ada

**Output:** `hello good day everyone`

**Catatan:** Karena tidak ada data yang disimpan, bagian `kamus` kosong.

**Pseudocode:**

```
program main_1

kamus
    { tidak ada variabel }

algoritma
    output("hello good day everyone")

endprogram
```

**Contoh:**

| Input       | Output                    |
| ----------- | ------------------------- |
| (tidak ada) | `hello good day everyone` |

---

## 8. modulo

**Deskripsi:** Program membaca dua bilangan bulat `a` dan `b`, lalu menampilkan sisa bagi `a` oleh `b`.

**Input:** `a` dan `b` (bilangan bulat, `b` tidak boleh 0)

**Output:** `a mod b`

**Catatan:** Jika `b = 0`, program Go akan berhenti dengan error (pembagian dengan nol).

**Pseudocode:**

```
program modulo

kamus
    a     : integer    { bilangan yang dibagi }
    b     : integer    { pembagi, tidak boleh 0 }
    hasil : integer    { sisa bagi a oleh b }

algoritma
    input(a, b)
    hasil <- a mod b
    output(hasil)

endprogram
```

**Contoh:**

| Input  | Output | Penjelasan     |
| ------ | ------ | -------------- |
| `17 5` | `2`    | 17 = 3 × 5 + 2 |

---

## 9. segitiga

**Deskripsi:** Program membaca tiga panjang sisi `a`, `b`, `c`, lalu menampilkan `true` jika ketiganya dapat membentuk segitiga dan `false` jika tidak.

**Aturan:** Ketiga sisi dapat membentuk segitiga jika jumlah setiap dua sisi lebih besar daripada sisi yang ketiga (ketaksamaan segitiga).

**Input:** `a`, `b`, `c` (bilangan bulat)

**Output:** `true` atau `false`

**Pseudocode:**

```
program segitiga

kamus
    a     : integer    { panjang sisi pertama }
    b     : integer    { panjang sisi kedua }
    c     : integer    { panjang sisi ketiga }
    hasil : boolean    { true jika a, b, c membentuk segitiga }

algoritma
    input(a, b, c)
    hasil <- (a + b > c) and (b + c > a) and (a + c > b)
    output(hasil)

endprogram
```

**Contoh:**

| Input   | Output  | Penjelasan             |
| ------- | ------- | ---------------------- |
| `3 4 5` | `true`  | 3+4>5, 4+5>3, 3+5>4    |
| `1 2 3` | `false` | 1+2 tidak lebih dari 3 |

---

## 10. tigaDigit

**Deskripsi:** Program membaca bilangan tiga digit `x`, lalu memisahkannya menjadi digit ratusan (`d1`), puluhan (`d2`), dan satuan (`d3`), kemudian menampilkannya dengan label.

**Cara kerja:**

- `d1 = x div 100`
- `d2 = (x mod 100) div 10`
- `d3 = x mod 10`

**Input:** `x` (bilangan bulat tiga digit)

**Output:** Satu baris dengan format `d1 : ... d2 : ... d3 : ...`. Spasi pada teks mengikuti kode asli (dua spasi sebelum `d2`, satu spasi sebelum `d3`, dan satu spasi di akhir).

**Pseudocode:**

```
program tigaDigit

kamus
    x  : integer    { bilangan tiga digit }
    d1 : integer    { digit ratusan }
    d2 : integer    { digit puluhan }
    d3 : integer    { digit satuan }

algoritma
    input(x)
    d1 <- x div 100
    d2 <- (x mod 100) div 10
    d3 <- x mod 10
    output("d1 : ", d1, "  d2 : ", d2, " d3 : ", d3, " ")

endprogram
```

**Contoh:**

| Input | Output                   |
| ----- | ------------------------ |
| `456` | `d1 : 4  d2 : 5 d3 : 6 ` |
