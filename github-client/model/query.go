package model

const Apache2Query = `
	query {
		license (key: "apache-2.0"){
			name
			description
		}
	}
`

const MITQuery = `
	query {
		license (key: "MIT"){
			name
			description
		}
	}
`
