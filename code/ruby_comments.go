// Note: an unset cache version is treated by ruby as “”.
// Because of this, dirtying this isn’t as simple as deleting it – we need to
// actually set a new value.

// This byte sequence is what ruby expects.
// yes that’s a paren after the second 180, per ruby.

// Inserting and having an op is kinda weird: We already know
// state zero. But ruby supports it, so go does too.

// single geo query, don’t do anything. stupid and does not make sense
// but ruby does it. Changing this will break a lot of client tests.
// just be nice and fix it here.

// Ruby sets various defaults directly in the structure and expects them to appear in cache.
// For consistency, we’ll do the same thing.
