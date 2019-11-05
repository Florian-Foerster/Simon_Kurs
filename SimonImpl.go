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
}

func (s *Spielfeld)Zeichnen(){  //Üntzelmann, Förster
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Das aktuelle Spielfeld ist auf dem Bildschirm ausgegeben.
    gfx.Stiftfarbe(255,255,255)
    gfx.Vollrechteck(0,0,230,10)
	gfx.Stiftfarbe(0,0,0)
	gfx.Schreibe(0,0,s.Modus)
	s.grüne.Zeichnen()
	s.rote.Zeichnen()
	s.gelbe.Zeichnen()
	s.blaue.Zeichnen()
	gfx.Stiftfarbe(255,255,255)
	gfx.Vollrechteck(0,220,220,20)
	gfx.Stiftfarbe(0,0,0)
	gfx.Schreibe(0,220,"bisher geschafft:" + strconv.Itoa(int(s.bisherGeschafft)))
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
