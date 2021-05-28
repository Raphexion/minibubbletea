# Mini Bubbletea

My attempt at the smallest possible bubbletea tui.
Quit by pressing "q".

## Key insights

The hardest part of this trivial program is the `tick()` function and how it is used.
Especially we check what type of message we get in `Update(msg tea.Msg)`. We should
only return the `tick()` if we get a `tickMsg()`.

So the tick only comes once. Then we need to trigger it again.
