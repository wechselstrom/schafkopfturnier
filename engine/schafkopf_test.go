package main

import (
	"testing"
	pb "github.com/wechselstrom/schafkopfturnier/proto"
)

var schellen_solo = pb.Spiel{Spieltyp:pb.Spieltyp_Solo, Farbe:pb.Farbe_Schellen}
var grass_geier = pb.Spiel{Spieltyp:pb.Spieltyp_Geier, Farbe:pb.Farbe_Grass}
var wenz = pb.Spiel{Spieltyp:pb.Spieltyp_Wenz, Farbe:pb.Farbe_Natura}
var eichel_wenz = pb.Spiel{Spieltyp:pb.Spieltyp_Wenz, Farbe:pb.Farbe_Eichel}
var aufdblaue = pb.Spiel{Spieltyp:pb.Spieltyp_Sauspiel, Farbe:pb.Farbe_Grass}

var alte = pb.Karte{Wert:pb.Wert_Ober, Farbe:pb.Farbe_Eichel}
var rote = pb.Karte{Wert:pb.Wert_Ober, Farbe:pb.Farbe_Herz}
var die_blaue = pb.Karte{Wert:pb.Wert_Sau, Farbe:pb.Farbe_Grass}
var blau_wenz = pb.Karte{Wert:pb.Wert_Unter, Farbe:pb.Farbe_Grass}
var grass_7er = pb.Karte{Wert:pb.Wert_Sieben, Farbe:pb.Farbe_Grass}
var eichel_10er = pb.Karte{Wert:pb.Wert_Zehn, Farbe:pb.Farbe_Eichel}

func TestTrumpfwert(t *testing.T) {
    spiel := schellen_solo
    got := trumpfwert(spiel, alte)
    soll := 16
    if got != soll {
        t.Errorf("wert vom alten im %v = %d; want %d", spiel, got, soll)
    }
    spiel = eichel_wenz
    got = trumpfwert(spiel, alte)
    soll =  5
    if got != soll {
        t.Errorf("wert vom alten im %v = %d; want %d", spiel, got, soll)
    }
    spiel = schellen_solo
    got = trumpfwert(spiel, blau_wenz)
    soll =  11
    if got != soll {
        t.Errorf("wert vom blau wenz im %v = %d; want %d", spiel, got, soll)
    }
}

func __sticht(t *testing.T, spiel pb.Spiel, k1 pb.Karte, k2 pb.Karte, soll bool) {
	got := sticht(spiel, k1, k2)
	if got != soll {
	    t.Errorf("%v sticht %v im %v = %t; want %t", k1, k2, spiel, got, soll)
	}
}

func TestSticht(t *testing.T) {
    __sticht(t, schellen_solo, alte, rote, true)
    __sticht(t, schellen_solo, rote, alte, false)
    __sticht(t, schellen_solo, die_blaue, alte, false)
    __sticht(t, wenz, eichel_10er, alte, true)
    __sticht(t, wenz, eichel_10er, grass_7er, false)
    __sticht(t, wenz, grass_7er, eichel_10er, false)
    __sticht(t, grass_geier, grass_7er, eichel_10er, true)
    __sticht(t, aufdblaue, grass_7er, eichel_10er, false)
}

func TestWerSticht(t *testing.T) {

}
