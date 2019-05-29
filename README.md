# Shanty
Shanty is an app to help when writing music to show notes that go together.

## Usage

```/bin/bash
./shanty --method circle --note c --number 5
Root Note: c
g
d
a
e
b
```

```/bin/bash
./shanty --method major --note c                         
Root Note: c
[d e f g a b c]
```

```/bin/bash
./shanty --method naturalMinor --note c
Root Note: c
[d d# f g g# a# c]

```

```/bin/bash
./shanty --method harmonicMinor --note c
Root Note: c
[d d# f g g# b c]
```

```/bin/bash
./shanty --method melodicMinor --note c
Root Note: c
[d d# f g a b c]
```

### TODO
Need to change flats and sharps when appropriate
Need to add blues and other scales
