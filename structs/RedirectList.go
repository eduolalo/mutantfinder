package structs

import "math/rand"

// RedirectList es una estructura cntenedora de listas de páginas para redirigir tráfico
// no deseado
type RedirectList struct {
	URL []string
}

// FunnyPage Método para obtener una página chistosa de la lista
func (s *RedirectList) FunnyPage() string {

	s.fill()
	return s.URL[rand.Intn(len(s.URL))]
}

// fill llena el slice con urls de páginas chistosas
func (s *RedirectList) fill() {
	s.URL = []string{
		"http://www.everydayim.com/",
		"https://www.instagram.com/jinushikeisuke/",
		"http://randomcolour.com/",
		"https://hackertyper.net/",
		"https://www.instagram.com/failedtatto/",
		"http://www.planecrashinfo.com/lastwords.htm",
		"https://www.instagram.com/tipsforjesus/",
		"http://www.staggeringbeauty.com/",
		"https://www.instagram.com/mothmeister/",
		"https://www.ratemypoo.com/",
		"https://www.instagram.com/baddiewinkle/",
		"https://www.creepypasta.com/",
		"https://www.instagram.com/tasteofstreep/",
		"https://www.forbes.com/",
		"https://www.instagram.com/craptaxidermy/",
		"http://tinytuba.com/",
		"https://www.instagram.com/benjihultsch/",
		"http://beesbeesbees.com/",
		"http://www.partridgegetslucky.com/",
		"https://www.instagram.com/tonydetroit/",
		"http://www.koalastothemax.com/",
		"http://heeeeeeeey.com/",
		"http://tinytuba.com/",
		"http://corndog.io/",
		"http://thatsthefinger.com/",
		"http://cant-not-tweet-this.com/",
		"http://weirdorconfusing.com/",
		"http://eelslap.com/",
		"http://www.staggeringbeauty.com/",
		"http://burymewithmymoney.com/",
		"http://endless.horse/",
		"http://www.trypap.com/",
		"http://www.republiquedesmangues.fr/",
		"http://www.movenowthinklater.com/",
		"http://www.partridgegetslucky.com/",
		"http://www.rrrgggbbb.com/",
		"http://beesbeesbees.com/",
		"http://www.koalastothemax.com/",
		"http://www.everydayim.com/",
		"http://randomcolour.com/",
		"http://cat-bounce.com/",
		"http://chrismckenzie.com/",
		"http://hasthelargehadroncolliderdestroyedtheworldyet.com/",
		"http://ninjaflex.com/",
		"http://ihasabucket.com/",
		"http://corndogoncorndog.com/",
		"http://www.hackertyper.com/",
		"https://pointerpointer.com",
		"http://imaninja.com/",
		"http://www.ismycomputeron.com/",
		"http://www.nullingthevoid.com/",
		"http://www.muchbetterthanthis.com/",
		"http://www.yesnoif.com/",
		"http://iamawesome.com/",
		"http://www.pleaselike.com/",
		"http://crouton.net/",
		"http://corgiorgy.com/",
		"http://www.wutdafuk.com/",
		"http://unicodesnowmanforyou.com/",
		"http://www.crossdivisions.com/",
		"http://tencents.info/",
		"http://www.patience-is-a-virtue.org/",
		"http://www.theendofreason.com/",
		"http://pixelsfighting.com/",
		"http://isitwhite.com/",
		"http://onemillionlols.com/",
		"http://www.omfgdogs.com/",
		"http://oct82.com/",
		"http://chihuahuaspin.com/",
		"http://www.blankwindows.com/",
		"http://dogs.are.the.most.moe/",
		"http://tunnelsnakes.com/",
		"http://www.trashloop.com/",
		"http://www.ascii-middle-finger.com/",
		"http://spaceis.cool/",
		"http://www.donothingfor2minutes.com/",
		"http://buildshruggie.com/",
		"http://buzzybuzz.biz/",
		"http://yeahlemons.com/",
		"http://burnie.com/",
		"http://wowenwilsonquiz.com",
		"https://thepigeon.org/",
		"http://notdayoftheweek.com/",
		"http://www.amialright.com/",
		"http://nooooooooooooooo.com/",
		"https://cant-not-tweet-this.com/",
	}
}
