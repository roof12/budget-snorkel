module roof12/budget-snorkel

go 1.19

require github.com/notnil/chess v1.9.0

// Bug breaks MoveHistory(): https://github.com/notnil/chess/issues/106
// Patched repo: https://github.com/mono424/chess
replace github.com/notnil/chess => ../../../vendor/chess
