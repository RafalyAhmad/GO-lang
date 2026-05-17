package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Student struct {
	KTM  string `json:"ktm"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var students []Student

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(students)

	case "POST":
		var student Student
		_ = json.NewDecoder(r.Body).Decode(&student)
		students = append(students, student)
		json.NewEncoder(w).Encode(student)

	case "PATCH":
		// 1. Ambil KTM dari URL Query (contoh: /students?ktm=2215061011)
		ktmTarget := r.URL.Query().Get("ktm")
		if ktmTarget == "" {
			http.Error(w, "Parameter 'ktm' dibutuhkan", http.StatusBadRequest)
			return
		}

		// 2. Ambil data perubahan dari Body Request
		var updatedData Student
		_ = json.NewDecoder(r.Body).Decode(&updatedData)

		// 3. Cari mahasiswa berdasarkan KTM di dalam slice
		found := false
		for i := 0; i < len(students); i++ {
			if students[i].KTM == ktmTarget {
				found = true

				// 4. Update hanya field yang dikirim (tidak kosong)
				if updatedData.Name != "" {
					students[i].Name = updatedData.Name
				}
				if updatedData.Age != 0 {
					students[i].Age = updatedData.Age
				}

				// Kirim respon data yang sudah berhasil diupdate
				json.NewEncoder(w).Encode(students[i])
				return
			}
		}

		// Jika KTM tidak ditemukan di dalam slice
		if !found {
			http.Error(w, "Mahasiswa dengan KTM tersebut tidak ditemukan", http.StatusNotFound)
		}

	case "DELETE":
		// 1. Ambil KTM yang mau dihapus dari Query Parameter (contoh: /students?ktm=2215061011)
		ktmYangDicari := r.URL.Query().Get("ktm")
		if ktmYangDicari == "" {
			http.Error(w, "Parameter 'ktm' wajib diisi", http.StatusBadRequest)
			return
		}

		indexKetemu := -1
		// 2. Cari posisinya di dalam slice students
		for i, student := range students {
			if student.KTM == ktmYangDicari {
				indexKetemu = i
				break
			}
		}

		// 3. Jika KTM tidak ditemukan
		if indexKetemu == -1 {
			http.Error(w, "Mahasiswa dengan KTM tersebut tidak ditemukan", http.StatusNotFound)
			return
		}

		// 4. Proses Hapus dari Slice
		// Caranya: gabungkan slice sebelum indexKetemu dan slice setelah indexKetemu
		students = append(students[:indexKetemu], students[indexKetemu+1:]...)

		// 5. Berikan respon sukses
		responSukses := map[string]string{"message": "Data mahasiswa berhasil dihapus"}
		json.NewEncoder(w).Encode(responSukses)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	students = make([]Student, 2)

	students[0] = Student{KTM: "2215061011", Name: "Ahmad Rafaly", Age: 22}
	students[1] = Student{KTM: "2215061012", Name: "John Doe", Age: 23}

	// isi data sementara
	students = append(students, Student{KTM: "2022", Name: "Nisya", Age: 22})
	students = append(students, Student{KTM: "2022", Name: "Ananda", Age: 22})
	students = append(students, Student{KTM: "2022", Name: "Shaliha", Age: 22})
	fmt.Println(students)

	http.HandleFunc("/students", GetStudents)
	log.Println("Server starting on localhost:8080")
	http.ListenAndServe("localhost:8080", nil)

}
