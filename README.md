# Objs

Objs is a minimalist, alpha-stage object store that aims to store, backup, and allow access to files
dependent on users, permissions, and capabilities.

## Usage

|              Command             |                                                       Usage                                                      |
|:--------------------------------:|:----------------------------------------------------------------------------------------------------------------:|
|            `objs list`           | Returns a complete list of all of the objects in the store, along with the user who added it and the time stamp. |
|    `objs post [file to post]`    |        Posts the given file to the object store, and returns the UUID that can be used to access the file.       |
| `objs get [UUID] [file to hold]` | Gets the file at the given UUID, and places it in the specified file.                                            |
| `objs put [UUID] [file to put]`  | Puts the specified file in the object with the given UUID, discarding the old content.                           |
| `objs remove [UUID]`             | Removes the object with the given UUID.                                                                          |

## What's Next

Next, I plan to introduce a backup feature, which would allow for logical backups of each object to
take place when the user enters a command, and the ability to revert an object to a previous backup,
without modifying any other files.