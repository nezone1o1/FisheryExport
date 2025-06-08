package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

type Ikan struct {
	Nama   string
	Jumlah int
	Berat  float64
}

type Ekspor struct {
	NamaIkan     string
	DaerahTujuan string
	Tanggal      string
	JumlahEkspor int
	BeratEkspor  float64
	Pendapatan   float64
}

type SuhuLokasi struct {
	Daerah       string
	SuhuRata     float64
	TanggalCatat string
}

type KondisiEkosistem struct {
	Daerah             string
	Status             string
	TingkatOverfishing string
	TanggalEvaluasi    string
}

const MAX = 100

var dataIkan [MAX]Ikan
var jumlahData int

var dataEkspor [MAX]Ekspor
var jumlahEkspor int

var dataSuhu [MAX]SuhuLokasi
var jumlahSuhu int

var dataEkosistem [MAX]KondisiEkosistem
var jumlahEkosistem int

func inputData() {
	if jumlahData >= MAX {
		fmt.Println("Data penuh!")
		return
	}

	fmt.Print("Masukkan nama ikan: ")
	reader.ReadString('\n')
	dataIkan[jumlahData].Nama = readLine()

	fmt.Print("Masukkan jumlah tangkapan (ekor): ")
	fmt.Scan(&dataIkan[jumlahData].Jumlah)

	fmt.Print("Masukkan berat total (kg): ")
	fmt.Scan(&dataIkan[jumlahData].Berat)

	jumlahData++
}

func tampilData() {
	fmt.Println("\nData Hasil Tangkapan Ikan:")
	for i := 0; i < jumlahData; i++ {
		ikan := dataIkan[i]
		fmt.Printf("%d. %s\n", i+1, ikan.Nama)
		fmt.Printf("Jumlah : %d ekor\n", ikan.Jumlah)
		fmt.Printf("Berat total : %.2fkg\n\n", ikan.Berat)
	}
}

func inputEkspor() {
	if jumlahEkspor >= MAX {
		fmt.Println("Data ekspor penuh!")
		return
	}

	fmt.Print("Masukkan nama ikan untuk diekspor: ")
	reader.ReadString('\n')
	nama := readLine()
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Ikan tidak ditemukan dalam data tangkapan.")
		return
	}

	fmt.Print("Masukkan daerah tujuan: ")
	dataEkspor[jumlahEkspor].DaerahTujuan = readLine()

	fmt.Print("Masukkan tanggal ekspor (DD-MM-YYYY): ")
	dataEkspor[jumlahEkspor].Tanggal = readLine()

	fmt.Print("Masukkan jumlah ekspor (ekor): ")
	fmt.Scan(&dataEkspor[jumlahEkspor].JumlahEkspor)

	fmt.Print("Masukkan berat ekspor (kg): ")
	fmt.Scan(&dataEkspor[jumlahEkspor].BeratEkspor)

	fmt.Print("Masukkan total pendapatan ekspor (Rp): ")
	fmt.Scan(&dataEkspor[jumlahEkspor].Pendapatan)

	if dataEkspor[jumlahEkspor].JumlahEkspor > dataIkan[idx].Jumlah {
		fmt.Println("Jumlah ekspor melebihi jumlah tangkapan!")
		return
	}

	dataEkspor[jumlahEkspor].NamaIkan = nama
	dataIkan[idx].Jumlah -= dataEkspor[jumlahEkspor].JumlahEkspor
	jumlahEkspor++
	fmt.Println("Data ekspor berhasil ditambahkan.")
}

