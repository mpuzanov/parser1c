package parser1c

import (
	"testing"
)

var testData1, testData2 string

func init() {
	testData1 =
		`1CClientBankExchange
	ВерсияФормата=1.02
	Кодировка=Windows
	Отправитель=
	Получатель=
	ДатаСоздания=12.07.2019
	ВремяСоздания=11:43:24
	ДатаНачала=04.07.2019
	ДатаКонца=12.07.2019
	РасчСчет=12345678901234567890
	СекцияРасчСчет
	ДатаНачала=04.07.2019
	ДатаКонца=12.07.2019
	НачальныйОстаток=7133.27
	РасчСчет=12345678901234567890
	ВсегоСписано=150840
	ВсегоПоступило=170529.31
	КонечныйОстаток=26822.58
	КонецРасчСчет	
	СекцияДокумент=Платежное поручение
	Номер=162
	Дата=10.07.2019
	Сумма=1500.00
	ДатаПоступило=11.07.2019
	ПлательщикСчет=40703810935000003034
	Плательщик=ТСН "И. ЗАКИРОВА, 9"
	ПлательщикИНН=1841049380
	ПлательщикРасчСчет=40703810935000003034
	ПлательщикБанк1=АКБ "ИЖКОМБАНК" (ПАО)
	ПлательщикБИК=049401871
	ПлательщикКорсчет=30101810900000000871
	ПолучательСчет=40802810964280000719
	ПолучательБанк1=Филиал "Пермский" ПАО КБ "УБРиР"
	ВидПлатежа=Электронно
	ВидОплаты=01
	НазначениеПлатежа=Оплата за июнь 2019 НДС не облагается.
	КонецДокумента
	СекцияДокумент=Платежное требование
	Номер=8172053
	Дата=09.07.2019
	Сумма=50.00
	ДатаСписано=09.07.2019
	ПлательщикБИК=045773883
	Получатель=Комиссия за перевод с исп.систем удаленного доступа (межбанк), ф-л «Пермский», ОО «Ижевский»   (ИП)
	ПолучательИНН=6608008004
	ПолучательРасчСчет=70601810644282740377
	ПолучательБанк1=Филиал "Пермский" ПАО КБ "УБРиР"
	ПолучательБИК=045773883
	ПолучательКорсчет=30101810500000000883
	ВидПлатежа=Электронно
	НазначениеПлатежа1=Комиссия за перевод с исп.систем удаленного доступа (межбанк), ф-л «Пермский», ОО «Ижевский»   (ИП) за 09.07.2019 г. в кол-ве штук : 2 согл. дог. банковского счета № 6428-РС-689 от 08.10.2014
	НазначениеПлатежа2= НДС не облагается
	КонецДокумента
	`

	testData2 = `
	СекцияДокумент=Платежное поручение
	Номер=162
	Дата=10.07.2019
	Сумма=1500.00
	ДатаПоступило=11.07.2019
	ПлательщикСчет=40703810935000003034
	Плательщик=ТСН "И. ЗАКИРОВА, 9"
	ПлательщикИНН=1841049380
	ПлательщикРасчСчет=40703810935000003034
	ПлательщикБанк1=АКБ "ИЖКОМБАНК" (ПАО)
	ПлательщикБИК=049401871
	ПлательщикКорсчет=30101810900000000871
	ПолучательСчет=40802810964280000719
	Получатель=ИП ПУЗАНОВ МИХАИЛ АНАТОЛЬЕВИЧ
	ПолучательИНН=183201986640
	ПолучательРасчСчет=40802810964280000719
	ПолучательБанк1=Филиал "Пермский" ПАО КБ "УБРиР"
	ПолучательБИК=045773883
	ПолучательКорсчет=30101810500000000883
	ВидПлатежа=Электронно
	ВидОплаты=01
	НазначениеПлатежа=Оплата за июнь 2019 НДС не облагается.
	`

}

func TestImportData(t *testing.T) {
	want := 2
	got, err := ImportData(testData1)
	if err != nil {
		t.Errorf("got=%v, expected=%v, error: %v", got, want, err)
	} else {
		if len(got.Docs) != want {
			t.Errorf("got=%v, expected=%v", got, want)
		}
	}
}

func TestGetStringDoc(t *testing.T) {
	want := 2
	got := getStringDoc(testData1)
	if len(got) != want {
		t.Errorf("got=%v, expected=%v", got, want)
	}
}

func TestGetValueName(t *testing.T) {
	testCases := []struct {
		desc string
		name string
		want string
	}{
		{
			desc: "Сумма",
			name: "Сумма",
			want: "1500.00",
		},
		{
			desc: "Дата",
			name: "Дата",
			want: "10.07.2019",
		},
		{
			desc: "ПлательщикРасчСчет",
			name: "ПлательщикРасчСчет",
			want: "40703810935000003034",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := getValueName(tC.name, testData2)
			if got != tC.want {
				t.Errorf("%s, got=%v, expected=%v", tC.desc, got, tC.want)
			}
		})
	}
}
