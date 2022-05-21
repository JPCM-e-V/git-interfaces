# Modules

|Name|Provided interface|Used interfaces|
|-|-|-|
|git-interfaces-http|`git http protocol`|<ul><li>`git interfaces redis repo`</li><li>`git interfaces git utils`</li></ul>|
|git-interfaces-redis-repo|`git interfaces redis repo`|<ul><li>`redis server`</li></ul>|
|git-interfaces-git-utils|`git interfaces git utils`|<ul></ul>|


# Interfaces

## `git interfaces redis repo`

Go module with:

- function `Init ()`:
  
  Initializes this module

- struct `GitRef`

  `id: string` Id of the ref (20-byte, 40-hexchar sha1-sum of the object)

  `name: string` Name of the ref (e.g. `HEAD` or `refs/heads/main`)

  `type: string` The type of the ref

  `symref-target: string` The target if `type` is `symref`

- function `LsRefs (string): []GitRef`:

  Returns all object-ids with names and options

- function `GetObject (string): string`:

  Returns the content of the object.

- function `AddObject (string)`:

  Adds a new object.

- function `SetRef (GitRef)`:

  Sets / Adds the ref.

- function `DeleteRef (string)`:

  Deletes a ref.

  

## `git interfaces git utils`

- function `WriteGitProtocol (io.Writer, []string)`:

  Writes the strings as lines to the writer

- function `ReadGitProtocol (io.ReadCloser): []string`

  Reads the strings from the reader.