func tampilEkspor() {
	fmt.Println("\nData Ekspor Ikan:")
	for i := 0; i < jumlahEkspor; i++ {
		e := dataEkspor[i]
		fmt.Printf("%d. %s\n", i+1, e.NamaIkan)
		fmt.Printf("Daerah Tujuan : %s\n", e.DaerahTujuan)
		fmt.Printf("Tanggal : %s\n", e.Tanggal)
		fmt.Printf("Jumlah Ekspor : %d\n", e.JumlahEkspor)
		fmt.Printf("Berat Ekspor : %.2f kg\n", e.BeratEkspor)
		fmt.Printf("Pendapatan : RP%.2f\n\n", e.Pendapatan)
	}
}
func inputSuhuLokasi() {
	if jumlahSuhu >= MAX {
		fmt.Println("Data suhu penuh!")
		return
	}

	fmt.Print("Masukkan nama daerah lokasi tangkap: ")
	reader.ReadString('\n')
	dataSuhu[jumlahSuhu].Daerah = readLine()

	fmt.Print("Masukkan suhu rata-rata (°C): ")
	fmt.Scanln(&dataSuhu[jumlahSuhu].SuhuRata)

	fmt.Print("Masukkan tanggal pencatatan (DD-MM-YYYY): ")
	dataSuhu[jumlahSuhu].TanggalCatat = readLine()

	jumlahSuhu++
	fmt.Println("Data suhu berhasil disimpan.")
}

func tampilSuhuLokasi() {
	fmt.Println("\nData Suhu Lokasi Tangkap:")
	for i := 0; i < jumlahSuhu; i++ {
		s := dataSuhu[i]
		fmt.Printf("%d. Daerah: %s\n", i+1, s.Daerah)
		fmt.Printf("   Suhu Rata-rata: %.2f°C\n", s.SuhuRata)
		fmt.Printf("   Tanggal Pencatatan: %s\n\n", s.TanggalCatat)
	}
}

func inputEkosistem() {
	if jumlahEkosistem >= MAX {
		fmt.Println("Data ekosistem penuh!")
		return
	}

	fmt.Print("Masukkan nama daerah: ")
	reader.ReadString('\n')
	dataEkosistem[jumlahEkosistem].Daerah = readLine()

	fmt.Print("Masukkan status lingkungan (Baik/Tercemar/Kritis): ")
	dataEkosistem[jumlahEkosistem].Status = readLine()

	fmt.Print("Masukkan tingkat overfishing (Rendah/Sedang/Tinggi): ")
	dataEkosistem[jumlahEkosistem].TingkatOverfishing = readLine()

	fmt.Print("Masukkan tanggal evaluasi (DD-MM-YYYY): ")
	dataEkosistem[jumlahEkosistem].TanggalEvaluasi = readLine()

	jumlahEkosistem++
	fmt.Println("Data kondisi ekosistem berhasil disimpan.")
}

func tampilEkosistem() {
	fmt.Println("\nData Kondisi Ekosistem dan Lingkungan:")
	for i := 0; i < jumlahEkosistem; i++ {
		e := dataEkosistem[i]
		fmt.Println(e.Daerah)
		fmt.Printf("status : %s \n", e.Status)
		fmt.Printf("OverFishing : %s \n", e.TingkatOverfishing)
		fmt.Printf("Tanggal : %s \n", e.TanggalEvaluasi)
	}
}

func insertionSortJumlahIkan() {
	for i := 1; i < jumlahData; i++ {
		temp := dataIkan[i]
		j := i - 1
		for j >= 0 && dataIkan[j].Jumlah > temp.Jumlah {
			dataIkan[j+1] = dataIkan[j]
			j--
		}
		dataIkan[j+1] = temp
	}
	fmt.Println("Data ikan telah diurutkan berdasarkan jumlah (menaik) menggunakan Insertion Sort.")
}

func selectionSortPendapatanEkspor() {
	for i := 0; i < jumlahEkspor-1; i++ {
		maxIdx := i
		for j := i + 1; j < jumlahEkspor; j++ {
			if dataEkspor[j].Pendapatan > dataEkspor[maxIdx].Pendapatan {
				maxIdx = j
			}
		}
		if maxIdx != i {
			dataEkspor[i], dataEkspor[maxIdx] = dataEkspor[maxIdx], dataEkspor[i]
		}
	}
	fmt.Println("Data ekspor telah diurutkan berdasarkan pendapatan (menurun) menggunakan Selection Sort.")
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahData; i++ {
		if strings.EqualFold(dataIkan[i].Nama, nama) {
			return i
		}
	}
	return -1
}

