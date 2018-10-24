{{- $short := (shortname .Name "err" "res" "sqlstr" "db" "XOLog") -}}
{{- $table := (schema .Schema .Table.TableName) -}}
{{- if .Comment -}}
// {{ .Comment }}
{{- else -}}
// {{ .Name }} represents a row from '{{ $table }}'.
{{- end }}
type {{ .Name }} struct {
{{- range .Fields }}
	{{ .Name }} {{ retype .Type }} `json:"{{ .Col.ColumnName }}"` // {{ .Col.ColumnName }}
{{- end }}
{{- if .PrimaryKey }}

_exists, _deleted bool
{{ end }}
}


type {{ .Name }}Service interface {
	 Does{{ .Name }}Exists({{ $short }} *{{ .Name }})(bool,error)
	 Insert{{ .Name}}({{ $short }} *{{ .Name }})(error)
	 Update{{ .Name}}({{ $short }} *{{ .Name }})(error)
	 Upsert{{ .Name }}({{ $short }} *{{ .Name }}) (error)
	 Delete{{ .Name }}({{ $short }} *{{ .Name }}) (error)
	 GetAll{{ .Name }}s() ([]*{{ .Name }}, error)
	 GetChunked{{ .Name }}s(limit int,offset int) ([]*{{ .Name }}, error)

}

type {{ .Name }}ServiceImpl struct {
		DB XODB
}

{{ if .PrimaryKey }}
// Exists determines if the {{ .Name }} exists in the database.
func ( serviceImpl *{{ .Name }}ServiceImpl) Does{{.Name}}Exists({{ $short }} *{{ .Name }}) (bool,error) {
		panic("not yet implemented")
}


// Insert inserts the {{ .Name }} to the database.
func (serviceImpl *{{ .Name }}ServiceImpl) Insert{{ .Name }}({{ $short }} *{{ .Name }}) error {
	var err error


{{ if .Table.ManualPk }}
	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO {{ $table }} (` +
		`{{ colnames .Fields }}` +
		`) VALUES (` +
		`{{ colvals .Fields }}` +
		`)`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields $short }})
	err = serviceImpl.DB.QueryRow(sqlstr, {{ fieldnames .Fields $short }}).Scan(&{{ $short }}.{{ .PrimaryKey.Name }})
	if err != nil {
		return err
	}
{{ else }}
	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO {{ $table }} (` +
		`{{ colnames .Fields .PrimaryKey.Name }}` +
		`) VALUES (` +
		`{{ colvals .Fields .PrimaryKey.Name }}` +
		`) RETURNING {{ colname .PrimaryKey.Col }}`

	// run query
	XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }})
	err = serviceImpl.DB.QueryRow(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}).Scan(&{{ $short }}.{{ .PrimaryKey.Name }})
	if err != nil {
		return err
	}
{{ end }}


	return nil
}

