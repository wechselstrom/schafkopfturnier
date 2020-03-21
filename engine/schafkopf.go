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
	if s.GetSpieltyp() != pb.Spieltyp_Wenz && k.GetWert() == pb.Wert_Ober {
		return true
	}
	if s.GetSpieltyp() != pb.Spieltyp_Geier && k.GetWert() == pb.Wert_Unter {
		return true
	}
	trumpffarbe := pb.Farbe_Herz
	switch s.GetSpieltyp() {
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
func get_sticht(s pb.Spiel) func(pb.Karte, pb.Karte) bool {
	return func(k1 pb.Karte, k2 pb.Karte) {
		t1, t2 := istTrumpf(s, k1), istTrumpf(s, k2)
		if t1 && t2 {
			return true //TODO: fix this
		}  else if t1 && !t2 {
			return true
		} else if !t1 && !t2 &&
		          k1.GetFarbe() == k2.GetFarbe() &&
		          k1 > k2 {
		       return true

		}
		return false
	}
}

//func werSticht(s pb.Spiel, ks pb.Karte[]) int {
//	return 5;
//}
