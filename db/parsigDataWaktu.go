package db

import (
	"sort"
	"strconv"
	"strings"
)

func SortDate(list []ListMoneyBuy) {
	sort.Slice(list, func(i, j int) bool {

		// Data pertama
		p1 := strings.Split(list[i].Date, " | ")
		tgl1 := strings.Split(p1[0], "-")
		jam1 := strings.Split(p1[1], ":")

		hari1, _ := strconv.Atoi(tgl1[0])
		bulan1, _ := strconv.Atoi(tgl1[1])
		tahun1, _ := strconv.Atoi(tgl1[2])

		jamPertama, _ := strconv.Atoi(jam1[0])
		menit1, _ := strconv.Atoi(jam1[1])
		detik1, _ := strconv.Atoi(jam1[2])

		// Data kedua
		p2 := strings.Split(list[j].Date, " | ")
		tgl2 := strings.Split(p2[0], "-")
		jam2 := strings.Split(p2[1], ":")

		hari2, _ := strconv.Atoi(tgl2[0])
		bulan2, _ := strconv.Atoi(tgl2[1])
		tahun2, _ := strconv.Atoi(tgl2[2])

		jamKedua, _ := strconv.Atoi(jam2[0])
		menit2, _ := strconv.Atoi(jam2[1])
		detik2, _ := strconv.Atoi(jam2[2])

		// Bandingkan
		if tahun1 != tahun2 {
			return tahun1 > tahun2
		}

		if bulan1 != bulan2 {
			return bulan1 > bulan2
		}

		if hari1 != hari2 {
			return hari1 > hari2
		}

		if jamPertama != jamKedua {
			return jamPertama > jamKedua
		}

		if menit1 != menit2 {
			return menit1 > menit2
		}

		return detik1 > detik2
	})
}
