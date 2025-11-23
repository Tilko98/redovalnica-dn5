package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
	"github.com/Tilko98/redovalnica-dn5/redovalnica"

)

func main() {
	cmd := &cli.Command{
		Name:  "redovalnica",
		Usage: "upravljanje ocen študentov",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Usage: "minimalno število ocen za pozitivno oceno",
				Value: 6,
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Usage: "najmanjša dovoljena ocena",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Usage: "največja dovoljena ocena",
				Value: 10,
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			// preberemo vrednosti flagov
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

			// testni študenti
			studenti := map[string]redovalnica.Student{
				"63210001": {Ime: "Ana", Priimek: "Novak", Ocene: []int{10, 9, 8}},
				"63210002": {Ime: "Boris", Priimek: "Kralj", Ocene: []int{6, 7, 5, 8}},
				"63210003": {Ime: "Janez", Priimek: "Novak", Ocene: []int{4, 5, 3, 5}},
			}

			// uporabiš funkcije iz paketa redovalnica
			redovalnica.DodajOceno(studenti, "63210001", 10, minOcena, maxOcena)
			redovalnica.IzpisVsehOcen(studenti)
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

			//fmt.Println("CLI se je uspešno izvedel.")
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
