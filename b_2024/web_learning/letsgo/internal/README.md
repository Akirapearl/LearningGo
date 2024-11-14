## Rememeber

It’s important to point out that the directory name internal carries a special meaning and behavior in Go: any packages which live under this directory can only be imported by code inside the parent of the internal directory.

Or, looking at it the other way, this means that any packages under internal cannot be
imported by code outside of our project.
