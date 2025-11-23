// Package redovalnica omogoča delo z ocenami študentov in izpis redovalnice.
package redovalnica

import "fmt"

// Student predstavlja študenta z imenom, priimkom in seznamom ocen
type Student struct {
Ime     string
Priimek string
Ocene   []int
}

// DodajOceno doda oceno študentu, če je v dovoljenem območju
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, minOcena int, maxOcena int) {
    // Preveri, ali je ocena v ustreznem območju
    if ocena < minOcena || ocena > maxOcena {
        fmt.Printf("Napaka: Ocena %d ni v veljavnem območju (0-10)\n", ocena)
        return
    }
    
    // Preveri, ali študent obstaja
    student, obstaja := studenti[vpisnaStevilka]
    if !obstaja {
        fmt.Println("Napaka: Študent z vpisno številko", vpisnaStevilka, "ne obstaja na seznamu")
        return
    }
    
    // Dodaj oceno študentu
    student.Ocene = append(student.Ocene, ocena)
    studenti[vpisnaStevilka] = student
    //fmt.Printf("Študentu %s %s je bila dodana ocena %d\n", student.ime, student.priimek, ocena)
}

func povprecje(studenti map[string]Student, vpisnaStevilka string, stOcen int) float64 {
    student, obstaja := studenti[vpisnaStevilka]
    if !obstaja {
        fmt.Println("Napaka: Študent z vpisno številko", vpisnaStevilka, "ne obstaja na seznamu")
        return -1.0
    }

    if len(student.Ocene) < stOcen{
        return 0.0
    }

    suma := 0
    for _, ocena := range student.Ocene {
        suma += ocena
    }

    avg := float64(suma) / float64(len(student.Ocene))
    if avg < 6.0 {
        return 0.0
    }
    return avg

}

// IzpisVsehOcen izpiše vse študente in njihove ocene.
func IzpisVsehOcen(studenti map[string]Student) {
    fmt.Println("Redovalnica:")
    for vpisnaStevilka, s := range studenti{
        fmt.Printf("%s - %s %s: %v\n", vpisnaStevilka, s.Ime, s.Priimek, s.Ocene)
    }
}

// IzpisiKoncniUspeh izpiše končni uspeh študentov glede na povprečje.
func IzpisiKoncniUspeh(studenti map[string]Student, stOcen int) {
    for vpisnaStevilka, s := range studenti {
        avg := povprecje(studenti, vpisnaStevilka, stOcen)
        if avg >= 9 {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Odličen študent!\n", s.Ime, s.Priimek, avg)
        } else if avg >= 6 && avg < 9 {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Povprečen študent\n", s.Ime, s.Priimek, avg)
        } else {
            fmt.Printf("%s %s: povprečna ocena %.2f -> Neuspešen študent\n", s.Ime, s.Priimek, avg)
        }
    }
}
