data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./loader",
  ]
}

variable "db_url" {
  type    = string
  default = "postgres://localhost:5432/mmr_project?user=postgres&password=this_is_a_hard_password1337&sslmode=disable"
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = var.db_url
  url = var.db_url
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}