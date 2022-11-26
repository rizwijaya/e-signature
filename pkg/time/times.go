package times

import (
	"fmt"
	"time"
)

var hari = [...]string{
	"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"}

var bulan = [...]string{
	"Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "Nopember", "Desember",
}

func TanggalJam(t time.Time) string {
	return fmt.Sprintf("%s, %02d %s %d | %02d:%02d WIB",
		hari[t.Weekday()], t.Day(), bulan[t.Month()-1][:3], t.Year(), t.Hour(), int(t.Minute()),
	)
}

func Tanggal(t time.Time) string {
	return fmt.Sprintf("%02d %s %d",
		t.Day(), bulan[t.Month()-1], t.Year())
}

func Jam(t time.Time) string {
	return fmt.Sprintf("%02d:%02d",
		t.Hour(), t.Minute())
}
