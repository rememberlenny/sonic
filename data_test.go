package sonic

var _SCHEMA_LESS = `
{
	"test": [
		1,
		true,
		null,
		"false",
		{
			"test": [
				1,
				true,
				null,
				"false",
				{
					"test": [
						...
					]
				}
				["a", "b"]
			]
		}
		["a", "b"]
	]
}
`