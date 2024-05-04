data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./db/models",
    "--dialect", "mysql",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://mysql/8/dev"
  migration {
    dir = "file://db/migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

