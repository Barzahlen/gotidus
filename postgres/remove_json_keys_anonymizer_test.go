package postgres

import "testing"

func TestRemoveJSONKeysAnonymizerBuild(t *testing.T) {
	tableName := "foo"
	columnName := "bar"

	cases := []struct {
		title string
		keys  []string

		expectedString string
	}{
		{
			title: "remove one key",
			keys:  []string{"one_key"},

			expectedString: `(
      SELECT
        concat('{', string_agg(to_json("key") || ':' || "value", ','), '}')::JSON
      FROM json_each(foo.bar WHERE key <> 'one_key')
    )`,
		},
		{
			title: "remove two keys",
			keys:  []string{"one_key", "another_key"},

			expectedString: `(
      SELECT
        concat('{', string_agg(to_json("key") || ':' || "value", ','), '}')::JSON
      FROM json_each(foo.bar WHERE key <> 'one_key' AND key <> 'another_key')
    )`,
		},
	}

	for _, c := range cases {
		t.Run(c.title, func(t *testing.T) {

			anonymizer := NewRemoveJSONKeysAnonymizer(c.keys)

			compareStrings(
				anonymizer.Build(tableName, columnName),
				c.expectedString,
				t,
			)

		})
	}
}