package main

import "fmt"

func main() {
	modules := map[int]string{
		104: "Grundlagen der Programmierung",
		117: "Web-Applikationen entwickeln",
		346: "Cloud-Lösungen konzipieren und realisieren",
	}

	fmt.Println("Modul 104:", modules[104])
	fmt.Println("Modul 117:", modules[117])
	fmt.Println("Modul 346:", modules[346])

	delete(modules, 117)
	//TO-DO löschen
	modules[201] = "Mobile Anwendungen entwickeln"
	//TO-DO hinzufügen
	modules[346] = "Cloud-Architekturen designen"

	fmt.Println(modules)
}
