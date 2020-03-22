package main

import (
	pb "github.com/wechselstrom/schafkopfturnier/proto"
)



var wert_zu_augen = map[pb.Wert]int{
	pb.Wert_Sieben:  0,
	pb.Wert_Acht:    0,
	pb.Wert_Neun:    0,
	pb.Wert_Zehn:   10,
	pb.Wert_Unter:   2,
	pb.Wert_Ober:    3,
	pb.Wert_Koenig:  4,
	pb.Wert_Sau:    11,
}


func istTrumpf(s pb.Spiel, k pb.Karte) bool {
	st := s.GetSpieltyp()
	if st != pb.Spieltyp_Wenz && k.GetWert() == pb.Wert_Ober {
		return true
	}
	if st != pb.Spieltyp_Geier && k.GetWert() == pb.Wert_Unter {
		return true
	}
	trumpffarbe := pb.Farbe_Herz
	switch st {
		case
		    pb.Spieltyp_Solo,
		    pb.Spieltyp_Wenz,
		    pb.Spieltyp_Geier:
		    trumpffarbe = s.GetFarbe()
	}
	if trumpffarbe == k.GetFarbe() {
		return true
	}
	return false
}
func trumpfwert(s pb.Spiel, k pb.Karte) int {
	st := s.GetSpieltyp()
	if st != pb.Spieltyp_Wenz &&
	   k.GetWert() == pb.Wert_Ober {
		return int(12+k.GetFarbe())
	} else if st != pb.Spieltyp_Geier &&
	   k.GetWert() == pb.Wert_Unter {
		return int(8+k.GetFarbe())
	} else {
		return int(k.GetWert())
	}
}

func sticht(s pb.Spiel, k1 pb.Karte, k2 pb.Karte) bool {
	t1, t2 := istTrumpf(s, k1), istTrumpf(s, k2)
	if t1 && t2 {
		return trumpfwert(s, k1) > trumpfwert(s, k2)
	} else if t1 && !t2 {
		return true
	} else if !t1 && !t2 &&
	          k1.GetFarbe() == k2.GetFarbe() &&
	          k1.GetWert() > k2.GetWert() {
	       return true

	}
	return false
}

func werSticht(s pb.Spiel, ks []pb.Karte) int {
	current_highest := 0
	for i, v := range ks[1:] {
		if sticht(s, v, ks[i]) {
			current_highest = i
		}
	}
	return current_highest
}

func darfErsteGespieltWerden(s pb.Spiel, erste pb.Karte, hand []pb.Karte, davongelaufen bool) bool {
	return !RufZwangMissachtet(s, erste, erste, hand, davongelaufen)
}

func darfNachfolgendeGespieltWerden(s pb.Spiel, erste pb.Karte, naechste pb.Karte, hand []pb.Karte, davongelaufen bool) bool {
	verboten := TrumpfZwangMissachtet(s, erste, naechste, hand) ||
	            FarbZwangMissachtet(s, erste, naechste, hand) ||
	            RufZwangMissachtet(s, erste, naechste, hand, davongelaufen) ||
	            RufSchmierVerbotMissachtet(s, erste, naechste, hand, davongelaufen)
	return !verboten


}
func TrumpfZwangMissachtet(s pb.Spiel, erste pb.Karte, naechste pb.Karte, hand []pb.Karte) bool {
	te, tn := istTrumpf(s, erste), istTrumpf(s, naechste)
	if !te || //erste ist kein Trumpf also gilt kein Trumpfzwang
	   tn  || //trumpf wurde zugegeben
	   !inHand(hand, istTrumpf) { // Kein Trumpf auf der Hand, also entfällt Trumpfzwang
		return false
	}
	return true
}

func FarbZwangMissachtet(s pb.Spiel, erste pb.Karte, naechste pb.Karte, hand []pb.Karte) bool {
	if !inHand(hand, istFarbe(s, handkarte, erste.GetFarbe())) ||// Keine passende Farbe auf der Hand, also entfällt Farbzwang
	   istFarbe(s, naechste, erste.GetFarbe()) {
		//farbe wurde zugegeben
		return false
	}
	return true
}

func RufZwangMissachtet(s pb.Spiel, erste pb.Karte, naechste pb.Karte, hand []pb.Karte, davongelaufen bool) bool {
	if pb.Spiel.GetSpieltyp() != pb.Spieltyp_Sauspiel || //nur im Sauspiel 
	   istFarbe(s, erste, s.GetFarbe()) || //wenn die ruffarbe angespielt wird relevant
	   !inHand(hand, pb.Karte{Wert:pb.Wert_Sau, Farbe:s.GetFarbe()}) || // Sau nicht auf der Hand also entfällt Rufzwang
	   naechste == (pb.Karte{Wert:pb.Wert_Sau, Farbe:erste.GetFarbe()}) || // Es wird gerade zugegeben
	   DarfDavonLaufen(s, hand) || davongelaufen { // darf davonlaufen oder ist bereits davongelaufen, rufzwang entfällt.
		return false
	}
	return true
}

func RufSchmierVerbotMissachtet(s pb.Spiel, erste pb.Karte, naechste pb.Karte, hand []pb.Karte, davongelaufen bool) bool {
	if naechste != (pb.Karte{Wert:pb.Wert_sau, Farbe:s.GetFarbe()}) || // nicht die Sau der Spielfarbe
	  pb.Spiel.GetSpieltyp() != pb.Spieltyp_Sauspiel || // kein Sauspiel
	  !istFarbe(s, erste, s.GetFarbe()) || //es wurde gesucht
	  davongelaufen || //bereits davongelaufen
	  len(hand) == 1 { // letzte Handkarte
		return false
	}
	return true
}

