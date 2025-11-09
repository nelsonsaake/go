package dsc

func ShouldSkip(table, model, generator string) bool {

	data := map[string]string{
		"generator":  generator,
		"model":      model,
		"table":      table,
		"table_name": table,
		"tablename":  table,
	}

	return New().WithData(data).WithRSrc("dbs/model").WithID(table).Is("excluded2")
}
