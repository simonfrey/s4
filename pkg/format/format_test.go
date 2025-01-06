package format_test

import (
	"github.com/simonfrey/s4/pkg/format"
	"testing"
)

type testCase struct {
	TravelFormat string
	UsesAES      bool
	Version      float32
	Data         string
	ValidFormat  bool
}

func TestOldFormat(t *testing.T) {
	oldFormat := []testCase{
		{
			TravelFormat: "====== BEGIN [s4 v0.5 || S4]======\nHk0kxzGjgswt3IWI2j1flybl\n====== END   [s4 v0.5||S4]======\n",
			UsesAES:      false,
			Version:      0.5,
			Data:         "Hk0kxzGjgswt3IWI2j1flybl",
			ValidFormat:  true,
		},
		{
			TravelFormat: "====== BEGIN [s4 v0.5 || S4]======\ncWBDjs1dabgdyvmTrmeLtd5t\n====== END   [s4 v0.5||S4]======",
			UsesAES:      false,
			Version:      0.5,
			Data:         "cWBDjs1dabgdyvmTrmeLtd5t",
			ValidFormat:  true,
		},
		{
			TravelFormat: "====== BEGIN [s4 v0.5 || S4]======\n79XrU7KiobypPcsJqhYa7boW\n====== END   [s4 v0.5||S4]======\n",
			UsesAES:      false,
			Version:      0.5,
			Data:         "79XrU7KiobypPcsJqhYa7boW",
			ValidFormat:  true,
		},
		{
			TravelFormat: "====== BEGIN [s4 v0.5 || S4]======\n79XrU7KiobypPcsJqhYa7boW\n======   [s4 v0.5S4]======\n",
			UsesAES:      false,
			Version:      0.5,
			Data:         "79XrU7KiobypPcsJqhYa7boW",
			ValidFormat:  true,
		},
		{
			TravelFormat: "====== BEGIN [s4 v0.5    || AES+S4]=======\nyM3cvdrIxDS7Y7LH8RfwEceaz0FgzvsT8btGOFQJqa\n6OCio9Xz1fPV89KgoKq/Iu823RltznehJQBYpilIqYfWoTNSjg9ZV39DsgzJlzymX6a09wjDfGgq5Kaq01a0wVb4N\n6u7sYuy9C8Skds9v5gRAZMZm4\n====== END   [s4 v0.5||AES+S4]======\n",
			UsesAES:      true,
			Version:      0.5,
			Data:         "yM3cvdrIxDS7Y7LH8RfwEceaz0FgzvsT8btGOFQJqa6OCio9Xz1fPV89KgoKq/Iu823RltznehJQBYpilIqYfWoTNSjg9ZV39DsgzJlzymX6a09wjDfGgq5Kaq01a0wVb4N6u7sYuy9C8Skds9v5gRAZMZm4",
			ValidFormat:  true,
		},
		{
			TravelFormat: "====== BEGIN [s4 v0.5 || AES+S4]======\nVK+628YsyzieyqhWjpQigRU\n3JfEI7mNY2ecz7pip/4s6Cio9Xz1fPV89KgoKq/Iu823RltznehJQBYpilIqYfWoTNSjg9ZV39DsgzJlzymX6a09wjDfGgq5Kaq01a0wVb4N6u7sYuy9C8Skds9v5gRAZMZm4\n====== END   [s4 v0.5||AES+S4]======\n",
			UsesAES:      true,
			Version:      0.5,
			Data:         "VK+628YsyzieyqhWjpQigRU3JfEI7mNY2ecz7pip/4s6Cio9Xz1fPV89KgoKq/Iu823RltznehJQBYpilIqYfWoTNSjg9ZV39DsgzJlzymX6a09wjDfGgq5Kaq01a0wVb4N6u7sYuy9C8Skds9v5gRAZMZm4",
			ValidFormat:  true,
		},
		{
			TravelFormat: "======BEGIN[s4v0.5||AES+S4]======\nVP3X/KW2X3xvyUSFk53A7ICfR4gMan9pG8NVljFe6wACCio9Xz1fPV89KgoKFcK7hKesmVPeROw6GXufT3uhwgug7fojG84L0x/HnYuuKGIKUQFy++hrRxYr8Ane6GRopubFlauUovyAQ/KWJASuXwvy7dLsSg==\n====== END   [s4 v0.5||AES+S4]======\n",
			UsesAES:      true,
			Version:      0.5,
			Data:         "VP3X/KW2X3xvyUSFk53A7ICfR4gMan9pG8NVljFe6wACCio9Xz1fPV89KgoKFcK7hKesmVPeROw6GXufT3uhwgug7fojG84L0x/HnYuuKGIKUQFy++hrRxYr8Ane6GRopubFlauUovyAQ/KWJASuXwvy7dLsSg==",
			ValidFormat:  true,
		},
	}
	newFormat := []testCase{
		{
			TravelFormat: "[s4 v0.5 || S4\nHk0kxzGjgswt3IWI2j1flybl\n]",
			UsesAES:      false,
			Version:      0.5,
			Data:         "Hk0kxzGjgswt3IWI2j1flybl",
			ValidFormat:  true,
		},
		{
			TravelFormat: "[s4 v0.5 S4 hUjjRUlRlA==]",
			UsesAES:      false,
			Version:      0.5,
			Data:         "hUjjRUlRlA==",
			ValidFormat:  true,
		},
	}
	v06Format := []testCase{
		{
			TravelFormat: "[s4 v0.6 AES+S4\n  rbec s6r8 97xw 8352 fkkf\n  7696 k8zx 2zp8 363r 4e4p\n  67tr b2ha 8h2w 48px rp9y\n  f334 4par fw44 pzrp hzw8\n  9c75 6kxh 9w87 rwzh z6p3\n  yt7w 36f7 sxzw 5rwf t7gb\n  3x9w cw3x cef8 72pg kh34\n]",
			UsesAES:      true,
			Version:      0.6,
			ValidFormat:  true,
			Data:         "rbecs6r897xw8352fkkf7696k8zx2zp8363r4e4p67trb2ha8h2w48pxrp9yf3344parfw44pzrphzw89c756kxh9w87rwzhz6p3yt7w36f7sxzw5rwft7gb3x9wcw3xcef872pgkh34",
		},
		{
			TravelFormat: "[s4 v0.6 S4\n  9fp9 kbtg 9y62 k4\n]",
			UsesAES:      false,
			Version:      0.6,
			ValidFormat:  true,
			Data:         "9fp9kbtg9y62k4",
		},
	}
	createFormatTestCases := []testCase{
		{
			TravelFormat: format.CreateTravelFormat(format.Format{
				UseAES:  false,
				Version: 0.5,
				Data:    "abc",
			}),
			UsesAES:     false,
			Version:     0.5,
			Data:        "abc",
			ValidFormat: true,
		},
		{
			TravelFormat: format.CreateTravelFormat(format.Format{
				UseAES:                    false,
				Version:                   0.5,
				OptimizedHumandReadbility: true,
				Data:                      "012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940",
			}),
			UsesAES:     false,
			Version:     0.5,
			Data:        "012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940",
			ValidFormat: true,
		},
		{
			TravelFormat: format.CreateTravelFormat(format.Format{
				UseAES:                    false,
				Version:                   0.6,
				OptimizedHumandReadbility: true,
				Data:                      "012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940",
			}),
			UsesAES:     false,
			Version:     0.6,
			Data:        "012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940012345678910111213141516171819202122232425262728293031323334353637383940",
			ValidFormat: true,
		},
	}

	testCases := append(append(append(oldFormat, newFormat...), createFormatTestCases...), v06Format...)

	for k, tC := range testCases {
		valid := format.IsTravelValidFormat(tC.TravelFormat)
		if valid != tC.ValidFormat {
			t.Errorf("[TestCase %d] Expected format valid to be %t but got %t", k, tC.ValidFormat, valid)
		}
		if !valid {
			continue
		}

		f, err := format.ParseTravelFormat(tC.TravelFormat)
		if err != nil {
			t.Errorf("[TestCase %d] Could not ParseTravelFormat: %s", k, err)
		}

		// Check if the parsed data is valid
		if f.UseAES != tC.UsesAES {
			t.Errorf("[TestCase %d] Expected UseAES to be %t but got %t", k, tC.UsesAES, f.UseAES)
		}
		if f.Version != tC.Version {
			t.Errorf("[TestCase %d] Expected Version to be %f but got %f", k, tC.Version, f.Version)
		}
		if f.Data != tC.Data {
			t.Errorf("[TestCase %d] Expected Data to be %s but got %s", k, tC.Data, f.Data)
		}
	}
}
