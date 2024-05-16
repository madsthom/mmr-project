data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "postgres://localhost:5432/mmr_project?user=postgres&password=this_is_a_hard_password1337&sslmode=disable"
  url = "postgres://localhost:5432/mmr_project?user=postgres&password=this_is_a_hard_password1337&sslmode=disable"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}