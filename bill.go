package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

var mail = [...]struct {
	from, subject string
	body          *template.Template
}{
	{
		from:    "the LRS (Larn Revenue Service)",
		subject: "undeclared income",
		body: template.Must(template.New("body").Parse(
			"\n   We have heard you survived the caverns of Larn.  Let me be the" +
				"\nfirst to congratulate you on your success.  It was quite a feat." +
				"\nIt was also very profitable for you..." +
				"\n\n   The Dungeon Master has informed us that you brought" +
				"\n{{.Gold}} gold pieces back with you from your journey.  As the" +
				"\ncounty of Larn is in dire need of funds, we have spared no time" +
				"\nin preparing your tax bill.  You owe {{.Tax}} gold pieces as" +
				"\nof this notice, and is due within 5 days.  Failure to pay will" +
				"\nmean penalties.  Once again, congratulations, We look forward" +
				"\nto your future successful expeditions.")),
	},
	{
		from:    "His Majesty King Wilfred of Larndom",
		subject: "a noble deed",
		body: template.Must(template.New("body").Parse(
			"\n   I have heard of your magnificent feat, and I, King Wilfred," +
				"\nforthwith declare today to be a national holiday.  Furthermore," +
				"\nhence three days, ye be invited to the castle to receive the" +
				"\nhonour of Knight of the realm.  Upon thy name shall it be written..." +
				"\n\nBravery and courage be yours." +
				"\n\nMay you live in happiness forevermore...")),
	},
	{
		from:    "Count Endelford",
		subject: "You Bastard!",
		body: template.Must(template.New("body").Parse(
			"\n   I have heard (from sources) of your journey.  Congratulations!" +
				"\nYou Bastard!  With several attempts I have yet to endure the" +
				" caves,\nand you, a nobody, makes the journey!  From this time" +
				" onward, bewarned\nupon our meeting you shall pay the price!")),
	},
	{
		from:    "Mainair, Duke of Larnty",
		subject: "High Praise",
		body: template.Must(template.New("body").Parse(
			"\n   With certainty, a hero I declare to be amongst us!  A nod of" +
				"\nfavour I send to thee.  Me thinks Count Endelford this day of" +
				"\nright breath'eth fire as of dragon of whom ye are slayer.  I" +
				"\nyearn to behold his anger and jealously.  Should ye choose to" +
				"\nunleash some of thy wealth upon those who be unfortunate, I," +
				"\nDuke Mainair, shall equal thy gift also.")),
	},
	{
		from:    "St. Mary's Children's Home",
		subject: "these poor children",
		body: template.Must(template.New("body").Parse(
			"\n   News of your great conquests has spread to all of Larndom." +
				"\nMight I have a moment of a great adventurers's time?  We here at" +
				"\nSt. Mary's Children's Home are very poor, and many children are" +
				"\nstarving.  Disease is widespread and very often fatal without" +
				"\ngood food.  Could you possibly find it in your heart to help us" +
				"\nin our plight?  Whatever you could give will help much." +
				"\n(your gift is tax deductible)")),
	},
	{
		from:    "The National Cancer Society of Larn",
		subject: "hope",
		body: template.Must(template.New("body").Parse(
			"\nCongratulations on your successful expedition.  We are sure much" +
				"\ncourage and determination were needed on your quest.  There are" +
				"\nmany though, that could never hope to undertake such a journey" +
				"\ndue to an enfeebling disease -- cancer.  We at the National" +
				"\nCancer Society of Larn wish to appeal to your philanthropy in" +
				"\norder to save many good people -- possibly even yourself a few" +
				"\nyears from now.  Much work needs to be done in researching this" +
				"\ndreaded disease, and you can help today.  Could you please see it" +
				"\nin your heart to give generously?  Your continued good health" +
				"\ncan be your everlasting reward.")),
	},
}

/*
 * function to mail the letters to the player if a winner
 */

func mailbill() {
	data := map[string]int{
		"Gold": c[GOLD],
		"Tax":  c[GOLD] * TAXRATE,
	}
	for _, m := range mail {
		var buf bytes.Buffer
		if err := m.body.Execute(&buf, data); err != nil {
			log.Printf("Failed executing mail body template: %v", err)
			continue
		}
		// TODO: send mail for real.
		log.Printf("\nFrom: %s\nSubject: %s\n\n%s", m.from, m.subject, buf.Bytes())
	}
	os.Exit(0)
}
