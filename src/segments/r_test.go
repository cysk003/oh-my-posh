package segments

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestR(t *testing.T) {
	cases := []struct {
		Case           string
		ExpectedString string
		Version        string
		HasRscript     bool
		HasR           bool
		HasRexe        bool
	}{
		{Case: "Rscript 4.2.0", ExpectedString: "4.2.0", HasRscript: true, Version: "Rscript (R) version 4.2.0 (2022-04-22)"},
		{Case: "Rscript 4.1.3", ExpectedString: "4.1.3", HasRscript: true, Version: "R scripting front-end version 4.1.3 (2022-03-10)"},
		{Case: "Rscript 4.1.3 patched", ExpectedString: "4.1.3", HasRscript: true, Version: "R scripting front-end version 4.1.3 Patched (2022-03-10 r81896)"},
		{Case: "Rscript 4.0.0", ExpectedString: "4.0.0", HasRscript: true, Version: "R scripting front-end version 4.0.0 (2020-04-24)"},
		{Case: "Rscript devel", ExpectedString: "4.2.0", HasRscript: true, Version: "R scripting front-end version 4.2.0 Under development (unstable) (2022-03-14 r81896)"},

		{Case: "R 4.1.2", ExpectedString: "4.1.2", HasR: true, Version: "R version 4.1.2 (2021-11-01) -- \"Bird Hippie\""},
		{Case: "R 4.1.3 patched", ExpectedString: "4.1.3", HasR: true, Version: "R version 4.1.3 Patched (2022-03-10 r81896) -- \"One Push-Up\""},
		{Case: "R 4.0.0", ExpectedString: "4.0.0", HasR: true, Version: "R version 4.0.0 (2020-04-24) -- \"Arbor Day\""},

		{Case: "R.exe 4.1.2", ExpectedString: "4.1.2", HasRexe: true, Version: "R version 4.1.2 (2021-11-01) -- \"Bird Hippie\""},
		{Case: "R.exe 4.1.3 patched", ExpectedString: "4.1.3", HasRexe: true, Version: "R version 4.1.3 Patched (2022-03-10 r81896) -- \"One Push-Up\""},
		{Case: "R.exe 4.0.0", ExpectedString: "4.0.0", HasRexe: true, Version: "R version 4.0.0 (2020-04-24) -- \"Arbor Day\""},
	}
	for _, tc := range cases {
		params := &mockedLanguageParams{
			cmd:           "R",
			versionParam:  "--version",
			versionOutput: tc.Version,
			extension:     "*.R",
		}
		env, props := getMockedLanguageEnv(params)

		env.On("HasCommand", "Rscript").Return(tc.HasRscript)
		env.On("RunCommand", "Rscript", []string{"--version"}).Return(tc.Version, nil)
		env.On("HasCommand", "R.exe").Return(tc.HasRexe)
		env.On("RunCommand", "R.exe", []string{"--version"}).Return(tc.Version, nil)

		r := &R{}
		r.Init(props, env)

		assert.True(t, r.Enabled(), fmt.Sprintf("Failed in case: %s", tc.Case))
		assert.Equal(t, tc.ExpectedString, renderTemplate(env, r.Template(), r), fmt.Sprintf("Failed in case: %s", tc.Case))
	}
}
