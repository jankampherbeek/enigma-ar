/*
 *  Enigma Astrology Research.
 *  Copyright (c) Jan Kampherbeek.
 *  Enigma is open source.
 *  Please check the file copyright.txt in the root of the source for further details.
 */

package meta

import (
	"enigma-ar/domain"
	"image/color"
)

// DefaultConfig returns the default configuration
func DefaultConfig() domain.Config {
	basic := createBasic()
	orbs := createOrbs()
	aspects := createAspects()
	points := createPoints()
	prog := createProg()
	return domain.Config{
		Basic:   basic,
		Orbs:    orbs,
		Aspects: aspects,
		Points:  points,
		Prog:    prog,
	}

}

func createBasic() domain.ConfigBasic {
	return domain.ConfigBasic{
		Houses:   domain.HousesPlacidus,
		Ayan:     domain.AyanNone,
		ObsPos:   domain.ObsPosGeocentric,
		ProjType: domain.ProjType2D,
		Wheel:    domain.WheelTypeSignsEqual,
	}
}

func createOrbs() domain.ConfigOrbs {
	return domain.ConfigOrbs{
		BaseOrbAspects:   10,
		BaseOrbMidpoints: 1.6,
		OrbDeclMidpoints: 0.5,
		OrbParallels:     1.0,
		OrbTransits:      1.0,
		OrbSecDir:        1.0,
		OrbSymDir:        1.0,
		OrbPrimDir:       1.0,
	}
}

func createAspects() []domain.ConfigAspect {
	var (
		red = color.NRGBA{
			R: 255,
			G: 0,
			B: 0,
			A: 255,
		}
		green = color.NRGBA{
			R: 0,
			G: 255,
			B: 0,
			A: 255,
		}
		gray = color.NRGBA{
			R: 128,
			G: 128,
			B: 128,
			A: 255,
		}
		white = color.NRGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	)

	return []domain.ConfigAspect{
		createSpecAspect(domain.Conjunction, true, false, 100.0, '\uE700', white),
		createSpecAspect(domain.Opposition, true, true, 100.0, '\uE710', red),
		createSpecAspect(domain.Trine, true, true, 80.0, '\uE720', green),
		createSpecAspect(domain.Square, true, true, 80.0, '\uE730', red),
		createSpecAspect(domain.Sextile, true, true, 60.0, '\uE740', green),
		createSpecAspect(domain.SemiSextile, true, false, 20.0, '\uE750', gray),
		createSpecAspect(domain.Inconjunct, true, true, 30.0, '\uE760', gray),
		createSpecAspect(domain.SemiSquare, true, false, 20.0, '\uE770', gray),
		createSpecAspect(domain.SesquiQuadrate, true, false, 20.0, '\uE780', gray),
		createSpecAspect(domain.Quintile, true, false, 20.0, '\uE790', gray),
		createSpecAspect(domain.BiQuintile, true, false, 20.0, '\uE800', gray),
		createSpecAspect(domain.Septile, true, false, 20.0, '\uE810', gray),
		createSpecAspect(domain.Vigintile, false, false, 10.0, '\uE820', gray),
		createSpecAspect(domain.SemiQuintile, false, false, 10.0, '\uE830', gray),
		createSpecAspect(domain.TriDecile, false, false, 10.0, '\uE840', gray),
		createSpecAspect(domain.BiSeptile, false, false, 10.0, '\uE850', gray),
		createSpecAspect(domain.TriSeptile, false, false, 10.0, '\uE860', gray),
		createSpecAspect(domain.Novile, false, false, 10.0, '\uE870', gray),
		createSpecAspect(domain.BiNovile, false, false, 10.0, '\uE880', gray),
		createSpecAspect(domain.QuadraNovile, false, false, 10.0, '\uE890', gray),
		createSpecAspect(domain.Undecile, false, false, 10.0, '\uE900', gray),
		createSpecAspect(domain.Centile, false, false, 10.0, '\uE910', gray),
	}
}

func createSpecAspect(actAsp domain.Aspect, isUsed, show bool, orbF float64, glyph rune, col color.NRGBA) domain.ConfigAspect {
	return domain.ConfigAspect{
		ActualAspect: actAsp,
		IsUsed:       isUsed,
		ShowInChart:  show,
		OrbFactor:    orbF,
		Glyph:        glyph,
		Color:        col,
	}
}

