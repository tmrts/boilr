# Creating Templates
At the top-level of your repository include an optional "project.json"
file that contains the default values that you'd like to substitute

```json
{
    "Name": "example-project",
    "Author": "Tamer Tas",
    "Email": "contact@tmrts.com",
    "PrintHomeDir": true,
    "License": [
        "MIT",
        "GNU GPL v3.0",
        "Apache Software License 2.0"
    ]
}
```

Now, create a `template` folder that contains all the files that you'd like to
be part of your project template. When using a template, the contents of this
folder will be parsed and copied to the target directory requested by user

`template` directory:
```txt
template/
    LICENSE
    README.md
    {{Name}}.go
    {{time "Mon Jan 2 15:04:05 -0700 MST 2006"}}.log
```

`LICENSE` file:
```txt
{{if eq License "MIT"}}
// MIT License

{{else if eq License "GNU GPL v3.0"}}
// GNU GPL v3.0 License

{{else if eq License "Apache Software License 2.0"}}
// Apache License

{{end}}
```

`README` file:
```markdown
This project was created by {{Author}}.

This project is under the {{License}} license.

For more information please send an e-mail to `{{Email}}`.

{{if PrintHomeDir}}
During the project creation the home directory path was `{{env "HOME" | toLower}}`.
{{end}}
```

## File/Directory Name Templating

File/Directory names can also be templated:

- `{{Name}}.go` file will be `example-project.go`.
- `{{time "Mon_Jan_2_15:04_2006"}}.log` file will be formatted with the given example
time format using the current time. It will become `Mon_Dec_14_15:08_2015.log`

**Note:**
- Defined values are by convention, capital CamelCase and functions are lowercase camelCase.
- The user will be prompted for a choice for each value in the `project.json` template
- Only the contents of the `template` folder will be copied.
