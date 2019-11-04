package main

import "gfx"
import "./Lampe"
//import "zufallszahlen"
//import "time"
//import "strconv"

type Spielfeld struct {
	Farbenfolge []uint8    //Slice von uin8-Werten, das die Anzeige-Reihenfolge der Lampen enthält
	                       //0=grün, 1=rot, 2=blau, 3=gelb
	bisherGeschafft uint8  //wie viele Lampen der Spieler bisher richtig gemerkt hat
	Zeigedauer uint64      //wie lange die Lampen beim Vorspielen aufleuchten, in ns
	Modus string           //für die Anzeige: "vorspielen" oder "nachspielen"
	grüne, rote, blaue, gelbe Lampe.Lampe //die vier Lapmen
}

func Initialisieren () Spielfeld {
	
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Die Farbenfolge ist mit 100 zufälligen erten zwischen 0 und 3 gefüllt.
//      bisherGeschafft ist auf 0 gesetzt.
//      Die vier Lampen sind initialisiert und aus.
//      Der Modus ist "vorspielen"
//      Das so initialisierte Spielfeld ist zurückgegeben.

	if gfx.FensterOffen != true {
	fmt.Println("Mach das Fenster erst mal auf!!!")
	}
	var b = make ([]uint8, 100)  //für das Festlegen der Farbenfolge
	
	s.grüne = Lampe.Initialisieren(10,10,100,0,255,0) 
	s.rot = Lampe.Initialisieren(120,10,100,255,0,0)
	s.gelb= Lampe.Initialisieren(10,120,100,0,255,255)
	s.blau= Lampe.Initialisieren(120,120,100,0,0,255)
	s.bisherGeschafft = 0
	s.Zeigedauer = 500000000 //in ms
	s.Modus = "Vorspielen"
	zufallszahlen.Randomisieren()
	for i:=0; i<len(b); i++{
		b[i] = uint8(zufallszahlen.NZufallszahl(0,3))
	}
	return s
}

func (s *Spielfeld)Zeichnen(){
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Das aktuelle Spielfeld ist auf dem Bildschirm ausgegeben.

}

func (s *Spielfeld)Vorspielen(){
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Es haben nacheinander die vom Spieler bisher richtig geklickten Lampen
//      plus eine zusätzliche aufgeleuchtet. Der Modus ist auf "vorspielen" gesetzt.

}

func (s*Spielfeld)Nachspielen()bool{
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Wenn der Spieler die Lampen in der richtigen Reihenfolge angeklickt hat,
//      so ist true zurück gegeben. 
//      Wenn nicht, so ist false zurückgegeben.
//      Der Modus ist auf "nachspielen" gesetzt.
   return true
}

func main(){
	gfx.Fenster(230, 230)
    gfx.TastaturLesen1()
}
