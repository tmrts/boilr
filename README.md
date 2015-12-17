# tmplt

Are you doing the same steps over and over again every time you start a new programming project?

`tmplt` is a powerful language-agnostic command-line project templating scaffolding tool here to help you.

# Features
- Blazingly Fast
- No dependencies (NodeJS, Python Interpreter etc.)
- Full power of [golang templates](https://golang.org/pkg/text/template/) (Easy to learn & powerful)

It is under *heavy construction* at the moment.

# Installation
Download install script and run `./install` to install `tmplt`.

The `tmplt` binary will be installed to `~/bin/tmplt`.

# Usage
Use `tmplt help` to get the list of available commands.

## Reporting Issues
You can report issues directly from the command-line. Use the command, `tmplt report`.
A markdown file will be opened where the first line is the issue title and the rest 
is the issue body. After creating the issue, save & exit the editor and you will be
prompted for github credentials needed to create the issue.

## Reporting Issues
You can report issues directly from the command-line. Use the command, `tmplt report`.
After entering your github credentials, a markdown file will be opened
where the first line is the issue title and the rest is the issue body.

## Download Template
In order to download a template from a github repository, use the following command:

```bash 
tmplt download <github-repo-path> <template-name>
tmplt download tmrts/tmplt-example example
``` 

The downloaded template will be saved to local `tmplt` registry.

## Save Local Template
In order to save a template from filesystem to the template registry use the following command:

```bash 
tmplt save <template-path> <template-name>
tmplt save ~/tmplt-example example
``` 

The saved template will be saved to local `tmplt` registry.

## Use Template
In order to use a template from template registry use the following command:

```bash 
tmplt use <template-name> <target-dir>
tmplt use example ~/Workspace/example-project/
``` 

You will be prompted for values when using a template.

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
    {{now "Mon Jan 2 15:04:05 -0700 MST 2006"}}.log
```

`LICENSE` file:
```txt
{{if License == "MIT"}}
// MIT License

{{else if License == "GNU GPL v3.0"}}
// GNU GPL v3.0 License

{{else if License == "Apache Software License 2.0"}}
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

Directory/File names can also be templated: 

- `{{Name}}.go` file will be `example-project.go`.
- `{{now "Mon_Jan_2_15:04_2006"}}.log` file will be formatted with the given example
time format using the current time. It will become `Mon_Dec_14_15:08_2015.log`

**Note:**
- Defined values are by convention, capital CamelCase and functions are lowercase camelCase.
- The user will be prompted for a choice for each value in the `project.json` template
- Only the contents of the `template` folder will be copied.

# Feedback
If you'd like to contribute, share your opinions or learn more, please feel free to open an issue.
At this stage, user feedback is of **utmost importance**, every contribution is welcome however small it may be.
