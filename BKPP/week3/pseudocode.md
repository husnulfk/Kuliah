# Pseudocode dan Dokumentasi BKPP Week 3

Dokumen ini berisi konversi tiga program Go menjadi pseudocode dengan struktur:
`program` → `kamus` → `algoritma` → `endfor` → `endprogram`.

## Daftar Program

| No  | Program     | Fungsi Singkat                                                         |
| --- | ----------- | ---------------------------------------------------------------------- |
| 1   | hadiah      | Menentukan apakah adik berhak hadiah (total gol minimal 10)            |
| 2   | jumlahDeret | Menjumlahkan deret 4/n + 4/(n+1) + ... + 4/m dan mencetak tiap langkah |
| 3   | tabungan    | Menghitung total tabungan kakak setelah n minggu (naik Rp 2500/minggu) |

## Konvensi Penulisan

| Notasi                                          | Arti                                                                                                       | Padanan di Go                             |
| ----------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ----------------------------------------- |
| `input(x)`                                      | membaca nilai ke variabel x                                                                                | `fmt.Scan(&x)`                            |
| `output(x)`                                     | menampilkan nilai x                                                                                        | `fmt.Println(x)`                          |
| `x <- y`                                        | memberi nilai y ke x                                                                                       | `x = y`                                   |
| `real(x)`                                       | mengubah bilangan bulat x menjadi real                                                                     | `float64(x)`                              |
| `for i <- a to b do ... endfor`                 | perulangan dari a sampai b (inklusif)                                                                      | `for i := a; i <= b; i++`                 |
| `for input(x); kondisi; input(x) do ... endfor` | perulangan dengan sentinel: baca x di awal, ulangi selama kondisi benar, baca x lagi di akhir tiap putaran | `for fmt.Scan(&x); kondisi; fmt.Scan(&x)` |

---

## 1. hadiah

**Deskripsi:** Program membaca sederetan bilangan yang menyatakan jumlah gol adik, diakhiri bilangan sentinel `-1`. Seluruh bilangan sebelum sentinel dijumlahkan. Adik mendapat hadiah jika total gol sekurang-kurangnya 10.

**Input:** Sebaris bilangan bulat yang diakhiri `-1` (sentinel, tidak ikut dijumlahkan).

**Output:** `true` jika adik diberi hadiah, `false` jika tidak.

**Rumus:** hadiah = (total >= 10)

**Cara kerja:** Bilangan pertama dibaca sebelum perulangan. Selama bilangan belum `-1`, bilangan ditambahkan ke `total`, lalu bilangan berikutnya dibaca. Setelah perulangan berhenti, kondisi `total >= 10` langsung menghasilkan nilai boolean sehingga tidak diperlukan percabangan.

**Pseudocode:**

```
program hadiah

kamus
    n      : integer    { gol yang dibaca }
    total  : integer    { jumlah seluruh gol }
    hadiah : boolean    { true jika adik berhak hadiah }

algoritma
    total <- 0
    for input(n); n != -1; input(n) do
        total <- total + n
    endfor
    hadiah <- (total >= 10)
    output(hadiah)

endprogram
```

**Contoh:**

| Input                | Output  | Penjelasan           |
| -------------------- | ------- | -------------------- |
| `3 2 0 1 2 1 1 -1`   | `true`  | total = 10, 10 >= 10 |
| `1 1 0 1 0 2 1 0 -1` | `false` | total = 6, 6 < 10    |

---

## 2. jumlahDeret

**Deskripsi:** Program membaca dua bilangan `n` dan `m`, lalu menjumlahkan suku-suku `4/i` untuk `i` dari `n` sampai `m`. Setiap kali satu suku ditambahkan, jumlah sementara langsung dicetak.

**Input:** `n` dan `m` (bilangan bulat)

**Output:** Jumlah sementara setelah tiap suku, masing-masing satu baris, sebanyak `m - n + 1` baris. Jika `n > m`, tidak ada yang dicetak.

**Rumus:** suku ke-i = 4 / i

**Catatan:**

- `i` harus diubah ke real sebelum pembagian agar hasilnya tidak dibulatkan menjadi bilangan bulat.
- Nilai `n` tidak boleh 0 dan rentang `n` sampai `m` tidak boleh memuat 0, karena terjadi pembagian dengan nol.
- Hasil real bisa memiliki selisih kecil akibat pembulatan komputer.

**Pseudocode:**

```
program jumlahDeret

kamus
    n      : integer    { batas awal }
    m      : integer    { batas akhir }
    i      : integer    { penghitung perulangan }
    jumlah : real       { jumlah sementara deret }

algoritma
    input(n, m)
    jumlah <- 0
    for i <- n to m do
        jumlah <- jumlah + 4 / real(i)
        output(jumlah)
    endfor

endprogram
```

**Contoh:**

| Input | Output                                            | Penjelasan                         |
| ----- | ------------------------------------------------- | ---------------------------------- |
| `1 3` | `4` lalu `6` lalu `7.333333333333333` (kira-kira) | 4/1 = 4, + 4/2 = 6, + 4/3 ≈ 7,3333 |

---

## 3. tabungan

**Deskripsi:** Kakak menabung setiap hari Minggu. Besar tabungan minggu pertama diberikan, dan tiap minggu berikutnya tabungan bertambah Rp 2500 dari tabungan minggu sebelumnya. Program membaca tabungan minggu pertama dan jumlah minggu `minggu`, lalu menghitung total tabungan setelah `minggu` minggu.

**Input:** Baris pertama: `tabungan` (tabungan minggu pertama). Baris kedua: `minggu` (banyaknya minggu).

**Output:** Total tabungan kakak setelah `minggu` minggu.

**Rumus:** tabungan minggu ke-i = tabungan + (i - 1) × 2500

**Cara kerja:** Perulangan berjalan dari minggu 1 sampai `minggu`. Pada tiap minggu, besar tabungan minggu itu dihitung langsung dari rumus di atas lalu ditambahkan ke `total`, sehingga tidak perlu variabel tambahan untuk menyimpan setoran.

**Catatan:** Pada file `tabungan.go`, variabel `total` dipakai tetapi belum dideklarasikan, sehingga program tidak dapat dikompilasi. Pseudocode di bawah mendeklarasikan `total` dan mengisinya 0 di awal. Perbaikan di Go: tambahkan `total := 0` sebelum perulangan.

**Pseudocode:**

```
program tabungan

kamus
    tabungan : integer    { tabungan minggu pertama }
    minggu   : integer    { banyaknya minggu }
    total    : integer    { total tabungan }
    i        : integer    { penghitung perulangan (minggu ke-i) }

algoritma
    input(tabungan, minggu)
    total <- 0
    for i <- 1 to minggu do
        total <- total + tabungan + (i - 1) * 2500
    endfor
    output(total)

endprogram
```

**Contoh:**

| Input      | Output  | Penjelasan                         |
| ---------- | ------- | ---------------------------------- |
| `4000` `1` | `4000`  | 4000                               |
| `4000` `2` | `10500` | 4000 + 6500                        |
| `4000` `3` | `19500` | 4000 + 6500 + 9000                 |
| `4000` `5` | `45000` | 4000 + 6500 + 9000 + 11500 + 14000 |