{{ if ne (fieldnamesmulti .Fields $short .PrimaryKeyFields) "" }}
	// Update updates the {{ .Name }} in the database.
	func (serviceImpl *{{ .Name }}ServiceImpl) Update{{ .Name }}({{ $short }} *{{ .Name }}) error {
		var err error

		{{ if gt ( len .PrimaryKeyFields ) 1 }}
			// sql query with composite primary key
			const sqlstr = `UPDATE {{ $table }} SET (` +
				`{{ colnamesmulti .Fields .PrimaryKeyFields }}` +
				`) = ( ` +
				`{{ colvalsmulti .Fields .PrimaryKeyFields }}` +
				`) WHERE {{ colnamesquerymulti .PrimaryKeyFields " AND " (getstartcount .Fields .PrimaryKeyFields) nil }}`

			// run query
			XOLog(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
			_, err = serviceImpl.DB.Exec(sqlstr, {{ fieldnamesmulti .Fields $short .PrimaryKeyFields }}, {{ fieldnames .PrimaryKeyFields $short}})
		return err
		{{- else }}
			// sql query
			const sqlstr = `UPDATE {{ $table }} SET (` +
				`{{ colnames .Fields .PrimaryKey.Name }}` +
				`) = ( ` +
				`{{ colvals .Fields .PrimaryKey.Name }}` +
				`) WHERE {{ colname .PrimaryKey.Col }} = ${{ colcount .Fields .PrimaryKey.Name }}`

			// run query
			XOLog(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
			_, err = serviceImpl.DB.Exec(sqlstr, {{ fieldnames .Fields $short .PrimaryKey.Name }}, {{ $short }}.{{ .PrimaryKey.Name }})
			return err
		{{- end }}
	}

	// Save saves the {{ .Name }} to the database.
	/*
	func ({{ $short }} *{{ .Name }}) Save(db XODB) error {
		if {{ $short }}.Exists() {
			return {{ $short }}.Update(db)
		}

		return {{ $short }}.Insert(db)
	}
	*/

	// Upsert performs an upsert for {{ .Name }}.
	//
	// NOTE: PostgreSQL 9.5+ only
	func (serviceImpl *{{ .Name }}ServiceImpl) Upsert{{ .Name }}({{ $short }} *{{ .Name }}) error {
		var err error



		// sql query
		const sqlstr = `INSERT INTO {{ $table }} (` +
			`{{ colnames .Fields }}` +
			`) VALUES (` +
			`{{ colvals .Fields }}` +
			`) ON CONFLICT ({{ colnames .PrimaryKeyFields }}) DO UPDATE SET (` +
			`{{ colnames .Fields }}` +
			`) = (` +
			`{{ colprefixnames .Fields "EXCLUDED" }}` +
			`)`

		// run query
		XOLog(sqlstr, {{ fieldnames .Fields $short }})
		_, err = serviceImpl.DB.Exec(sqlstr, {{ fieldnames .Fields $short }})
		if err != nil {
			return err
		}


		return nil
}
{{ else }}
	// Update statements omitted due to lack of fields other than primary key
{{ end }}

// Delete deletes the {{ .Name }} from the database.
func (serviceImpl *{{ .Name }}ServiceImpl) Delete{{ .Name }}({{ $short }} *{{ .Name }}) error {
	var err error


	{{ if gt ( len .PrimaryKeyFields ) 1 }}
		// sql query with composite primary key
		const sqlstr = `DELETE FROM {{ $table }}  WHERE {{ colnamesquery .PrimaryKeyFields " AND " }}`

		// run query
		XOLog(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		_, err = serviceImpl.DB.Exec(sqlstr, {{ fieldnames .PrimaryKeyFields $short }})
		if err != nil {
			return err
		}
	{{- else }}
		// sql query
		const sqlstr = `DELETE FROM {{ $table }} WHERE {{ colname .PrimaryKey.Col }} = $1`

		// run query
		XOLog(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		_, err = serviceImpl.DB.Exec(sqlstr, {{ $short }}.{{ .PrimaryKey.Name }})
		if err != nil {
			return err
		}
	{{- end }}


	return nil
}

// GetAll{{ .Name }}s returns all rows from '{{ .Schema }}.{{ .Table.TableName }}',
// ordered by "created_at" in descending order.
func (serviceImpl *{{ .Name }}ServiceImpl)GetAll{{ .Name }}s() ([]*{{ .Name }}, error) {
    const sqlstr = `SELECT ` +
        `*` +
        `FROM {{ $table }}`

    q, err := serviceImpl.DB.Query(sqlstr)
    if err != nil {
        return nil, err
    }
    defer q.Close()

    // load results
    var res []*{{ .Name }}
    for q.Next() {
        {{ $short }} := {{ .Name }}{}

        // scan
        err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
        if err != nil {
            return nil, err
        }

        res = append(res, &{{ $short }})
    }

    return res, nil
}

// GetChunked{{ .Name }}s returns pagingated rows from '{{ .Schema }}.{{ .Table.TableName }}',
// ordered by "created_at" in descending order.
func(serviceImpl *{{ .Name }}ServiceImpl) GetChunked{{ .Name }}s(limit int,offset int) ([]*{{ .Name }}, error) {
    const sqlstr = `SELECT ` +
        `*` +
        `FROM {{ $table }} LIMIT $1 OFFSET $2`

    q, err := serviceImpl.DB.Query(sqlstr, limit,offset)
    if err != nil {
        return nil, err
    }
    defer q.Close()

    // load results
    var res []*{{ .Name }}
    for q.Next() {
        {{ $short }} := {{ .Name }}{}

        // scan
        err = q.Scan({{ fieldnames .Fields (print "&" $short) }})
        if err != nil {
            return nil, err
        }

        res = append(res, &{{ $short }})
    }

    return res, nil
}

{{- end }}
