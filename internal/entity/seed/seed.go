package seed

import "github.com/brianvoe/gofakeit/v6"

func (s Seeder) Seed() {
	if len(s.Entities) == 0 {
		return
	} else {
		if err := s.Client.Schema.Create(s.Context); err != nil {

		}
		for k := 0; k < len(s.Entities); k++ {
			switch s.Entities[k] {
			case "Faq":
				{
					for i := 0; i < s.SeedInt; i++ {
						control := true
						if i%2 == 0 {
							control = true
						} else {
							control = false
						}
						s.Client.Faq.Create().SetQuestion(gofakeit.Question()).SetAnswer(gofakeit.Name()).SetStatus(control).Exec(s.Context)
					}
					break
				}
			case "Message":
				{
					for i := 0; i < s.SeedInt; i++ {
						control := true
						if i%2 == 0 {
							control = true
						} else {
							control = false
						}
						s.Client.Message.Create().SetName(gofakeit.Name()).SetEmail(gofakeit.Email()).SetIP(gofakeit.IPv4Address()).SetMessage(gofakeit.LoremIpsumWord()).SetStatus(control).SetPhone(gofakeit.Phone()).SetSubject(gofakeit.LoremIpsumSentence(2)).Exec(s.Context)
					}
					break
				}
			default:
				break

			}
		}
	}

}
