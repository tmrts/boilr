# Usage
Use `boilr help` to get the list of available commands.

## Download Template
In order to download a template from a github repository, use the following command:

```bash
boilr template download <github-repo-path> <template-name>
boilr template download tmrts/boilr-license license
```

The downloaded template will be saved to local `boilr` registry.

## Save Local Template
In order to save a template from filesystem to the template registry use the following command:

```bash
boilr template save <template-path> <template-name>
boilr template save ~/boilr-license license
```

The saved template will be saved to local `boilr` registry.

## Use Template
In order to use a template from template registry use the following command:

```bash
boilr template use <template-name> <target-dir>
boilr template use license ~/Workspace/example-project/
``

You will be prompted for values when using a template.
