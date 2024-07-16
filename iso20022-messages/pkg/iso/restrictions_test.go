package iso_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/moov-io/base"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

func TestValidateMultipleErrors(t *testing.T) {
	var errs base.ErrorList = base.ErrorList{}
	iso.AddError(&errs, "Max35TextTCH", pacs_008_001_08.Max35Text("B20230931145322200000057A11712044729M").Validate())
	require.Len(t, errs, 1)
	require.ErrorContains(t, errs.Err(), "Max35TextTCH: B20230931145322200000057A11712044729M fails validation with pattern M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*")
}

func TestValidatePattern(t *testing.T) {
	pattern := `[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))((([0-1][0-9])|(2[0-3]))(([0-5][0-9])){2})[A-Z0-9]{11}.*`
	require.NoError(t, iso.ValidatePattern("20230713145322200000057A11712044729", pattern))
	require.Error(t, iso.ValidatePattern("20230931145322200000057A11712044729", pattern)) // invalid MMDD

	t.Run("concurrent", func(t *testing.T) {
		iterations := 1000

		var wg sync.WaitGroup
		wg.Add(iterations)

		for i := 0; i < iterations; i++ {
			go func(idx int) {
				wg.Done()

				patern := fmt.Sprintf("[0-9]{0,%d}", idx/5) // add/get regexes from cache
				require.NoError(t, iso.ValidatePattern("4", patern))
			}(i)
		}

		wg.Wait()
	})
}

func TestValidateEnumeration(t *testing.T) {
	enumVals := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	require.NoError(t, iso.ValidateEnumeration("A", enumVals...))
	require.Error(t, iso.ValidateEnumeration("Z", enumVals...))
}

func TestValidateLength(t *testing.T) {
	length := 5
	require.NoError(t, iso.ValidateLength("abcde", length))
	require.Error(t, iso.ValidateLength("abcdef", length))
	require.NoError(t, iso.ValidateLength("abµde", length))
	require.Error(t, iso.ValidateLength("abµdef", length))
}

func TestValidateMinLength(t *testing.T) {
	minLength := 5
	require.NoError(t, iso.ValidateMinLength("abcde", minLength))
	require.Error(t, iso.ValidateMinLength("abcd", minLength))
	require.NoError(t, iso.ValidateMinLength("abµde", minLength))
	require.Error(t, iso.ValidateMinLength("abµd", minLength))
}

func TestValidateMaxLength(t *testing.T) {
	maxLength := 5
	require.NoError(t, iso.ValidateMaxLength("abcde", maxLength))
	require.Error(t, iso.ValidateMaxLength("abcdef", maxLength))
	require.NoError(t, iso.ValidateMaxLength("abµde", maxLength))
	require.Error(t, iso.ValidateMaxLength("abµdef", maxLength))
}

func TestValidateMinInclusive(t *testing.T) {
	minVal := 3
	require.Error(t, iso.ValidateMinInclusive(3, minVal))
	require.NoError(t, iso.ValidateMinInclusive(5, minVal))
	require.Error(t, iso.ValidateMinInclusive(2, minVal))
}

func TestValidateMaxInclusive(t *testing.T) {
	maxVal := 3
	require.Error(t, iso.ValidateMaxInclusive(3, maxVal))
	require.NoError(t, iso.ValidateMaxInclusive(2, maxVal))
	require.Error(t, iso.ValidateMaxInclusive(5, maxVal))
}

func TestValidateMinExclusive(t *testing.T) {
	minVal := 3
	require.NoError(t, iso.ValidateMinExclusive(3, minVal))
	require.NoError(t, iso.ValidateMinExclusive(5, minVal))
	require.Error(t, iso.ValidateMinExclusive(2, minVal))
}

func TestValidateMaxExclusive(t *testing.T) {
	maxVal := 3
	require.NoError(t, iso.ValidateMaxExclusive(3, maxVal))
	require.NoError(t, iso.ValidateMaxExclusive(2, maxVal))
	require.Error(t, iso.ValidateMaxExclusive(5, maxVal))
}

func TestValidateFractionDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, iso.ValidateFractionDigits("3.1415", maxVal))
	require.NoError(t, iso.ValidateFractionDigits("3.14159", maxVal))
	require.Error(t, iso.ValidateFractionDigits("3.141592", maxVal))
}

func TestValidateTotalDigits(t *testing.T) {
	maxVal := 5
	require.NoError(t, iso.ValidateTotalDigits("3.141", maxVal))
	require.NoError(t, iso.ValidateTotalDigits("3.1415", maxVal))
	require.Error(t, iso.ValidateTotalDigits("3.14159", maxVal))
}
