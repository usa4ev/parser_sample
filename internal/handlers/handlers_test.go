package handlers

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_Scan(t *testing.T) {
	type (
		node struct {
			val      int
			children []any
		}
	)

	tests := []struct {
		name     string
		testFile string
		want     company
	}{
		{
			"Юр. лицо",
			"testdata/ooo",
			company{
				Name: "ОБЩЕСТВО С ОГРАНИЧЕННОЙ ОТВЕТСТВЕННОСТЬЮ \"КЕХ ЕКОММЕРЦ\"",
				INN:  "7710668349",
				KPP:  "771001001",
				CEO:  "Правдивый Владимир Анатольевич",
			},
		},
		{
			"AO",
			"testdata/ao",
			company{
				Name: "АКЦИОНЕРНОЕ ОБЩЕСТВО \"ИНДУСТРИАЛЬНЫЙ ПАРК \"РОДНИКИ\"",
				INN:  "3701046986",
				KPP:  "370101001",
				CEO:  "Голятин Андрей Олегович",
			},
		},
		{
			"ЗАО",
			"testdata/zao",
			company{
				Name: "ЗАКРЫТОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО \"ФАРМАЛЬЯНС\"",
				INN:  "4401130886",
				KPP:  "440101001",
				CEO:  "Иванов Денис Сергеевич",
			},
		},
		{
			"ХП",
			"testdata/hp",
			company{
				Name: "ХОЗЯЙСТВЕННОЕ ПАРТНЕРСТВО \"ИНДУСТРИАЛЬНЫЙ ПАРК \"ВОЛГОРЕЧЕНСКИЙ\"\"",
				INN:  "4431004649",
				KPP:  "443101001",
				CEO:  "Зайцев Андрей Иванович",
			},
		},
	}

	wd,_ := os.Getwd()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, _ := os.Open(filepath.Join(wd, tt.testFile))
			company, err := ScrapCompany(f)
			require.NoError(t, err)

			assert.Equalf(t, tt.want, company, "company was not filled properly")

		})
	}

}