func createPoints() []domain.ConfigPoint {
	return []domain.ConfigPoint{
		createSpecPoint(domain.Sun, true, true, 100.0, '\uE200'),
		createSpecPoint(domain.Moon, true, true, 100.0, '\uE201'),
		createSpecPoint(domain.Mercury, true, true, 80.0, '\uE202'),
		createSpecPoint(domain.Venus, true, true, 80.0, '\uE203'),
		createSpecPoint(domain.Earth, false, false, 100.0, '\uE204'),
		createSpecPoint(domain.Mars, true, true, 80.0, '\uE205'),
		createSpecPoint(domain.Jupiter, true, true, 60.0, '\uE206'),
		createSpecPoint(domain.Saturn, true, true, 60.0, '\uE207'),
		createSpecPoint(domain.Uranus, true, true, 60.0, '\uE208'),
		createSpecPoint(domain.Neptune, true, true, 60.0, '\uE209'),
		createSpecPoint(domain.Pluto, true, true, 60.0, '\uE210'),
		createSpecPoint(domain.Chiron, true, true, 40.0, '\uE400'),
		createSpecPoint(domain.Nessus, false, false, 10.0, '\uE401'),
		createSpecPoint(domain.Pholus, false, false, 10.0, '\uE402'),
		createSpecPoint(domain.Varuna, false, false, 10.0, '\uE403'),
		createSpecPoint(domain.Ixion, false, false, 10.0, '\uE404'),
		createSpecPoint(domain.Quaoar, false, false, 10.0, '\uE405'),
		createSpecPoint(domain.Haumea, false, false, 10.0, '\uE406'),
		createSpecPoint(domain.Eris, false, false, 10.0, '\uE407'),
		createSpecPoint(domain.Sedna, false, false, 10.0, '\uE408'),
		createSpecPoint(domain.Orcus, false, false, 10.0, '\uE409'),
		createSpecPoint(domain.Makemake, false, false, 10.0, '\uE410'),
		createSpecPoint(domain.Ceres, false, false, 10.0, '\uE411'),
		createSpecPoint(domain.Pallas, false, false, 10.0, '\uE412'),
		createSpecPoint(domain.Juno, false, false, 10.0, '\uE413'),
		createSpecPoint(domain.Vesta, false, false, 10.0, '\uE414'),
		createSpecPoint(domain.Hygieia, false, false, 10.0, '\uE415'),
		createSpecPoint(domain.Astraea, false, false, 10.0, '\uE416'),
		createSpecPoint(domain.Huya, false, false, 10.0, '\uE417'),
		createSpecPoint(domain.Ascendant, true, true, 80.0, '\uE500'),
		createSpecPoint(domain.Mc, false, false, 80.0, '\uE501'),
		createSpecPoint(domain.Vertex, false, false, 10.0, '\uE502'),
		createSpecPoint(domain.EastPoint, false, false, 10.0, '\uE503'),
		createSpecPoint(domain.NodeMean, true, true, 60.0, '\uE523'),
		createSpecPoint(domain.NodeTrue, false, false, 60.0, '\uE525'),
		createSpecPoint(domain.ApogeeMean, false, false, 40.0, '\uE530'),
		createSpecPoint(domain.ApogeeCorrected, false, false, 40.0, '\uE531'),
		createSpecPoint(domain.ApogeeDuval, false, false, 40.0, '\uE531'),
		createSpecPoint(domain.CupidoUra, false, false, 10.0, '\uE600'),
		createSpecPoint(domain.HadesUra, false, false, 10.0, '\uE601'),
		createSpecPoint(domain.ZeusUra, false, false, 10.0, '\uE602'),
		createSpecPoint(domain.KronosUra, false, false, 10.0, '\uE603'),
		createSpecPoint(domain.ApollonUra, false, false, 10.0, '\uE604'),
		createSpecPoint(domain.AdmetosUra, false, false, 10.0, '\uE605'),
		createSpecPoint(domain.VulcanusUra, false, false, 10.0, '\uE606'),
		createSpecPoint(domain.PoseidonUra, false, false, 10.0, '\uE607'),
		createSpecPoint(domain.PersephoneRam, false, false, 40.0, '\uE608'),
		createSpecPoint(domain.HermesRam, false, false, 40.0, '\uE609'),
		createSpecPoint(domain.DemeterRam, false, false, 40.0, '\uE610'),
		createSpecPoint(domain.Isis, false, false, 40.0, '\uE611'), // also TransPluto
		createSpecPoint(domain.PersephoneCarteret, false, false, 40.0, '\uE612'),
		createSpecPoint(domain.VulcanusCarteret, false, false, 40.0, '\uE613'),
	}
}

func createSpecPoint(actPoint domain.ChartPoint, isUsed, show bool, orbF float64, glyph rune) domain.ConfigPoint {
	return domain.ConfigPoint{
		ActualPoint: actPoint,
		IsUsed:      isUsed,
		ShowInChart: show,
		OrbFactor:   orbF,
		Glyph:       glyph,
	}
}

func createProg() domain.ConfigProg {
	return domain.ConfigProg{
		TransitPoints: []domain.ChartPoint{
			domain.Sun, domain.Moon, domain.Mercury, domain.Venus, domain.Mars, domain.Jupiter, domain.Saturn,
			domain.Uranus, domain.Neptune, domain.Pluto, domain.Chiron, domain.NodeMean,
		},
		SecDirPoints: []domain.ChartPoint{
			domain.Sun, domain.Moon, domain.Mercury, domain.Venus, domain.Mars, domain.Jupiter,
		},
		SymDirPoints: []domain.ChartPoint{
			domain.Sun, domain.Moon, domain.Mercury, domain.Venus, domain.Mars, domain.Jupiter, domain.Saturn,
			domain.Uranus, domain.Neptune, domain.Pluto, domain.Chiron, domain.NodeMean, domain.Ascendant, domain.Mc,
		},
		SymDirTimeKey: domain.SymKeyOneDegree,
		PrimDirSign: []domain.ChartPoint{domain.Sun, domain.Moon, domain.Mercury, domain.Venus, domain.Mars,
			domain.Jupiter, domain.Saturn, domain.Ascendant, domain.Mc},
		PrimDirProm: []domain.ChartPoint{domain.Sun, domain.Moon, domain.Mercury, domain.Venus, domain.Mars,
			domain.Jupiter, domain.Saturn},
		PrimDirMethod:  domain.MethodPlacidus,
		PrimDirTimeKey: domain.PdKeyNaibod,
		PrimDirMundane: true,
		SolarRelocate:  false,
	}
}
