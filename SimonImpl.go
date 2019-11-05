package main

//Some stupid comment
//Irgendwas Sinnloses von Schüler x 

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
}

func (s *Spielfeld)Zeichnen(){
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Das aktuelle Spielfeld ist auf dem Bildschirm ausgegeben.

}

func (s *Spielfeld)Vorspielen(){
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Es haben nacheinander die vom Spieler bisher richtig geklickten Lampen
//      plus eine zusätzliche aufgeleuchtet. Der Modus ist auf "vorspielen" gesetzt.
	s.Modus = "vorspielen"
	s.Zeichnen()
	//s.bisherGeschafft++	besser am Ende von Nachspielen (sonst am Anfang von Nachspielen falsche Anzeige)
	for i:=uint8(0); i<=(*s).bisherGeschafft;i++ {
		switch (*s).Farbenfolge[i] {
		case 0:
		s.grüne.Umschalten()
		s.grüne.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer))
		s.grüne.Umschalten()
		s.grüne.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer/2))
		case 1:
		s.rote.Umschalten()
		s.rote.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer))
		s.rote.Umschalten()
		s.rote.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer/2))
		case 2:
		s.blaue.Umschalten()
		s.blaue.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer))
		s.blaue.Umschalten()
		s.blaue.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer/2))
		case 3:
		s.gelbe.Umschalten()
		s.gelbe.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer))
		s.gelbe.Umschalten()
		s.gelbe.Zeichnen()
		time.Sleep(time.Duration(s.Zeigedauer/2))
		}
	}
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