func sequentialSearchMultiple(nama string) {
	nama = strings.ToLower(nama)
	ditemukan := false

	fmt.Printf("Hasil pencarian ikan yang mengandung '%s':\n", nama)
	for i := 0; i < jumlahData; i++ {
		if strings.Contains(strings.ToLower(dataIkan[i].Nama), nama) {
			fmt.Printf("Nama: %s, Jumlah: %d, Berat: %.2fkg\n", dataIkan[i].Nama, dataIkan[i].Jumlah, dataIkan[i].Berat)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ditemukan ikan dengan nama yang mengandung:", nama)
	}
}

func insertionSortSuhu() {
	for i := 1; i < jumlahSuhu; i++ {
		temp := dataSuhu[i]
		j := i - 1
		for j >= 0 && dataSuhu[j].SuhuRata > temp.SuhuRata {
			dataSuhu[j+1] = dataSuhu[j]
			j--
		}
		dataSuhu[j+1] = temp
	}
}

func binarySearchSuhu(target float64) int {
	low := 0
	high := jumlahSuhu - 1
	for low <= high {
		mid := (low + high) / 2
		if dataSuhu[mid].SuhuRata == target {
			return mid
		} else if dataSuhu[mid].SuhuRata < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func carilokasiBerdasarkanSuhu() {
	if jumlahSuhu == 0 {
		fmt.Println("Data suhu kosong.")
		return
	}

	var target float64
	fmt.Print("Masukkan suhu rata-rata yang ingin dicari: ")
	fmt.Scan(&target)

	insertionSortSuhu()
	idx := binarySearchSuhu(target)

	if idx == -1 {
		fmt.Println("Tidak ditemukan lokasi dengan suhu rata-rata tersebut.")
		return
	}

	i := idx
	for i >= 0 && dataSuhu[i].SuhuRata == target {
		i--
	}
	i++

	fmt.Println("\nData ditemukan:")
	for i < jumlahSuhu && dataSuhu[i].SuhuRata == target {
		fmt.Printf("Daerah: %s\n", dataSuhu[i].Daerah)
		fmt.Printf("Suhu: %.2f°C\n", dataSuhu[i].SuhuRata)
		fmt.Printf("Tanggal Pencatatan: %s\n\n", dataSuhu[i].TanggalCatat)
		i++
	}
}

func menu() {
	for {
		fmt.Println("\n===== MENU UTAMA =====")
		fmt.Println("1. Input Data Ikan")
		fmt.Println("2. Tampilkan Data Ikan")
		fmt.Println("3. Input Data Ekspor")
		fmt.Println("4. Tampilkan Data Ekspor")
		fmt.Println("5. Input Suhu Lokasi")
		fmt.Println("6. Tampilkan Suhu Lokasi")
		fmt.Println("7. Input Kondisi Ekosistem")
		fmt.Println("8. Tampilkan Kondisi Ekosistem")
		fmt.Println("9. Urutkan Jumlah Ikan (Insertion Sort)")
		fmt.Println("10. Urutkan Pendapatan Ekspor (Selection Sort)")
		fmt.Println("11. Cari Lokasi Berdasarkan Suhu (Binary Search)")
		fmt.Println("12. Cari Ikan (Kemiripan Nama - Sequential Search)")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilih int
		fmt.Scan(&pilih)
		switch pilih {
		case 1:
			inputData()
		case 2:
			tampilData()
		case 3:
			inputEkspor()
		case 4:
			tampilEkspor()
		case 5:
			inputSuhuLokasi()
		case 6:
			tampilSuhuLokasi()
		case 7:
			inputEkosistem()
		case 8:
			tampilEkosistem()
		case 9:
			insertionSortJumlahIkan()
		case 10:
			selectionSortPendapatanEkspor()
		case 11:
			carilokasiBerdasarkanSuhu()
		case 12:
			fmt.Print("Masukkan kata kunci nama ikan: ")
			reader.ReadString('\n')
			nama := readLine()
			sequentialSearchMultiple(nama)
		case 0:
			fmt.Println("Terima kasih!")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	menu()
}
