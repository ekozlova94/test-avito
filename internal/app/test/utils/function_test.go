package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_IsEmailValid(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool
	}{
		{
			name:     "Correct_1",
			email:    "test@gmail.com",
			expected: true,
		},
		{
			name:     "Correct_2",
			email:    "test@yandex.ru",
			expected: true,
		},
		{
			name:     "Correct_3",
			email:    "test@google.com",
			expected: true,
		},
		{
			name:     "Incorrect_1",
			email:    "test@gmail.con",
			expected: false,
		},
		{
			name:     "Incorrect_2",
			email:    "test@example-google.com",
			expected: false,
		},
		{
			name:     "Incorrect_3",
			email:    "eccbg5cgi9dd6djgnvlj7itntcc8dduij908hdinfgj",
			expected: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsEmailValid(tc.email)
			require.Equal(t, tc.expected, result)
		})
	}
}

func Test_IsURLValid(t *testing.T) {
	testCases := []struct {
		name     string
		link     string
		expected string
		err      error
	}{
		{
			name:     "Correct_1",
			link:     "https://m.avito.ru/penza/predlozheniya_uslug/elektrik._elektrika_bez_posredniko_186174581",
			expected: "186174581",
			err:      nil,
		},
		{
			name:     "Correct_2",
			link:     "https://www.avito.ru/penza/predlozheniya_uslug/elektromontazh_elektrik_909388249",
			expected: "909388249",
			err:      nil,
		},
		{
			name:     "Correct_3",
			link:     "https://www.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika._vse_vidy_rabot_1796651105",
			expected: "1796651105",
			err:      nil,
		},
		{
			name:     "Incorrect_1",
			link:     "https://m.avit.ru/penza/predlozheniya_uslug/elektrik._elektrika_bez_posredniko_186174581",
			expected: "",
			err:      ErrIncorrect,
		},
		{
			name:     "Incorrect_2",
			link:     "http://m.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika_elektromontazh_elektrik_909388249",
			expected: "",
			err:      ErrIncorrect,
		},
		{
			name:     "Incorrect_3",
			link:     "https://avito.ru/penza/predlozheniya_uslug/uslugi_elektrika._vse_vidy_rabot_1796651105",
			expected: "",
			err:      ErrIncorrect,
		},
		{
			name:     "Incorrect_4",
			link:     "https://www.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika._vse_vidy_rabot_t1796651105",
			expected: "",
			err:      ErrIncorrect,
		},
		{
			name:     "Incorrect_5",
			link:     "https://www.avito.ru/penza/predlozheniya_uslug/uslugi_elektrika._vse_vidy_rabot",
			expected: "",
			err:      ErrIncorrect,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := IsURLValid(tc.link)
			require.Equal(t, tc.expected, result)
			require.Equal(t, tc.err, err)
		})
	}
}
