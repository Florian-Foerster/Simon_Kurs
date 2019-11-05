package main

//Some stupid comment
//Irgendwas Sinnloses von Schüler x 

import "gfx"
import "./Lampe"
import "zufallszahlen"
import "time"
import "strconv"

type Spielfeld struct {
	Farbenfolge []uint8    //Slice von uin8-Werten, das die Anzeige-Reihenfolge der Lampen enthält
	                       //0=grün, 1=rot, 2=blau, 3=gelb
	bisherGeschafft uint8  //wie viele Lampen der Spieler bisher richtig gemerkt hat
	Zeigedauer uint64      //wie lange die Lampen beim Vorspielen aufleuchten, in ns
	Modus string           //für die Anzeige: "vorspielen" oder "nachspielen"
	grüne, rote, blaue, gelbe Lampe.Lampe //die vier Lapmen
}

func Initialisieren () Spielfeld{  //Wendt, Opitz
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Die Farbenfolge ist mit 100 zufälligen erten zwischen 0 und 3 gefüllt.
//      bisherGeschafft ist auf 0 gesetzt.
//      Die vier Lampen sind initialisiert und aus.
//      Der Modus ist "vorspielen"
//      Das so initialisierte Spielfeld ist zurückgegeben.
	if gfx.FensterOffen() != true {
	fmt.Println("Mach das Fenster erst mal auf!!!")
	}
	var s Spielfeld
	s.Farbenfolge = make ([]uint8, 100)  //für das Festlegen der Farbenfolge
	s.grüne = Lampe.Initialisieren(10,10,100,0,255,0) 
	s.rote = Lampe.Initialisieren(120,10,100,255,0,0)
	s.gelbe= Lampe.Initialisieren(10,120,100,0,255,255)
	s.blaue= Lampe.Initialisieren(120,120,100,0,0,255)
	s.bisherGeschafft = 0
	s.Zeigedauer = 500000000 //in ms
	s.Modus = "Vorspielen"
	zufallszahlen.Randomisieren()
	for i:=0; i<len(s.Farbenfolge); i++{
		s.Farbenfolge[i] = uint8(zufallszahlen.Zufallszahl(0,3))
	}
	return s
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

func (s *Spielfeld) Nachspielen () bool {
//Vor.: Ein Grafikfenster ist geöffnet.
//Eff.: Wenn der Spieler die Lampen in der richtigen Reihenfolge angeklickt hat,
//      so ist true zurück gegeben. 
//      Wenn nicht, so ist false zurückgegeben.
//      Der Modus ist auf "nachspielen" gesetzt.
	var t uint8
	var st int8
	var x,y uint16
	(*s).Modus="Nachspielen"
	s.Zeichnen()
	for i:=0;i<=int((*s).bisherGeschafft);i++ {
		s.Zeichnen()
		for {
			t,st,x,y= gfx.MausLesen1()	
			if t==1 && st==1 {break}
		}
		s.Zeichnen()
		if (*s).rote.IstGeklickt(1,1,x,y) {
				(*s).rote.Umschalten()
				s.Zeichnen()
				time.Sleep(time.Duration((*s).Zeigedauer))
				(*s).rote.Umschalten()
				s.Zeichnen()				
			if (*s).Farbenfolge[i]!=1 {return false}
		} else if (*s).grüne.IstGeklickt(1,1,x,y) {
				(*s).grüne.Umschalten()
				s.Zeichnen()
				time.Sleep(time.Duration((*s).Zeigedauer))
				(*s).grüne.Umschalten()
				s.Zeichnen()	
				
			if (*s).Farbenfolge[i]!=0 {return false}
		} else if (*s).blaue.IstGeklickt(1,1,x,y) {
				(*s).blaue.Umschalten()
				s.Zeichnen()
				time.Sleep(time.Duration((*s).Zeigedauer))
				(*s).blaue.Umschalten()
				s.Zeichnen()	
				
			if (*s).Farbenfolge[i]!=2 {return false}
		} else if (*s).gelbe.IstGeklickt(1,1,x,y) {
				(*s).gelbe.Umschalten()
				s.Zeichnen()
				time.Sleep(time.Duration((*s).Zeigedauer))
				(*s).gelbe.Umschalten()
				s.Zeichnen()	
			if (*s).Farbenfolge[i]!=3 {return false}
		} else {
			i--
		}
	}
	time.Sleep(time.Duration((*s).Zeigedauer))
	(*s).bisherGeschafft++
	//(*s).Zeigedauer-1000
	return true
}

/// Bla
func main(){
	gfx.Fenster(230, 230)
    	var Simon Spielfeld
    	Simon=Initialisieren()
    	var beendet bool;
    	beendet=false
    	for beendet == false {
       Simon.Vorspielen()
       beendet= ! Simon.Nachspielen()
    }
}
