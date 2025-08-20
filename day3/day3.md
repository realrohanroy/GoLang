Day 3 — Go Learning Journey 🚀
✅ What we covered today:

Slices Recap

Capacity (cap) grows (usually doubles) when slices run out of space.

Slices are references → multiple slices can point to the same underlying array.

Modifying a subslice can affect the original slice.

Maps in Go

Declared maps with make and literals.

Added, updated, and deleted keys.

Checked if a key exists using the comma-ok idiom (val, ok := m[key]).

Iterated over maps with for range.

📝 Exercise done:

Built a small Student Grades Manager using maps:

Added "Charlie".

Updated "Alice".

Deleted "Bob".

Checked for "David".

Printed the whole map at each step.

Bonus: Iterated to print all student grades.

🌟 Key Takeaways:

Slices = Dynamic, reference type. Watch out when sub-slicing — changes reflect on the original.

Maps = Key-value store, unordered. Use delete and comma-ok idiom wisely.

Both slices and maps are reference types in Go (very important).

🔜 For Day 4 (tomorrow):

More practice with maps (nested maps, map of slices).

Introduction to structs (custom types).

Possibly combine slices + maps + structs into a mini project (e.g., student management system).