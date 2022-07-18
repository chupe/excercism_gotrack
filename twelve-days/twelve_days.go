package twelve

import "fmt"

type Day struct{ ordinal, gift string }

var Days = [12]Day{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

func Verse(i int) string {
	i -= 1
	ordinal := Days[i].ordinal
	var gifts string
	for j := i; j >= 0; j-- {
		gifts += Days[j].gift
		if j > 0 {
			gifts += ", "
		}
		if j == 1 {
			gifts += "and "
		}
	}
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", ordinal, gifts)
}

func Song() string {
	var song string
	for i := 1; i <= 12; i++ {
		song += Verse(i)
		if i < 12 {
			song += "\n"
		}
	}
	return song
}
