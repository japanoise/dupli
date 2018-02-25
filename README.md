# dupli

`dupli` is a quick clone of `fdupes -rd`. It recurses through directories
looking for duplicates, then prompts the user which they'd like to keep.

The purpose of writing it was to allow my friends who use Windows to have a
way to remove duplicates from their image folders.

Licensed MIT.

## Usage

Pass `dupli` the list of directories you want to check - e.g.

```
dupli mydirectory1 myotherdirectory
```

Or, run it with no arguments to just check the current directory:

```
cd mydirectory
dupli
```

Then, follow the prompts to clean up your duplicates.
