package jabufaker

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Const alphabet for use random data with string
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Func init for first run
func init() {
	// Run rand.Seed
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	// Get total character on const alphabet
	k := len(alphabet)

	// Loop through n
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random person name
func RandomPerson() string {
	// Create list person name
	persons := []string{"Ari", "Ayu", "Aulia", "Anggi", "Agus", "Ade", "Arya", "Amel", "Andi", "Bayu", "Bagas", "Budi", "Bagus", "Bastian", "Ben", "Chika", "Cinta", "Citra", "Cakra", "Candra", "Darius", "Dimas", "Deo", "Dean", "Dinda", "Dika", "Dodi", "Ernes", "Erwin", "Eka", "Elin", "Elsa", "Ema", "Ela", "Fikri", "Fitri", "Fika", "Fani", "Fina", "Farid", "Fadel", "Galih", "Gading", "Guntur", "Gilang", "Geri", "Gibran", "Hamidah", "Hilda", "Hilmi", "Hisyam", "Haikal", "Harun", "Ita", "Ilham", "Indra", "Ikbal", "Irwan", "Ivan", "Irfan", "Ian", "Joko", "Josua", "Jonathan", "Jeri", "Jefri", "Karin", "Kirana", "Keisya", "Kevin", "Keyla", "Luna", "Lala", "Larisa", "Latif", "Lukman", "Mila", "Monika", "Maya", "Mira", "Malik", "Nila", "Nanda", "Naila", "Nisa", "Niko", "Nida", "Oki", "Okta", "Omar", "Oskar", "Olivia", "Putra", "Putri", "Paul", "Pinkan", "Pedro", "Qiqi", "Qafi", "Qori", "Qamar", "Queen", "Rafi", "Rafa", "Ririn", "Riska", "Rian", "Salsa", "Sinta", "Syifa", "Syahrul", "Samuel", "Tika", "Tristan", "Tobi", "Toni", "Tina", "Ulfa", "Usman", "Ulya", "Utari", "Umi", "Vania", "Virza", "Vincent", "Valdo", "Vino", "Wisnu", "Wulan", "Winda", "William", "Wira", "Yuda", "Yuli", "Yolanda", "Yusron", "Yosep"}

	// Get length person
	n := len(persons)

	// Return person name with rand mucn as length persons
	return persons[rand.Intn(n)]
}

// RandomMonet generates a random amount of money
func RandomMoney() int64 {
	// Use function RandomInt for get number from 0 - 1000
	return RandomInt(0, 100)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	// Create list of currency
	currencies := []string{"IDR", "USD", "EUR"}
	// Get len currency
	n := len(currencies)
	// Return currency with rand much as length currencies
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	var randInt int64
	// random int
	switch rand.Intn(3) {
	case 1:
		randInt = RandomInt(1001, 2000)
	case 2:
		randInt = RandomInt(2001, 3000)
	default:
		randInt = RandomInt(1, 1000)
	}

	// Random person
	person := strings.ToLower(RandomPerson())

	email := fmt.Sprintf("%s%d@test.com", person, randInt)
	return email
}

// RandomProvince generates a random some province of based on Sumatera Island
func RandomProvince() string {
	// Create list of province
	province := []string{"ACEH", "SUMATERA UTARA", "SUMATERA BARAT", "RIAU", "JAMBI"}
	// Get length of province
	n := len(province)
	// Return with rand much as length province
	return province[rand.Intn(n)]
}

// RandomRegency generates a random some regencies of based on selected from function `RandomProvince`
func RandomRegency(province string) string {
	var regencies []string

	switch province {
	case "ACEH":
		regencies = []string{"KABUPATEN SIMEULUE", "KABUPATEN ACEH SINGKIL", "KABUPATEN ACEH SELATAN", "KABUPATEN ACEH TENGGARA", "KABUPATEN ACEH TIMUR", "KABUPATEN ACEH TENGAH", "KABUPATEN ACEH BARAT", "KABUPATEN ACEH BESAR", "KABUPATEN PIDIE", "KABUPATEN BIREUEN", "KABUPATEN ACEH UTARA", "KABUPATEN ACEH BARAT DAYA", "KABUPATEN GAYO LUES", "KABUPATEN ACEH TAMIANG", "KABUPATEN NAGAN RAYA", "KABUPATEN ACEH JAYA", "KABUPATEN BENER MERIAH", "KABUPATEN PIDIE JAYA", "KOTA BANDA ACEH", "KOTA SABANG", "KOTA LANGSA", "KOTA LHOKSEUMAWE", "KOTA SUBULUSSALAM"}

	case "SUMATERA UTARA":
		regencies = []string{"KABUPATEN NIAS", "KABUPATEN MANDAILING NATAL", "KABUPATEN TAPANULI SELATAN", "KABUPATEN TAPANULI TENGAH", "KABUPATEN TAPANULI UTARA", "KABUPATEN TOBA SAMOSIR", "KABUPATEN LABUHAN BATU", "KABUPATEN ASAHAN", "KABUPATEN SIMALUNGUN", "KABUPATEN DAIRI", "KABUPATEN KARO", "KABUPATEN DELI SERDANG", "KABUPATEN LANGKAT", "KABUPATEN NIAS SELATAN", "KABUPATEN HUMBANG HASUNDUTAN", "KABUPATEN PAKPAK BHARAT", "KABUPATEN SAMOSIR", "KABUPATEN SERDANG BEDAGAI", "KABUPATEN BATU BARA", "KABUPATEN PADANG LAWAS UTARA", "KABUPATEN PADANG LAWAS", "KABUPATEN LABUHAN BATU SELATAN", "KABUPATEN LABUHAN BATU UTARA", "KABUPATEN NIAS UTARA", "KABUPATEN NIAS BARAT", "KOTA SIBOLGA", "KOTA TANJUNG BALAI", "KOTA TANJUNG BALAI", "KOTA PEMATANG SIANTAR", "KOTA TEBING TINGGI", "KOTA MEDAN", "KOTA BINJAI", "KOTA PADANGSIDIMPUAN", "KOTA GUNUNGSITOLI"}

	case "SUMATERA BARAT":
		regencies = []string{"KABUPATEN KEPULAUAN MENTAWAI", "KABUPATEN PESISIR SELATAN", "KABUPATEN SOLOK", "KABUPATEN SIJUNJUNG", "KABUPATEN TANAH DATAR", "KABUPATEN PADANG PARIAMAN", "KABUPATEN AGAM", "KABUPATEN LIMA PULUH KOTA", "KABUPATEN PASAMAN", "KABUPATEN SOLOK SELATAN", "KABUPATEN DHARMASRAYA", "KABUPATEN PASAMAN BARAT", "KOTA PADANG", "KOTA SOLOK", "KOTA SAWAH LUNTO", "KOTA PADANG PANJANG", "KOTA BUKITTINGGI", "KOTA PAYAKUMBUH", "KOTA PARIAMAN"}

	case "RIAU":
		regencies = []string{"KABUPATEN KUANTAN SINGINGI", "KABUPATEN INDRAGIRI HULU", "KABUPATEN INDRAGIRI HILIR", "KABUPATEN PELALAWAN", "KABUPATEN S I A K", "KABUPATEN KAMPAR", "KABUPATEN ROKAN HULU", "KABUPATEN BENGKALIS", "KABUPATEN ROKAN HILIR", "KABUPATEN KEPULAUAN MERANTI", "KOTA PEKANBARU", "KOTA D U M A I"}

	case "JAMBI":
		regencies = []string{"KABUPATEN KERINCI", "KABUPATEN MERANGIN", "KABUPATEN SAROLANGUN", "KABUPATEN BATANG HARI", "KABUPATEN MUARO JAMBI", "KABUPATEN TANJUNG JABUNG TIMUR", "KABUPATEN TANJUNG JABUNG BARAT", "KABUPATEN TEBO", "KABUPATEN BUNGO", "KOTA JAMBI", "KOTA SUNGAI PENUH"}

	default:
		regencies = []string{""}
	}

	// Get length of regencies
	n := len(regencies)

	// Return with rand much as length regencies
	return regencies[rand.Intn(n)]
}
